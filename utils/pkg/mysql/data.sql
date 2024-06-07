# 数据库初始化

-- 创建库
create database if not exists wu_user;

-- 切换库
use wu_user;
-- 用户表
create table if not exists users
(
    id           bigint auto_increment comment 'id' primary key,
    userAccount  varchar(256)                           not null comment '账号',
    userPassword varchar(512)                           not null comment '密码',
    userAvatar   varchar(1024)                          null comment '用户头像',
    userProfile  varchar(1024)                          null comment '用户简介',
    tags         varchar(1024)                          null comment '标签列表（json 数组）',
    userRole     varchar(256) default 'user'            not null comment '用户角色：user/admin',
    createTime   timestamp    default CURRENT_TIMESTAMP not null comment '用户创建时间',
    updateTime   timestamp    default null              null on update CURRENT_TIMESTAMP comment '用户更新时间',
    isDelete     tinyint      default 0                 not null comment '是否删除'
    ) comment '用户表' collate = utf8mb4_unicode_ci;
