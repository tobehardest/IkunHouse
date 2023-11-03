CREATE TABLE `favorite` (
                        `id` varchar(32) NOT NULL,
                        `user_id` bigint(20) DEFAULT NULL,
                        `video_id` bigint(20) DEFAULT NULL,
                        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `id` (`id`),
                        UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
