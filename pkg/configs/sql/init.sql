CREATE TABLE `user`
(
    id         int(11) auto_increment primary key,
    username   varchar(31)  not null,
    password   varchar(255) not null,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime
) engine InnoDB
  default charset utf8mb4;