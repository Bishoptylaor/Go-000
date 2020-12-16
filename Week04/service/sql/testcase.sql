CREATE TABLE `testcase_user` (
 `uid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
 `ct` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
 `ut` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新时间',
 `name` varchar(128) NOT NULL DEFAULT 0 COMMENT '用户名',
 `headimg` varchar(128) NOT NULL DEFAULT 0 COMMENT '头像',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';
