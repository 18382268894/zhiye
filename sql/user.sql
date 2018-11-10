use zhiye;
create table if not exists `users`(
  `id` int(11) unsigned not null auto_increment comment '用户ID',
  `username` varchar(50)  not null comment '用户名',
  `password` varchar(50) not null comment '密码',
  `phone` bigint(11) unsigned null comment '手机号码',
  `email` varchar(50) null comment '电子邮箱',
  `real_name` varchar(50) null default '' comment '真实姓名',
  `create_time` timestamp not null default current_timestamp comment '创建时间',
  `last_ip` varchar(20) null default '' comment '最后一次登录ip',
  `last_time` timestamp null comment '最后一次登录时间',
  `status` tinyint unsigned not null default 0 comment '账号状态，0代表未激活，1代表已激活',
  primary key (`id`),
  unique (`username`),
  unique (`email`),
  unique (`phone`)
)engine = InnoDB charset = utf8mb4 collate = utf8mb4_general_ci row_format = dynamic;
