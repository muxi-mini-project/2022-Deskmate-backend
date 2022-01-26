DROP DATABASE IF EXISTS `deskmate`;

CREATE DATABASE `deskmate`;

USE `deskmate`;

-- 用户
CREATE TABLE `users`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `student_id` VARCHAR(100) NOT NULL,-- 学号
    `password` VARCHAR(100) NOT NULL,-- 密码
    `name` VARCHAR(100) NOT NULL,
    `college` VARCHAR(100) NOT NULL,-- 学院
    -- `major` VARCHAR(100) NOT NULL,-- 专业
    `grade` VARCHAR(100) NOT NULL,
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 名片
CREATE TABLE `cards`(  
    -- 似乎不需要单独列出tag的表，直接设置就可以了
    `id` INT NOT NULL AUTO_INCREMENT COMMENT "名片id", 
    `users_id` VARCHAR(255) NOT NULL COMMENT "用户学号",
    `avatar` VARCHAR(255) COMMENT "头像",
    `nickname` VARCHAR(100) COMMENT "昵称",
    `declaration` VARCHAR(100) COMMENT "同桌宣言",
    `status` VARCHAR(100) NOT NULL,
    `infor` VARCHAR(100) COMMENT "简要信息",
    `tag1`  VARCHAR(100) COMMENT "标签一",
    `tag2`  VARCHAR(100) COMMENT "标签二",
    `tag3`  VARCHAR(100) COMMENT "标签三",
    `tag4`  VARCHAR(100) COMMENT "标签四",
    `tag5`  VARCHAR(100) COMMENT "标签五",
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 打卡
CREATE TABLE `signs`(
    `id` INT NOT NULL AUTO_INCREMENT COMMENT "打卡id",
    `time` INT NOT NULL COMMENT "打卡天数",
    `daily` VARCHAR(255) COMMENT "打卡日报",
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `users_signs`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `users_id1` VARCHAR(255) NOT NULL,
    `users_id2` VARCHAR(255) NOT NULL,
    `signs_id` INT NOT NULL,
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 申请同桌
CREATE TABLE `applications`(
    `id` INT NOT NULL AUTO_INCREMENT COMMENT "同桌申请id",
    `users_id1` VARCHAR(255) NOT NULL COMMENT "申请者",
    `users_id2` VARCHAR(255) NOT NULL COMMENT "申请对象",
    `result` VARCHAR(255) COMMENT "是否同意", -- 默认为拒绝
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 标签，我这里直接加到名片里去了，应该用不到了
CREATE TABLE `tags`(
    `id` INT NOT NULL AUTO_INCREMENT COMMENT "标签id",
    `tags_name` VARCHAR(100) NOT NULL COMMENT "标签内容",
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `cards_tags`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `cards_id`INT NOT NULL,
    `tags_id`INT NOT NULL,
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;