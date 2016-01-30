CREATE DATABASE `issue_management` DEFAULT CHARACTER SET utf8;

USE `issue_management`;

CREATE TABLE IF NOT EXISTS `issue` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(1024) NOT NULL COMMENT 'タイトル',
  `source` varchar(1024) NOT NULL COMMENT '情報元',
  `detail` text NOT NULL COMMENT '詳細',
  `priority` tinyint(1) NOT NULL DEFAULT '1' COMMENT '対応状況(0:重要, 1:緊急, 2:その他)',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '対応状況(0:完了, 1:未対応)',
  `created` datetime NOT NULL COMMENT '作成日時',
  `updated` datetime NOT NULL COMMENT '更新日時',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='脆弱性情報';

CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(1024) NOT NULL COMMENT 'ユーザ名',
  `password` varchar(1024) NOT NULL COMMENT 'パスワード',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='ユーザー';

--GRANT SELECT, INSERT, UPDATE, DELETE ON issue_management.* TO 'issue_manager'@'localhost' IDENTIFIED BY 'Example_passwd123_';

-- CREATE INDEX issue_emergency_no_complite ON issue(`title`);
