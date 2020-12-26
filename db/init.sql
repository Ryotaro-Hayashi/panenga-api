CREATE DATABASE IF NOT EXISTS panenga_db;

USE panenga_db;

CREATE TABLE panels(
  id              INTEGER PRIMARY KEY AUTO_INCREMENT,
  title           VARCHAR(64) NOT NULL,
  panel_image     VARCHAR(255) NOT NULL UNIQUE,
  created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
);

INSERT INTO panels (title, panel_image) VALUES ('三井アウトレットパーク札幌北広島', 'https://file24-d.kuku.lu/files/20201226-0812_c12fae45310a831a7a39834c9ecd9dcb.jpg');

INSERT INTO panels (title, panel_image) VALUES ('青森ねぶた祭り', 'https://file24-d.kuku.lu/files/20201226-0814_1e834f3d94c0baf3e5e914dd7135181e.jpg');

INSERT INTO panels (title, panel_image) VALUES ('函館湯川温泉', 'https://file14-d.kuku.lu/files/20201226-1653_dccf804db713b5a6ee66256fdcffb68b.jpg');

INSERT INTO panels (title, panel_image) VALUES ('函館五稜郭タワー', 'https://file14-d.kuku.lu/files/20201226-1651_537299245553a7d06578ebe54fa18b59.png');

INSERT INTO panels (title, panel_image) VALUES ('北海道新幹線・新函館北斗駅', 'https://file14-d.kuku.lu/files/20201226-1651_988537a094b4e60d128b7ecc64ec7d70.jpg');

INSERT INTO panels (title, panel_image) VALUES ('福岡のどこか', 'https://file14-d.kuku.lu/files/20201226-1650_0f3624c43d357b05a98dc7529db909d7.jpeg');