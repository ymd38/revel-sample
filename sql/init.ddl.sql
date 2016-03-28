CREATE DATABASE `issue_management` DEFAULT CHARACTER SET utf8;

USE `issue_management`;

CREATE TABLE IF NOT EXISTS `issue` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(1024) NOT NULL COMMENT 'タイトル',
  `source` varchar(1024) NOT NULL COMMENT '情報元',
  `detail` text NOT NULL COMMENT '詳細',
  `priority` tinyint(1) NOT NULL DEFAULT '1' COMMENT '対応状況(0:重要, 1:緊急, 2:その他)',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '対応状況(0:完了, 1:未完了)',
  `limit` int NOT NULL  COMMENT '期限日(UNIX_TIMESTAMP)',
  `created` int NOT NULL COMMENT '作成日時(UNIX_TIMESTAMP)',
  `updated` int NOT NULL COMMENT '更新日時(UNIX_TIMESTAMP)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='脆弱性情報';


CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `mailaddr` varchar(1024) NOT NULL COMMENT 'メールアドレス',
  `password` varchar(1024) NOT NULL COMMENT 'パスワード',
  `created` int NOT NULL COMMENT '作成日時(UNIX_TIMESTAMP)',
  `updated` int NOT NULL COMMENT '更新日時(UNIX_TIMESTAMP)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='ユーザー';

CREATE TABLE IF NOT EXISTS `service` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(1024) NOT NULL COMMENT 'サービス名',
  `start` int DEFAULT null COMMENT 'サービス開始日(UNIX_TIMESTAMP)',
  `end` int DEFAULT null COMMENT 'サービス終了日(UNIX_TIMESTAMP)',
  `created` int NOT NULL COMMENT '作成日時(UNIX_TIMESTAMP)',
  `updated` int NOT NULL COMMENT '更新日時(UNIX_TIMESTAMP)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='サービス';

CREATE TABLE IF NOT EXISTS `service_member` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `serviceid` int(11) unsigned NOT NULL COMMENT 'サービスID',
  `userid` int(11) unsigned NOT NULL COMMENT 'ユーザID',
  `created` int NOT NULL COMMENT '作成日時(UNIX_TIMESTAMP)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='サービス担当';

CREATE TABLE IF NOT EXISTS `service_issue` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `serviceid` int(11) unsigned NOT NULL COMMENT 'サービスID',
  `issueid` int(11) unsigned NOT NULL COMMENT 'ユーザID',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '対応状況(0:完了, 1:未確認, 2:対応不要、3:確認中、4:対応予定)',
  `reflectdate` int DEFAULT 0 COMMENT '対応予定日(UNIX_TIMESTAMP)',
  `memo` text DEFAULT NULL COMMENT 'メモ',
  `created` int NOT NULL COMMENT '作成日時(UNIX_TIMESTAMP)',
  `updated` int NOT NULL COMMENT '更新日時(UNIX_TIMESTAMP)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='サービス対応状況';

--index


DROP TABLE issue;
DROP TABLE user;
DROP TABLE service;
DROP TABLE service_member;
DROP TABLE service_issue;



--GRANT SELECT, INSERT, UPDATE, DELETE ON issue_management.* TO 'issue_manager'@'localhost' IDENTIFIED BY 'Example_passwd123_';

-- CREATE INDEX issue_emergency_no_complite ON issue(`title`);


--CREATE VIEW service_issue_view AS
--SELECT s.serviceid ServiceID, i.id IssueId, i.title IssueTitle, i.priority IssuePriority, s.status StatusCode, s.reflectdate ReflectDate
--FROM service_issue s INNER JOIN issue i ON s.issueid = i.id;
