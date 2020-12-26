CREATE DATABASE IF NOT EXISTS `panenga_db`;

USE `panenga_db`;

CREATE TABLE panels(
  id              INTEGER PRIMARY KEY AUTO_INCREMENT,
  title           VARCHAR(64) NOT NULL,
  panel_image     VARCHAR(255) NOT NULL UNIQUE,
  created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO panels (title, panel_image) VALUES ("三井アウトレットパーク札幌北広島", "https://file24-d.kuku.lu/files/20201226-0812_c12fae45310a831a7a39834c9ecd9dcb.jpg");

INSERT INTO panels (title, panel_image) VALUES ("青森ねぶた祭り", "https://file24-d.kuku.lu/files/20201226-0814_1e834f3d94c0baf3e5e914dd7135181e.jpg");