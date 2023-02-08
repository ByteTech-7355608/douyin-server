DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`             bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at`     datetime(3) DEFAULT NULL,
    `updated_at`     datetime(3) DEFAULT NULL,
    `deleted_at`     bigint DEFAULT 0,
    `username`       varchar(20)  NOT NULL,
    `password`       varchar(255) NOT NULL,
    `follow_count`   bigint unsigned DEFAULT '0',
    `follower_count` bigint unsigned DEFAULT '0',
    PRIMARY KEY (`id`),
    KEY              `idx_user_deleted_at` (`deleted_at`),
    UNIQUE (`username`, `deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` bigint DEFAULT 0,
    `uid`        bigint unsigned DEFAULT NULL,
    `to_uid`     bigint unsigned DEFAULT NULL,
    `content`    longtext,
    PRIMARY KEY (`id`),
    KEY          `idx_message_deleted_at` (`deleted_at`),
    KEY          `idx_message_uid` (`uid`),
    KEY          `idx_message_to_uid` (`to_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `like`;
CREATE TABLE `like`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` bigint DEFAULT 0,
    `uid`        bigint unsigned DEFAULT NULL,
    `vid`        bigint unsigned DEFAULT NULL,
    `action`     tinyint(1) DEFAULT '0', -- 0代表未点赞，1代表点赞
    PRIMARY KEY (`id`),
    KEY          `idx_like_deleted_at` (`deleted_at`),
    KEY          `idx_like_uid` (`uid`),
    KEY          `idx_like_vid` (`vid`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`
(
    `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` DATETIME(3) DEFAULT NULL,
    `updated_at` DATETIME(3) DEFAULT NULL,
    `deleted_at` bigint DEFAULT 0,
    `vid`        BIGINT UNSIGNED NOT NULL,
    `uid`        BIGINT UNSIGNED NOT NULL,
    `content`    LONGTEXT COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`id`),
    KEY          `idx_comment_deleted_at` (`deleted_at`),
    KEY          `idx_comment_vid` (`vid`),
    KEY          `idx_comment_uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `relation`;
CREATE TABLE `relation`
(
    `id`           BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at`   DATETIME(3) DEFAULT NULL,
    `updated_at`   DATETIME(3) DEFAULT NULL,
    `deleted_at`   bigint DEFAULT 0,
    `concerner_id` BIGINT UNSIGNED NOT NULL,
    `concerned_id` BIGINT UNSIGNED NOT NULL,
    `action`       tinyint(1) DEFAULT '0',  -- 0代表未关注，1代表关注
    PRIMARY KEY (`id`),
    KEY            `idx_relation_deleted_at` (`deleted_at`),
    KEY            `idx_relation_concerner_id` (`concerner_id`),
    KEY            `idx_relation_concerned_id` (`concerned_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`
(
    `id`             bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at`     datetime(3) DEFAULT NULL,
    `updated_at`     datetime(3) DEFAULT NULL,
    `deleted_at`     bigint DEFAULT 0,
    `play_url`       varchar(255) NOT NULL,
    `cover_url`      varchar(255) NOT NULL,
    `favorite_count` bigint unsigned DEFAULT '0',
    `comment_count`  bigint unsigned DEFAULT '0',
    `title`          varchar(255) NOT NULL,
    `uid`            bigint unsigned NOT NULL,
    PRIMARY KEY (`id`),
    KEY              `idx_video_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
