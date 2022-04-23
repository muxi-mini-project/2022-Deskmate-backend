DROP DATABASE IF EXISTS `deskmate`;

CREATE DATABASE `deskmate`;

USE `deskmate`;

-- 用户
CREATE TABLE `users`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `student_id` VARCHAR(100) NOT NULL UNIQUE COMMENT "用户学号",-- 学号
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
    `users_id` VARCHAR(255) NOT NULL UNIQUE COMMENT "用户学号", -- 防止一个用户创建多个名片，将学号设置为不能重复
    `avatar` VARCHAR(255) COMMENT "头像",
    `sha`   VARCHAR(255) NULL,
    `path` VARCHAR(255) NULL,
    `nickname` VARCHAR(100) COMMENT "昵称",
    `declaration` VARCHAR(100) COMMENT "同桌宣言",
    `status` VARCHAR(100) NOT NULL COMMENT "状态", -- 0为没有同桌，1为有同桌
    `infor` VARCHAR(100) COMMENT "简要信息",
    `tag1`  VARCHAR(100) COMMENT "标签一",
    `tag2`  VARCHAR(100) COMMENT "标签二",
    `tag3`  VARCHAR(100) COMMENT "标签三",
    `tag4`  VARCHAR(100) COMMENT "标签四",
    `tag5`  VARCHAR(100) COMMENT "标签五",
    `grade` VARCHAR(100) COMMENT "年级",
    `college` VARCHAR(100) COMMENT "学院",
    `major` VARCHAR(100) COMMENT "专业",
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 打卡
CREATE TABLE `dailyrecords`(
    `id` INT NOT NULL AUTO_INCREMENT COMMENT "打卡id",
    `users_id1` VARCHAR(255) NOT NULL,
    `users_id2` VARCHAR(255) NOT NULL,
    `time`  int  COMMENT "打卡天数",
    `status` VARCHAR(255) COMMENT "打卡状态", -- 默认为进行中，结束后为已结束，用来判断是否需要打卡
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 存储打卡消息
CREATE TABLE `messages`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `dailyrecords_id` INT NOT NULL COMMENT "打卡id", -- 这个是来识别对应的某次打卡
    `time` VARCHAR(255)  NOT NULL COMMENT "发送消息的时间", -- 这里到底用什么类型还得考虑,暂时通过代码获取此刻时间再存入time中
    `information` VARCHAR(255) COMMENT "打卡的内容",
    `user_id` VARCHAR(255) COMMENT "发送人",
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE `users_signs`(-- 这个表不需要了，全部放打卡的表里了
    `id` INT NOT NULL AUTO_INCREMENT,
    `users_id1` VARCHAR(255) NOT NULL,
    `users_id2` VARCHAR(255) NOT NULL,
    `signs_id` INT NOT NULL,
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 申请同桌
CREATE TABLE `applycations`( -- 这里把原来的表名applications改成applycations了
    `id` INT NOT NULL AUTO_INCREMENT COMMENT "同桌申请id",
    `users_id1` VARCHAR(255) NOT NULL COMMENT "申请者",
    `users_id2` VARCHAR(255) NOT NULL COMMENT "申请对象",
    `result` VARCHAR(255) COMMENT "是否同意", -- 默认为空，1为同意，0为拒绝
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 标签，我这里直接加到名片里去了，应该用不到了
CREATE TABLE `tags`(
    `id` INT NOT NULL AUTO_INCREMENT COMMENT "标签id",
    `tags_name` VARCHAR(100) NOT NULL COMMENT "标签内容",
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 这里也应该用不到了
CREATE TABLE `cards_tags`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `cards_id`INT NOT NULL,
    `tags_id`INT NOT NULL,
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `updates`( -- 记录更新
    `id` INT NOT NULL AUTO_INCREMENT,
    `dailyrecords_id` INT NOT NULL COMMENT "打卡id",
    `time` VARCHAR(255)  NOT NULL COMMENT "记录更新的时间",
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;