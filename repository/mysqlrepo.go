package repository

import (
	"database/sql"

	"github.com/tsongpon/yoneebook/model"
	"github.com/tsongpon/yoneebook/query"
)

type MysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(db *sql.DB) *MysqlRepository {
	repo := new(MysqlRepository)
	repo.db = db
	return repo
}

func (repo *MysqlRepository) GetStories(query query.StoryQuery) ([]model.Story, error) {
	stories := []model.Story{}
	sql := "SELECT id, title, content, author, created_time, modified_time FROM story ORDER BY created_time DESC LIMIT ? OFFSET ?"
	sqlLimitContent := "SELECT id, title, SUBSTRING(content, 1, 600), author, created_time, modified_time FROM story ORDER BY created_time DESC LIMIT ? OFFSET ?"
	if query.Shortcontent {
		sql = sqlLimitContent
	}
	result, err := repo.db.Query(sql, query.Limit, query.Offset)
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		s := model.Story{}
		err := result.Scan(&s.ID, &s.Title, &s.Content, &s.Author, &s.CreatedTime, &s.ModifiedTime)
		if err != nil {
			panic(err.Error())
		}
		stories = append(stories, s)
	}

	return stories, nil
}

// GetStory return story by given ID
func (repo *MysqlRepository) GetStory(id string) (model.Story, error) {
	sql := "SELECT id, title, content, author, created_time, modified_time FROM story WHERE id = ?"
	var s model.Story
	err := repo.db.QueryRow(sql, id).Scan(&s.ID, &s.Title, &s.Content, &s.Author, &s.CreatedTime, &s.ModifiedTime)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return s, nil
}

func (repo *MysqlRepository) SaveStory(story model.Story) (model.Story, error) {
	sql := "INSERT INTO story (id, title, content, author, created_time, modified_time) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := repo.db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	stmt.Exec(story.ID, story.Title, story.Content, story.Author, story.CreatedTime, story.ModifiedTime)

	return story, nil
}

func (repo *MysqlRepository) SaveStoryViewed(e model.StoryViewedEvent) error {
	sql := "INSERT INTO storyviewed (id, story_id, time, user_agent) VALUES (?, ?, ?, ?)"
	stmt, err := repo.db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	stmt.Exec(e.ID, e.StoryID, e.Time, e.UserAgent)

	return nil
}
