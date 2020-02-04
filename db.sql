DROP DATABASE IF EXISTS `food_service`;


CREATE DATABASE `food_service`;


USE food_service;


CREATE TABLE `canteen` (
  `id` tinyint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT 'cantenn',
  `storey` tinyint UNSIGNED NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4;


CREATE TABLE `restaurant` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT 'restaurant',
  `location` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT "餐厅地点，用canteen id来标识.",
  `introduction` varchar(255),
  `average_price` float COMMENT "平均价格",
  `picture_url` varchar(255) COMMENT "照片URL",
  `hot` bigint UNSIGNED COMMENT "热度，用于推荐食堂",
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
  `introduction` varchar(255),
  `ingredient` varchar(255) COMMENT "原料， 食材",
  `price` float COMMENT "价格",
  `picture_url` varchar(255) COMMENT "照片URL",
  `is_special` tinyint UNSIGNED COMMENT "是否为特色推荐",
  `hot` bigint UNSIGNED COMMENT "热度，用于华师必吃",
  -- 不知道要不要
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `name` (`name`),
  KEY `restaurant_id` (`restaurant_id`),
  KEY `is_special` (`is_special`)
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4;


INSERT INTO
  `canteen` (`name`, `storey`)
VALUES
  ("学子", 1),
  ("学子", 2),
  ("东一", 1),
  ("东一", 2),
  ("东二", 1),
  ("东二", 2),
  ("桂香园", 1),
  ("桂香园", 2),
  ("桂香园", 3),
  ("博雅园", 1);


INSERT INTO
  `restaurant` (
    `name`,
    `location`,
    `introduction`,
    `average_price`,
    `hot`
  )
VALUES
  ("巧媳妇山西面馆", 3, "好吃的", 12.0, 1111),
  ("机器人刀削面", 1, "好吃的", 11.0, 2222),
  ("bowser的面馆", 1, "一碗阳春面", 8.0, 3333);


INSERT INTO
  `food` (
    `name`,
    `restaurant_id`,
    `price`,
    `is_special`,
    `hot`,
    `introduction`,
    `ingredient`
  )
VALUES
  ("油泼面", 1, 12.0, 1, 111, "陕北风味", "面，辣子"),
  ("油泼辣子", 1, 1.0, 1, 1111, "陕北必备", "油，辣子");