CREATE TABLE `user` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
    `password` varchar(32)  NOT NULL DEFAULT '' COMMENT '密码',
    `avatar` varchar(256) NOT NULL DEFAULT '' COMMENT '头像',
    `mobile` varchar(128) NOT NULL DEFAULT '' COMMENT '手机号',
    `follow_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '关注数',
    `follower_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '粉丝数',
    `work_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '作品数数',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`id`),
    KEY `ix_update_time` (`update_time`),
    UNIQUE KEY `username` (`username`)
    UNIQUE KEY `uk_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
