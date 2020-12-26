CREATE DATABASE IF NOT EXISTS panenga_db;

USE panenga_db;

CREATE TABLE panels(
  id              INTEGER PRIMARY KEY AUTO_INCREMENT,
  title           VARCHAR(64) NOT NULL,
  panel_image     VARCHAR(255) NOT NULL UNIQUE,
  created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO panels (title, panel_image) VALUES ("テストデータ", "https://unit42.paloaltonetworks.com/wp-content/uploads/2019/07/golang-hacker.jpg");
