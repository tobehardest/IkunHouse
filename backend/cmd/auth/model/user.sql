CREATE TABLE `user` (
                        `id` varchar(32) NOT NULL,
                        `username` varchar(32)  NOT NULL,
                        `password` varchar(32)  NOT NULL,
                        `avatar` varchar(255)  DEFAULT '' NOT NULL,
                        `follow_count` bigint(20) NOT NULL DEFAULT '0',
                        `follower_count` bigint(20) NOT NULL DEFAULT '0',
                        `work_count` bigint(20) NOT NULL DEFAULT '0',
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `id` (`id`),
                        UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4

