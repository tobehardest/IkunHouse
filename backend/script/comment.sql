CREATE TABLE `fd_comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `pid` int(11) DEFAULT '0' COMMENT '父级id',
  `type` int(1) NOT NULL DEFAULT '1' COMMENT '评论类型:1视频, 2评论.',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '评论作者ID',
  #`to_uid` int(11) DEFAULT '0' COMMENT '被回复用户ID',
  #`post_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '评论帖子ID',
  `content` varchar(255) NOT NULL DEFAULT '' COMMENT '评论内容',
  `status` tinyint(4) DEFAULT '1' COMMENT '评论状态',
  `create_time` datetime DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `uid` (`uid`) USING BTREE,
  KEY `pid` (`pid`) USING BTREE,
  KEY `to_uid` (`to_uid`) USING BTREE,
  KEY `post_id` (`post_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='用户评论表';