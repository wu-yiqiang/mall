CREATE TABLE `user` (
                        `user_id` varchar(40) NOT NULL,
                        `username` varchar(64) NOT NULL DEFAULT "",
                        `password` varchar(64) NOT NULL DEFAULT "",
                        `email` varchar(64) DEFAULT "",
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT = '角色表';


CREATE TABLE `user_relation_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(40) NOT NULL COMMENT '用户ID',
  `role_id` varchar(40) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`id`),
  KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT = '用户角色关联表';



CREATE TABLE `menu` (
     `menu_id` varchar(40) NOT NULL,
     `name` varchar(400) NOT NULL,
     `code` varchar(400) NOT NULL,
     `path` varchar(400) NOT NULL,
     `serial` int NOT NULL,
     `parent_id` varchar(40),
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `is_deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
     PRIMARY KEY (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT = '菜单表';


CREATE TABLE `role_relation_menu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role_id` varchar(40) NOT NULL COMMENT '角色ID',
  `menu_id` varchar(40) NOT NULL COMMENT '用户ID',
  PRIMARY KEY (`id`),
  KEY `idx_menu_id` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT = '角色菜单关联表';

CREATE TABLE `button` (
     `button_id` varchar(40) NOT NULL,
     `name` varchar(400) NOT NULL,
     `code` varchar(400) NOT NULL,
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `is_deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
     PRIMARY KEY (`button_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT = '按钮表';

CREATE TABLE `menu_relation_button` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `menu_id` varchar(40) NOT NULL COMMENT '菜单ID',
  `button_id` varchar(40) NOT NULL COMMENT '按钮ID',
  PRIMARY KEY (`id`),
  KEY `idx_button_id` (`button_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT = '菜单按钮关联表';
