DROP DATABASE IF EXISTS `food_service`;


CREATE DATABASE `food_service`;


USE `food_service`;


CREATE TABLE `restaurant` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT 'restaurant',
  `location` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT "餐厅地点，用一个字节来标识.",
  `introduction` varchar(255),
  `average_price` int UNSIGNED COMMENT "平均价格",
  -- 不知道要不要
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `name` (`name`),
  KEY `location` (`location`)
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4;


CREATE TABLE `food` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `restaurant_id` int UNSIGNED NOT NULL DEFAULT 0,
  `name` varchar(255) NOT NULL DEFAULT 'food',
  `location` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT "餐厅地点，用一个字节来标识.",
  `introduction` varchar(255),
  `ingredient` varchar(255) COMMENT "原料， 食材",
  `price` float,
  -- 不知道要不要
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `name` (`name`),
  KEY `restaurant_id` (`restaurant_id`)
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4;