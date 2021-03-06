---- databse ----
USE blog_app_db;

---- drop ----
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `articles`;
DROP TABLE IF EXISTS `articles_tags`;
DROP TABLE IF EXISTS `tags`;

---- create ----
create table IF not exists `users`
(
    `id` int not null auto_increment,
    `login_id` varchar(30) not null,
    `password` varchar(50) not null,
    `email` varchar(255) not null,
    `name` varchar(255),
    `create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

create table IF not exists `articles`
(
    `id` int not null auto_increment,
    `title` varchar(255) not null,
    `body` text not null,
    `publish_flag` boolean not null default 1,
    `create_at` timestamp not null default current_timestamp,
    `update_at` timestamp not null default current_timestamp on update current_timestamp,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

create table IF not exists `articles_tags`
(
    `article_id` int not null,
    `tag_id` int not null,
    `create_at` timestamp not null default current_timestamp,
    `update_at` timestamp not null default current_timestamp on update current_timestamp,
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

create table IF not exists `tags`
(
    `id` int not null auto_increment,
    `name` varchar(255),
    `create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
