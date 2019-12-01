DROP DATABASE IF EXISTS `blog_app_db`;
CREATE DATABASE `blog_app_db` default character set utf8mb4;
GRANT ALL ON blog_app_db.* TO 'docker'@'%' IDENTIFIED BY 'docker';
FLUSH PRIVILEGES;
