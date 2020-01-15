DROP DATABASE IF EXISTS `food_service`;


CREATE DATABASE `food_service`;


USE `food_service`;


CREATE TABLE `restaurant` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT 'restaurant',
  `location` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT "餐厅地点，用一个字节来标识.",
  `introduction` varchar(255),
  `sales_volumn` int UNSIGNED COMMENT "月销量",
  -- 不知道要不要
  `createdAt` timestamp NULL DEFAULT NULL,
  `updatedAt` timestamp NULL DEFAULT NULL,
  `deletedAt` timestamp NULL DEFAULT NULL,
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
  `createdAt` timestamp NULL DEFAULT NULL,
  `updatedAt` timestamp NULL DEFAULT NULL,
  `deletedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `name` (`name`),
  KEY `restaurant_id` (`restaurant_id`)
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4;