DROP DATABASE IF EXISTS `deskmate`;

CREATE DATABASE `deskmate`;

USE `deskmate`;

CREATE TABLE `user`(
    `user_id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL,
    `major` VARCHAR(100) NOT NULL,
    `grade` INT NOT NULL,
    PRIMARY KEY(`user_id`)
);

CREATE TABLE `card`(
    `card_id` INT NOT NULL AUTO_INCREMENT COMMENT "名片id", 
    `user_id` INT NOT NULL COMMENT "用户id",
    `avatar` VARCHAR(255) COMMENT "头像",
    `nickname` VARCHAR(100) COMMENT "昵称",
    `declaration` VARCHAR(100) COMMENT "同桌宣言",
    `infor` VARCHAR(100) COMMENT "简要信息",
    PRIMARY KEY(`card_id`)
);

CREATE TABLE `sign`(
    `sign_id` INT NOT NULL AUTO_INCREMENT COMMENT "打卡id",
    `time` INT NOT NULL COMMENT "打卡天数",
    `daily` VARCHAR(255) COMMENT "打卡日报",
    PRIMARY KEY(`sign_id`)
);

CREATE TABLE `user_sign`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id1` INT NOT NULL,
    `user_id2` INT NOT NULL,
    `sign_id` INT NOT NULL,
    PRIMARY KEY(`id`)
);

CREATE TABLE `apply`(
    `apply_id` INT NOT NULL AUTO_INCREMENT COMMENT "同桌申请id",
    `user_id1` INT NOT NULL COMMENT "申请者",
    `user_id2` INT NOT NULL COMMENT "申请对象",
    PRIMARY KEY(`apply_id`)
);

CREATE TABLE `tag`(
    `tag_id` INT NOT NULL AUTO_INCREMENT COMMENT "标签id",
    `tag_name` VARCHAR(100) NOT NULL COMMENT "标签内容",
    PRIMARY KEY(`tag_id`)
);

CREATE TABLE `card_tag`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `card_id`INT NOT NULL,
    `tag_id`INT NOT NULL,
    PRIMARY KEY(`id`)
);