CREATE DATABASE hummingbird;

USE hummingbird;

CREATE TABLE account
(
    id          bigint unsigned                     not null comment 'ID' primary key,
    name        varchar(256)                        not null comment '名称',
    passwd      varchar(256)                        not null comment '密码',
    status      int                                 not null comment '状态',
    create_time timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    update_time timestamp default CURRENT_TIMESTAMP null
) COMMENT '账号表' COLLATE = utf8mb4_unicode_ci;