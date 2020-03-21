DROP DATABASE IF EXISTS `cln_arch`;
CREATE DATABASE `cln_arch`;
USE `cln_arch`;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `id` VARCHAR(36) PRIMARY KEY,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME DEFAULT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC;

DROP TABLE IF EXISTS `o_auth_states`;
CREATE TABLE `o_auth_states` (
    `id` VARCHAR(36) PRIMARY KEY,
    `state` VARCHAR(255),
    `expiry` DATETIME NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME DEFAULT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC;

DROP TABLE IF EXISTS `o_auth_tokens`;
CREATE TABLE `o_auth_tokens` (
    `user_id` VARCHAR(36) PRIMARY KEY,
    `access_token` VARCHAR(255) NOT NULL,
    `token_type` VARCHAR(255) NOT NULL,
    `refresh_token` VARCHAR(255) NOT NULL,
    `expiry` DATETIME NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at` DATETIME DEFAULT NULL,
    FOREIGN KEY(`user_id`) REFERENCES `users`(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC;
