-- 创建文件表基本表，用于秒传，逻辑关联 用户媒体表，用户云盘表格，用户剧集表
CREATE TABLE `tbl_file`
(
    `id`             int(11)       NOT NULL AUTO_INCREMENT,
    `file_sha1`      char(40)      NOT NULL DEFAULT '' COMMENT '文件hash',
    `file_name`      varchar(256)  NOT NULL DEFAULT '' COMMENT '文件名',
    `file_size`      bigint(20)             DEFAULT '0' COMMENT '文件大小',
    `file_addr`      varchar(1024) NOT NULL DEFAULT '' COMMENT '文件存储位置',
    `create_at`      datetime               default NOW() COMMENT '创建日期',
    `update_at`      datetime               default NOW() on update current_timestamp() COMMENT '更新日期',
    `status`         tinyint       NOT NULL DEFAULT '0' COMMENT '状态( 1可用/ 2禁用/ -1已删除等状态)',
    `minitype`       char(40)      NOT NULL DEFAULT '' COMMENT '文件具体类型',
    `ftype`          tinyint       NOT NULL DEFAULT '0' COMMENT '文件状态(0图片/1视频/2音乐/3文档/4压缩包)',
    `video_duration` time          NOT NULL DEFAULT '0' COMMENT '视频时长',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_file_hash` (`file_sha1`),
    KEY `idx_status` (`status`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


-- 创建用户媒体表(包含文件夹）
CREATE TABLE `tbl_user_file`
(
    `id`             int(11)       NOT NULL AUTO_INCREMENT,
    `pid`            int(11)       NOT NULL DEFAULT '-1' COMMENT '父文件夹id(-1为第一级)',
    `uid`            int(11)       NOT NULL DEFAULT '0',
    `phone`          varchar(128)           DEFAULT '' COMMENT '手机号',
    `file_sha1`      char(40)      NOT NULL DEFAULT '' COMMENT '文件hash',
    `file_sha1_pre`  char(40)      NOT NULL DEFAULT '' COMMENT '预览图文件hash(文件为空，文件夹为最新一张的预览图)',
    `file_name`      varchar(256)  NOT NULL DEFAULT '' COMMENT '文件名',
    `file_addr`      varchar(1024) NOT NULL DEFAULT '' COMMENT '文件存储位置',
    `file_size`      bigint(20)             DEFAULT '0' COMMENT '文件大小',
    `create_at`      datetime               default NOW() COMMENT '创建日期',
    `update_at`      datetime               default NOW() on update current_timestamp() COMMENT '更新日期',
    `delete_at`      datetime               default NOW() COMMENT '删除日期（回收站保留30天功能）',
    `status`         tinyint       NOT NULL DEFAULT '0' COMMENT '状态( 1 可用/ 2 禁用/ -1 已删除等状态)',
    `filetype`       tinyint       NOT NULL DEFAULT '-1' COMMENT '文件类型(文件夹 1/文件-1)',
    `minitype`       char(40)      NOT NULL DEFAULT '' COMMENT '文件类型',
    `ftype`          tinyint       NOT NULL DEFAULT '0' COMMENT '文件状态(0图片/1视频/2音乐/3文档/4压缩包)',
    `video_duration` time          NOT NULL DEFAULT '0' COMMENT '视频时长',
    PRIMARY KEY (`id`),
    KEY `idx_status` (`status`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 5
  DEFAULT CHARSET = utf8mb4;


-- 创建用户表
CREATE TABLE `tbl_user`
(
    `id`              int(11)       NOT NULL AUTO_INCREMENT,
    `user_name`       varchar(64)   NOT NULL DEFAULT '' COMMENT '用户名',
    `user_pwd`        varchar(256)  NOT NULL DEFAULT '' COMMENT '用户encoded密码',
    `email`           varchar(64)            DEFAULT '' COMMENT '邮箱',
    `phone`           varchar(128)           DEFAULT '' COMMENT '手机号',
    `email_validated` tinyint(1)             DEFAULT 0 COMMENT '邮箱是否已验证',
    `phone_validated` tinyint(1)             DEFAULT 0 COMMENT '手机号是否已验证',
    `photo_addr`      varchar(1024) NOT NULL DEFAULT '' COMMENT '用户头像',
    `photo_file_sha1` char(40)      NOT NULL DEFAULT '' COMMENT '用户头像文件hash',
    `space_size`      bigint(20)             DEFAULT '0' COMMENT '用户文件总大小',
    `signup_at`       datetime               DEFAULT CURRENT_TIMESTAMP COMMENT '注册日期',
    `last_active`     datetime               DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后活跃时间戳',
    `profile`         text COMMENT '用户属性',
    `status`          tinyint       NOT NULL DEFAULT '0' COMMENT '账户状态( 1 启用/ 2 禁用/ 3 锁定/ -1 标记删除等)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_phone` (`phone`),
    KEY `idx_status` (`status`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 5
  DEFAULT CHARSET = utf8mb4;

-- 创建用户token表
CREATE TABLE `tbl_user_token`
(
    `id`            int(11)     NOT NULL AUTO_INCREMENT,
    `uid`           int(11)     NOT NULL DEFAULT '0',
    `phone`         varchar(64) NOT NULL DEFAULT '' COMMENT '手机号',
    `user_token`    char(40)    NOT NULL DEFAULT '' COMMENT '用户登录token',
    `expiretimeStr` datetime COMMENT '过期时间',
    `expiretime`    bigint(64) COMMENT '过期时间',
    PRIMARY KEY (`id`),
    KEY `idx_uid` (`uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;



-- 云盘文件表(包含文件夹）为了和媒体文件分开，
CREATE TABLE `tbl_cloud_disk`
(
    `id`             int(11)       NOT NULL AUTO_INCREMENT,
    `pid`            int(11)       NOT NULL DEFAULT '-1' COMMENT '父文件夹id(-1为第一级)',
    `uid`            int(11)       NOT NULL DEFAULT '0',
    `phone`          varchar(128)           DEFAULT '' COMMENT '手机号',
    `file_sha1`      char(40)      NOT NULL DEFAULT '' COMMENT '文件hash',
    `file_sha1_pre`  char(40)      NOT NULL DEFAULT '' COMMENT '预览图文件hash(文件为空，文件夹为最新一张的预览图)',
    `file_name`      varchar(256)  NOT NULL DEFAULT '' COMMENT '文件名',
    `file_addr`      varchar(1024) NOT NULL DEFAULT '' COMMENT '文件存储位置',
    `file_size`      bigint(20)             DEFAULT '0' COMMENT '文件大小',
    `create_at`      datetime               default NOW() COMMENT '创建日期',
    `update_at`      datetime               default NOW() on update current_timestamp() COMMENT '更新日期',
    `delete_at`      datetime               default NOW() COMMENT '删除日期（回收站保留30天功能）',
    `status`         tinyint       NOT NULL DEFAULT '0' COMMENT '状态( 1 可用/ 2 禁用/ -1 已删除等状态)',
    `filetype`       tinyint       NOT NULL DEFAULT '-1' COMMENT '文件类型(文件夹 1/文件-1)',
    `minitype`       char(40)      NOT NULL DEFAULT '' COMMENT '文件类型',
    `ftype`          tinyint       NOT NULL DEFAULT '0' COMMENT '文件状态(0图片/1视频/2音乐/3文档/4压缩包)',
    `video_duration` time          NOT NULL DEFAULT '0' COMMENT '视频时长',
    PRIMARY KEY (`id`),
    KEY `idx_status` (`status`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 5
  DEFAULT CHARSET = utf8mb4;

