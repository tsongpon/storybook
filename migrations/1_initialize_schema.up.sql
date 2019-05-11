CREATE TABLE story
(
	id varchar(36) not null ,
	title varchar(255) not null,
	content text not null,
	author varchar(50) null,
	constraint story_pk
		primary key (id)
);
