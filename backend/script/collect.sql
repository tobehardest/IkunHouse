-- CREATE TABLE `vc_video_collect` (
--      `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
--      `uid` int(11) NOT NULL COMMENT '用户id',
--      `post_id` int(11) NOT NULL COMMENT '视频id',
--      PRIMARY KEY (`id`) USING BTREE
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='收藏表';

CREATE TABLE `collect_record` ( (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `biz_id` varchar(64) NOT NULL DEFAULT '' COMMENT '业务ID',
    `obj_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '收藏对象id',
    `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户ID',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`id`),
    KEY `ix_update_time` (`update_time`),
    UNIQUE KEY `uk_biz_obj_uid` (`biz_id`,`obj_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='收藏记录表';

CREATE TABLE `collect_count` (
      `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
      `biz_id` varchar(64) NOT NULL DEFAULT '' COMMENT '业务ID',
      `obj_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '点赞对象id',
      `collect_num` int(11) NOT NULL DEFAULT '0' COMMENT '收藏数',
      `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
      `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
      PRIMARY KEY (`id`),
      KEY `ix_update_time` (`update_time`),
      UNIQUE KEY `uk_biz_obj` (`biz_id`,`obj_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='收藏计数表';
