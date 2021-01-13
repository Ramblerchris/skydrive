-- 创建文件表
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


-- 创建用户文件表(包含文件夹）
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

-- 视频类型表 ：类型表:剧情，恐怖，纪实，励志，情感，战争，青春，动作，历史，喜剧，人文，搞笑，古装
CREATE TABLE `tbl_film_type`
(
    `film_type_id`        int(11)     NOT NULL AUTO_INCREMENT,
    `type_id`   int(11)     NOT NULL COMMENT '类型id',
    `type_name` varchar(64) NOT NULL COMMENT '视频类型表名称',
    `status`    tinyint     NOT NULL COMMENT '状态( 1 启用/ 2 禁用/ -1 标记删除等)',
    PRIMARY KEY (`film_type_id`),
    KEY `type_id` (`type_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


-- 视频分类表 ：电影，电视剧，纪录片，综艺
CREATE TABLE `tbl_film_classfy`
(
    `film_classfy_id`         int(11)     NOT NULL AUTO_INCREMENT,
    `class_id`   int(11)     NOT NULL COMMENT '分类id',
    `class_name` varchar(64) NOT NULL COMMENT '视频分类表',
    `status`     tinyint     NOT NULL COMMENT '状态( 1 启用/ 2 禁用/ -1 标记删除等)',
    PRIMARY KEY (`film_classfy_id`),
    KEY `class_id` (`class_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


-- 剧基本信息表:  名称，上映时间，简介，封面，导演，主演，国家，评分，语言，分类（分类表），类型（类型表），查询关键字，搜索关键字，图片资料，电影路径
CREATE TABLE `tbl_film_info`
(
    `film_info_id`         int(11)       NOT NULL AUTO_INCREMENT,
    `film_name`            varchar(256)  NOT NULL DEFAULT '' COMMENT '视频名称',
    `brief_introduction`   varchar(1024) NOT NULL DEFAULT '' COMMENT '简介',
    `film_cover`           varchar(256)  NOT NULL DEFAULT '' COMMENT '剧封面',
    `film_classfy_id`      int(11)       NOT NULL DEFAULT '0' COMMENT '分类id',
    `film_type_id`         int(11)       NOT NULL DEFAULT '0' COMMENT '类型id',
    `director_name`        varchar(256)  NOT NULL DEFAULT '' COMMENT '导演',
    `country`              varchar(50)   NOT NULL DEFAULT '' COMMENT '国家',
    `language`             varchar(50)   NOT NULL DEFAULT '' COMMENT '语言',
    `act_the_leading_role` varchar(256)  NOT NULL DEFAULT '' COMMENT '主演',
    `search_key`           varchar(256)  NOT NULL DEFAULT '' COMMENT '搜索关键字',
    `score`                int(11)       NOT NULL DEFAULT '0' COMMENT '评分',
    `show_at`              datetime               default NOW() COMMENT '上映时间',
    `create_at`            datetime               default NOW() COMMENT '添加时间',
    `status`               tinyint       NOT NULL COMMENT '状态( 1 可以查询/ 2不可查询/ -1 标记删除等)',
    `film_duration`        time          NOT NULL DEFAULT '0' COMMENT '视频时长',
    `photo_addr`           varchar(1024) NOT NULL DEFAULT '' COMMENT '图片资料，多张用;分割',
    PRIMARY KEY (`film_info_id`),
    KEY `idx_status` (`status`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


-- 剧播放文件列表
CREATE TABLE `tbl_video_playlist`
(
    `video_id`                 int(11)       NOT NULL AUTO_INCREMENT,
    `film_info_id`             int(11)       NOT NULL DEFAULT '0' COMMENT '剧基本信息表id',
    `video_name`               varchar(256)  NOT NULL DEFAULT '' COMMENT '单个视频名称',
    `video_brief_introduction` varchar(1024) NOT NULL DEFAULT '' COMMENT '单个视频简介',
    `video_cover`              varchar(256)  NOT NULL DEFAULT '' COMMENT '单个视频剧封面',
    `video_duration`           time          NOT NULL DEFAULT '0' COMMENT '单个视频时长',
    `score`                    int(11)       NOT NULL DEFAULT '0' COMMENT '单个视频评分',
    `create_at`                datetime               default NOW() COMMENT '添加时间',
    `status`                   tinyint       NOT NULL COMMENT '状态( 1 可以查询/ 2不可查询/ -1 标记删除等)',
    `file_addr`                varchar(1024) NOT NULL DEFAULT '' COMMENT '文件存储位置',
    PRIMARY KEY (`video_id`),
    KEY `idx_status` (`status`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;





