DROP DATABASE IF EXISTS `user`;
CREATE DATABASE `user`;
USE `user`;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `id` VARCHAR(36) PRIMARY KEY,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC;

DROP TABLE IF EXISTS `user_states`;
CREATE TABLE `user_states` (
    `user_id` VARCHAR(36) PRIMARY KEY,
    `session_id` VARCHAR(255) NOT NULL,
    `state` VARCHAR(255),
    `expiry` TIMESTAMP NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(`user_id`) REFERENCES `user`(`id`),
    FOREIGN KEY(`github_token_id`) REFERENCES `github_tokens`(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC;

DROP TABLE IF EXISTS `user_tokens`;
CREATE TABLE `user_token` (
    `user_id` VARCHAR(36) PRIMARY KEY,
    `token` VARCHAR(255),
    `expiry` TIMESTAMP NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(`user_id`) REFERENCES `user`(`id`),
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC;

DROP TABLE IF EXISTS `github_tokens`;
CREATE TABLE `github_tokens` (
    `user_id` VARCHAR(36) PRIMARY KEY,
    `token` VARCHAR(255) NOT NULL,
    `type` VARCHAR(255) NOT NULL,
    `refresh_token` VARCHAR(255) NOT NULL,
    `expiry` TIMESTAMP NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(`user_id`) REFERENCES `user`(`id`),
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC;
