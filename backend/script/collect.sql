CREATE TABLE `vc_video_collect` (
     `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
     `uid` int(11) NOT NULL COMMENT '用户id',
     `post_id` int(11) NOT NULL COMMENT '视频id',
     PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='收藏表';