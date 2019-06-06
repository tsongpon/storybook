CREATE TABLE storyviewed
(
	id varchar(36) not null ,
	story_id varchar(36) not null,
    user_agent varchar(255),
	time datetime not null,
	constraint userview_pk
		primary key (id)
);
