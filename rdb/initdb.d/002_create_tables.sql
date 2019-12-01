---- databse ----
USE blog_app_db;

---- drop ----
DROP TABLE IF EXISTS `articles`;
DROP TABLE IF EXISTS `articles_tags`;
DROP TABLE IF EXISTS `tags`;

---- create ----
create table IF not exists `articles`
(
    `id` int not null auto_increment,
    `title` varchar(255),
    `body` text,
    `create_at` timestamp not null default current_timestamp,
    `update_at` timestamp not null default current_timestamp on update current_timestamp,
    `logical_delete_flag` boolean not null default 0,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

create table IF not exists `articles_tags`
(
    `article_id` int not null,
    `tag_id` int not null,
    `create_at` timestamp not null default current_timestamp,
    `update_at` timestamp not null default current_timestamp on update current_timestamp,
    `logical_delete_flag` boolean not null default 0
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

create table IF not exists `tags`
(
    `id` int not null auto_increment,
    `name` varchar(255),
    `create_at` timestamp not null default current_timestamp,
    `update_at` timestamp not null default current_timestamp on update current_timestamp,
    `logical_delete_flag` boolean not null default 0,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
