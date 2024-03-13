

CREATE TABLE `parking_lots` (
                                `id` int NOT NULL AUTO_INCREMENT,
                                `name` varchar(255) NOT NULL,
                                `capacity` int NOT NULL,
                                `active` tinyint(1) NOT NULL,
                                `created_by` int NOT NULL,
                                `create_date` datetime DEFAULT NULL,
                                `update_date` datetime DEFAULT NULL,
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci


//
CREATE TABLE `parking_slots` (
                                 `id` int NOT NULL AUTO_INCREMENT,
                                 `parking_lot_id` int NOT NULL,
                                 `is_occupied` tinyint(1) NOT NULL DEFAULT '0',
                                 `is_maintenance` tinyint(1) NOT NULL DEFAULT '0',
                                 `create_date` datetime DEFAULT CURRENT_TIMESTAMP,
                                 `update_date` datetime DEFAULT NULL,
                                 PRIMARY KEY (`id`),
                                 KEY `parking_lot_id` (`parking_lot_id`),
                                 CONSTRAINT `parking_slots_ibfk_1` FOREIGN KEY (`parking_lot_id`) REFERENCES `parking_lots` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

//
CREATE TABLE `reservations` (
                                `id` int NOT NULL AUTO_INCREMENT,
                                `parking_lot_id` int NOT NULL,
                                `parking_slot_id` int NOT NULL,
                                `vehicle_id` int NOT NULL,
                                `in_time` datetime NOT NULL,
                                `out_time` datetime DEFAULT NULL,
                                `active` tinyint(1) DEFAULT NULL,
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

//
CREATE TABLE `users` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `phone` varchar(45) NOT NULL,
                         `email` varchar(45) DEFAULT NULL,
                         `name` varchar(45) DEFAULT NULL,
                         `type` enum('MANAGER','VEHICLE_OWNER') NOT NULL,
                         `create_date` datetime DEFAULT NULL,
                         `update_date` datetime DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

//
CREATE TABLE `vehicles` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `user_id` int NOT NULL,
                            `registration_no` varchar(45) DEFAULT NULL,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci




