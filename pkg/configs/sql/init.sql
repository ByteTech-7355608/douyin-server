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

DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
                           `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                           `created_at` datetime(3) DEFAULT NULL,
                           `updated_at` datetime(3) DEFAULT NULL,
                           `deleted_at` datetime(3) DEFAULT NULL,
                           `vid` bigint(20) unsigned NOT NULL,
                           `uid` bigint(20) unsigned NOT NULL,
                           `content` longtext COLLATE utf8mb4_bin NOT NULL,
                           PRIMARY KEY (`id`),
                           KEY `idx_comment_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS `relation`;
CREATE TABLE `relation` (
                            `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                            `created_at` datetime(3) DEFAULT NULL,
                            `updated_at` datetime(3) DEFAULT NULL,
                            `deleted_at` datetime(3) DEFAULT NULL,
                            `concerner_id` bigint(20) unsigned NOT NULL,
                            `concerned_id` bigint(20) unsigned NOT NULL,
                            PRIMARY KEY (`id`),
                            KEY `idx_relation_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;


DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                        `created_at` datetime(3) DEFAULT NULL,
                        `updated_at` datetime(3) DEFAULT NULL,
                        `deleted_at` datetime(3) DEFAULT NULL,
                        `username` longtext COLLATE utf8mb4_bin NOT NULL,
                        `password` longtext COLLATE utf8mb4_bin NOT NULL,
                        `follow_count` bigint(20) unsigned DEFAULT '0',
                        `follower_count` bigint(20) unsigned DEFAULT '0',
                        PRIMARY KEY (`id`),
                        KEY `idx_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS `video`;
CREATE TABLE `video` (
                         `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                         `created_at` datetime(3) DEFAULT NULL,
                         `updated_at` datetime(3) DEFAULT NULL,
                         `deleted_at` datetime(3) DEFAULT NULL,
                         `play_url` longtext COLLATE utf8mb4_bin NOT NULL,
                         `cover_url` longtext COLLATE utf8mb4_bin NOT NULL,
                         `favorite_count` bigint(20) unsigned DEFAULT '0',
                         `comment_count` bigint(20) unsigned DEFAULT '0',
                         `title` longtext COLLATE utf8mb4_bin NOT NULL,
                         `uid` bigint(20) unsigned NOT NULL,
                         PRIMARY KEY (`id`),
                         KEY `idx_vedio_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
