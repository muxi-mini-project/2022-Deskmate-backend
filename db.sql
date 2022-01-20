DROP DATABASE IF EXISTS `deskmate`;

CREATE DATABASE `deskmate`;

USE `deskmate`;

-- 用户
CREATE TABLE `users`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL,
    `college` VARCHAR(100) NOT NULL,-- 学院
    `major` VARCHAR(100) NOT NULL,-- 专业
    `grade` VARCHAR(100) NOT NULL,
    PRIMARY KEY(`id`)
);

-- 名片
CREATE TABLE `cards`(
    `id` INT NOT NULL AUTO_INCREMENT COMMENT "名片id", 
    `users_id` INT NOT NULL COMMENT "用户id",
    `avatar` VARCHAR(255) COMMENT "头像",
    `nickname` VARCHAR(100) COMMENT "昵称",
    `declaration` VARCHAR(100) COMMENT "同桌宣言",
    `infor` VARCHAR(100) COMMENT "简要信息",
    PRIMARY KEY(`id`)
);

-- 打卡
CREATE TABLE `signs`(
    `id` INT NOT NULL AUTO_INCREMENT COMMENT "打卡id",
    `time` INT NOT NULL COMMENT "打卡天数",
    `daily` VARCHAR(255) COMMENT "打卡日报",
    PRIMARY KEY(`id`)
);

CREATE TABLE `users_signs`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `users_id1` INT NOT NULL,
    `users_id2` INT NOT NULL,
    `signs_id` INT NOT NULL,
    PRIMARY KEY(`id`)
);

-- 申请同桌
CREATE TABLE `applications`(
    `id` INT NOT NULL AUTO_INCREMENT COMMENT "同桌申请id",
    `users_id1` INT NOT NULL COMMENT "申请者",
    `users_id2` INT NOT NULL COMMENT "申请对象",
    PRIMARY KEY(`id`)
);

-- 标签
CREATE TABLE `tags`(
    `id` INT NOT NULL AUTO_INCREMENT COMMENT "标签id",
    `tags_name` VARCHAR(100) NOT NULL COMMENT "标签内容",
    PRIMARY KEY(`id`)
);

CREATE TABLE `cards_tags`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `cards_id`INT NOT NULL,
    `tags_id`INT NOT NULL,
    PRIMARY KEY(`id`)
);