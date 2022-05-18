-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users
(
    `id`   int(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `name` varchar(255) NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS users;