create database hackernews;

use hackernews;

DROP TABLE IF EXISTS user;
CREATE TABLE user (
  id         INT AUTO_INCREMENT NOT NULL,
  username      VARCHAR(20) NOT NULL,
  password     VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);
