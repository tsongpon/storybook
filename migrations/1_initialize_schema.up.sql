CREATE TABLE story
(
	id varchar(36) not null ,
	title varchar(255) not null,
	content text not null,
	author varchar(50) null,
	created_time datetime not null,
	modified_time datetime not null,
	constraint story_pk
		primary key (id)
);
