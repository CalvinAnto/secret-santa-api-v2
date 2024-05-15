CREATE TABLE `game` (
  `id` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `size` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
CREATE TABLE `player` (
  `id` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `wishlist` text COLLATE utf8mb4_general_ci,
  `taken` tinyint(1) DEFAULT NULL,
  `receiver_id` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `game_id` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `receiver` (`receiver_id`),
  CONSTRAINT `receiver` FOREIGN KEY (`receiver_id`) REFERENCES `player` (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;