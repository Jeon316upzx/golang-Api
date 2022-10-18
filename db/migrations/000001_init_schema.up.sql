CREATE TABLE `accounts` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `balance` bigint NOT NULL,
  `currency` varchar(255) NOT NULL,
  `owner` int
);

CREATE TABLE `owner` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `created_at` timestamp NOT NULL,
  `owner_details` int NOT NULL
);

CREATE TABLE `owner_details` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `age` int,
  `image` varchar(255),
  `country` varchar(255)
);

CREATE TABLE `transfer` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `amount` bigint NOT NULL,
  `to_account` int NOT NULL,
  `from_account` int NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `transactions` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `type` ENUM ('credit', 'debit') NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `amount` bigint NOT NULL,
  `account` int NOT NULL
);

CREATE INDEX `accounts_index_0` ON `accounts` (`owner`);

CREATE INDEX `owner_index_1` ON `owner` (`owner_details`);

CREATE INDEX `transfer_index_2` ON `transfer` (`to_account`);

CREATE INDEX `transfer_index_3` ON `transfer` (`from_account`);

CREATE INDEX `transfer_index_4` ON `transfer` (`to_account`, `from_account`);

CREATE INDEX `transactions_index_5` ON `transactions` (`account`);

ALTER TABLE `accounts` ADD FOREIGN KEY (`owner`) REFERENCES `owner` (`id`);

ALTER TABLE `owner` ADD FOREIGN KEY (`owner_details`) REFERENCES `owner_details` (`id`);

ALTER TABLE `transfer` ADD FOREIGN KEY (`to_account`) REFERENCES `accounts` (`id`);

ALTER TABLE `transfer` ADD FOREIGN KEY (`from_account`) REFERENCES `accounts` (`id`);

ALTER TABLE `transactions` ADD FOREIGN KEY (`account`) REFERENCES `accounts` (`id`);
