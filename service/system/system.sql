CREATE TABLE `user` (
                        `user_id` varchar(40) NOT NULL,
                        `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL DEFAULT "",
                        `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL DEFAULT "",
                        `email` varchar(64) COLLATE utf8mb4_general_ci,
                        `gender` tinyint(4) NOT NULL DEFAULT '0',
                        `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`user_id`),
                        UNIQUE KEY `idx_username` (`username`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT = '用户表';


CREATE TABLE `role` (
     `role_id` varchar(40) NOT NULL,
     `name` varchar(400) NOT NULL,
     `code` varchar(400) NOT NULL,
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `is_deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
     PRIMARY KEY (`role_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COMMENT = '角色表';


CREATE TABLE `user_relation_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(32) NOT NULL COMMENT '用户ID',
  `role_id` varchar(32) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`id`),
  KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT = '用户角色关联表';
