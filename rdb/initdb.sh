#!/bin/bash
MYSQL=`which mysql`

$MYSQL -udev -pdev blog_app_db < /docker-entrypoint-initdb.d/init.sql

