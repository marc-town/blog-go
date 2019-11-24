DROP DATABASE IF EXISTS blog_app_db;
create database blog_app_db default character set utf8;
GRANT ALL ON blog_app_db.* TO developer@'localhost' IDENTIFIED BY 'dev';
FLUSH PRIVILEGES;

# create table
use blog_app_db;

# 社員マスタ
drop table if exists m_staff;
create table m_staff(
    id int not null auto_increment comment 'ID',
    staff_id varchar(20) not null comment '社員ID',
    staff_type_id int comment '社員種別ID',
    name varchar(15) comment '社員名',
    name_kana varchar(30) comment '社員名（かな）',
    entered_date date comment '入社日',
    birthday date comment '誕生日',
    telephone_number varchar(11) comment '電話番号',
    mail_address varchar(100) comment 'メールアドレス',
    password varchar(20) comment 'パスワード',
    create_at timestamp not null default current_timestamp comment '登録日',
    update_at timestamp not null default current_timestamp on update current_timestamp comment '更新日',
    PRIMARY KEY(id)
);

# 社員詳細テーブル
drop table if exists m_staff_info;
create table m_staff_info(
    id int not null auto_increment comment 'ID',
    staff_id int not null comment '社員ID',
    department_id int comment '部署ID',
    position_id int comment '役職ID',
    grade_id int default 1 comment 'グレードID',
    project_summary varchar(100) comment '案件概要',
    working_place varchar(50) comment '現場',
    occupation_id int comment '職種ID',
    specialty varchar(20) comment '得意分野',
    unit_price int comment '単価',
    rating int comment '評価',
    rater_id int comment '評価者ID',
    create_at timestamp not null default current_timestamp comment '登録日',
    update_at timestamp not null default current_timestamp on update current_timestamp comment '更新日',
    PRIMARY KEY(id)
);

# 社員種別マスタ
drop table if exists m_staff_type;
create table m_staff_type(
    staff_type_id int not null auto_increment comment '社員種別ID',
    staff_type_name varchar(20) comment '社員種別名',
    create_at timestamp not null default current_timestamp comment '登録日',
    update_at timestamp not null default current_timestamp on update current_timestamp comment '更新日',
    PRIMARY KEY(staff_type_id)
);

# 部署マスタ
drop table if exists m_department;
create table m_department(
    department_id int not null auto_increment comment '部署ID',
    department varchar(50) comment '部署名',
    create_at timestamp not null default current_timestamp comment '登録日',
    update_at timestamp not null default current_timestamp on update current_timestamp comment '更新日',
    PRIMARY KEY(department_id)
);

# 役職マスタ
drop table if exists m_position;
create table m_position(
    position_id int not null auto_increment comment '役職ID',
    position varchar(20) comment '役職名',
    create_at timestamp not null default current_timestamp comment '登録日',
    update_at timestamp not null default current_timestamp on update current_timestamp comment '更新日',
    PRIMARY KEY(position_id)
);

# グレードマスタ
drop table if exists m_grade;
create table m_grade(
    grade_id int not null auto_increment comment 'グレードID',
    grade varchar(10) not null comment 'グレード名',
    create_at timestamp not null default current_timestamp comment '登録日',
    update_at timestamp not null default current_timestamp on update current_timestamp comment '更新日',
    PRIMARY KEY(grade_id)
);

# 職種マスタ
drop table if exists m_occupation;
create table m_occupation(
    occupation_id int not null auto_increment comment '職種ID',
    occupation varchar(50) not null comment '職種名',
    create_at timestamp not null default current_timestamp comment '登録日',
    update_at timestamp not null default current_timestamp on update current_timestamp comment '更新日',
    PRIMARY KEY(occupation_id)
);

# 勤怠トランザクション
# year_month varchar(20) not null comment '年月',
drop table if exists t_attendance;
create table t_attendance(
    attendance_id int not null auto_increment comment '勤怠ID',
    staff_id int not null comment '社員ID',
    target_year_month varchar(20) not null comment '年月',
    target_day varchar(20) not null comment '日付',
    target_week_day varchar(20) comment '曜日',
    start_time varchar(20) comment '出勤時間',
    end_time varchar(20) comment '退勤時間',
    absence_from_work int(1) comment '欠勤',
    reason varchar(128) comment '欠勤理由',
    transportation_section int comment '区間',
    transportation_expenses int comment '交通費',
    note varchar(128) comment '備考',
    PRIMARY KEY(attendance_id)
);

# 評価トランザクション
drop table if exists t_rate;
create table t_rate(
    rate_id int not null auto_increment comment '年月',
    staff_id int not null comment '社員ID',
    target_year_month varchar(20) not null comment '社員ID',
    create_at timestamp not null default current_timestamp comment '登録日',
    update_at timestamp not null default current_timestamp on update current_timestamp comment '更新日',
    update_by int not null comment '評価者ID',
    PRIMARY KEY(rate_id)
);

# insert init data
# m_staff
insert into m_staff (staff_id, staff_type_id, name, name_kana, entered_date, birthday, telephone_number, mail_address, password) values
("admin", 1, "admin", "admin", "2019-08-07", "2017-07-28", "xxxyyyyzzzz", "admin@example.com", "admin"), # admin
("staff", 2, "staff", "staff", "2019-08-07", "2017-07-28", "xxxyyyyzzzz", "staff@example.com", "staff"), # nomal_staff
("gakkun", 2, "西平顎", "にしひらがきゅ", "2019-08-07", "2017-07-28", "xxxyyyyzzzz", "tester@example.com", "tester"); # test_staff

# t_staff_info
insert into m_staff_info (staff_id, department_id, position_id, grade_id, project_summary, working_place, occupation_id, specialty, unit_price, rating, rater_id) values
(1, 2, 1, 4, "ここより世界に痛みを", "木の葉隠れの里", 3, "新羅転生", 300, 10, 1),
(2, 1, 3, 1, "おしおし", "フィリピンpb", 2, "ぱふぱふ", 90, 10, 1),
(3, 1, 3, 1, "おぎゃー", "南国", 1, "うんち", 30,  10, 1);

# m_staff_type
insert into m_staff_type (staff_type_name) values
("admin"),
("nomal");

# m_department
insert into m_department (department) values
("開発事業部"),
("インフラ事業部"),
("営業部");

# m_position
insert into m_position (position) values
("会長"),
("平社員"),
("雑用");

# m_grade
insert into m_grade (grade_id, grade) values
(4, "Platinum"),
(3, "Gold"),
(2, "Silver"),
(1, "Green");

# m_occupation
insert into m_occupation (occupation) values
("フロントエンドエンジニア"),
("バックエンドエンジニア"),
("iOSエンジニア");
