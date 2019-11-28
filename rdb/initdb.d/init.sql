drop table if exists t_articles;
create table t_articles(
    id int not null auto_increment,
    title varchar(255),
    body text,
    create_at timestamp not null default current_timestamp,
    update_at timestamp not null default current_timestamp on update current_timestamp,
    logical_delete_flag boolean not null default 0,
    PRIMARY KEY(id)
);

drop table if exists c_articles_tags;
create table c_articles_tags(
    article_id int not null,
    tag_id int not null
    create_at timestamp not null default current_timestamp,
    update_at timestamp not null default current_timestamp on update current_timestamp,
    logical_delete_flag boolean not null default 0,
    PRIMARY KEY(id)
);

drop table if exists m_tags;
create table m_tags(
    id int not null auto_increment,
    name varchar(255),
    create_at timestamp not null default current_timestamp,
    update_at timestamp not null default current_timestamp on update current_timestamp,
    logical_delete_flag boolean not null default 0,
    PRIMARY KEY(id)
);
