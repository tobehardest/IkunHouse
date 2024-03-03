CREATE TABLE `video_tag` (
                       `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                       `tag_id` bigint(20) NOT NULL DEFAULT '' COMMENT '标签id',
                       `video_id` bigint(20) NOT NULL DEFAULT '' COMMENT '视频id',
                       `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                       `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
                       PRIMARY KEY (`id`),
                       KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='视频标签表';