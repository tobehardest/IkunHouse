CREATE TABLE `t_video` (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
   `uid` varchar(32) NOT NULL COMMENT '用户id',
--    `topic_id` bigint(11) DEFAULT '0' COMMENT '圈子id',
--    `discuss_id` bigint(20) DEFAULT '0' COMMENT '话题id',
--     `type_id` bigint(20) DEFAULT 0 COMMENT '类别id',
   `title` varchar(50) DEFAULT '' COMMENT '标题',
   `content` longtext CHARACTER SET utf8mb4 COMMENT '内容',
   `media` text COMMENT '文件',
   `cover_url` text COMMENT '视频封面',
   `read_count` int(255) DEFAULT '0' COMMENT '浏览量',
--    `post_top` int(1) DEFAULT '0' COMMENT '置顶',
   `type` int(1) DEFAULT '1' COMMENT '帖子类型：1图文2视频3文章',
   `address` varchar(255) DEFAULT NULL COMMENT '地址名称',
   `longitude` double(10,6) DEFAULT '0.000000' COMMENT '经度',
   `latitude` double(10,6) DEFAULT '0.000000' COMMENT '纬度',
   `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
--    `status` int(1) NOT NULL DEFAULT '0' COMMENT '状态0正常1审核',
   PRIMARY KEY (`id`) USING BTREE,
   KEY `uid` (`uid`) USING BTREE
--    KEY `topic_id` (`topic_id`) USING BTREE,
--    KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='帖子表';