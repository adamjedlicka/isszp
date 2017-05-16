SET FOREIGN_KEY_CHECKS=0 ;

DROP TABLE IF EXISTS `users` CASCADE
;

DROP TABLE IF EXISTS `time_records` CASCADE
;

DROP TABLE IF EXISTS `tasks` CASCADE
;

DROP TABLE IF EXISTS `rel_files` CASCADE
;

DROP TABLE IF EXISTS `projects` CASCADE
;

DROP TABLE IF EXISTS `permissions` CASCADE
;

DROP TABLE IF EXISTS `firms` CASCADE
;

DROP TABLE IF EXISTS `files` CASCADE
;

DROP TABLE IF EXISTS `comments` CASCADE
;

CREATE TABLE `users`
(
	`id` CHAR(36) NOT NULL,
	`user_name` VARCHAR(50) NOT NULL,
	`password` CHAR(60) NOT NULL,
	`first_name` VARCHAR(50) 	 NULL,
	`last_name` VARCHAR(50) 	 NULL,
	`permission` BIGINT NOT NULL,
	`deleted_at` DATETIME(0) 	 NULL,
	CONSTRAINT `PK_Users` PRIMARY KEY (`id`)
)

;

CREATE TABLE `time_records`
(
	`id` CHAR(36) NOT NULL,
	`description` TEXT 	 NULL,
	`time_in_ms` VARCHAR(50) NULL,
	`date` DATE NOT NULL,
	`start` TIME(0) NOT NULL,
	`end` TIME(0) 	 NULL,
	`user_id` CHAR(36) NOT NULL,
	`task_id` CHAR(36) NOT NULL,
	`deleted_at` DATETIME(0) 	 NULL,
	CONSTRAINT `PK_TimeRecords` PRIMARY KEY (`id`)
)

;

CREATE TABLE `tasks`
(
	`id` CHAR(36) NOT NULL,
	`name` VARCHAR(50) NOT NULL,
	`description` TEXT 	 NULL,
	`state` ENUM('free', 'active', 'revision', 'success', 'fail') NOT NULL,
	`start_date` DATE NOT NULL,
	`plan_end_date` DATE 	 NULL,
	`end_date` DATE 	 NULL,
	`maintainer_id` CHAR(36) NOT NULL,
	`worker_id` CHAR(36) 	 NULL,
	`project_id` CHAR(36) NOT NULL,
	`deleted_at` DATETIME(0) 	 NULL,
	CONSTRAINT `PK_Tasks` PRIMARY KEY (`id`)
)

;

CREATE TABLE `rel_files`
(
	`id` CHAR(36) NOT NULL,
	`file_id` CHAR(36) NOT NULL,
	`task_id` CHAR(36) 	 NULL,
	`project_id` CHAR(36) 	 NULL,
	`user_id` CHAR(36) NOT NULL,
	CONSTRAINT `PK_Rel_Files` PRIMARY KEY (`id`)
)

;

CREATE TABLE `projects`
(
	`id` CHAR(36) NOT NULL,
	`name` VARCHAR(50) NOT NULL,
	`code` VARCHAR(50) NOT NULL,
	`description` TEXT 	 NULL,
	`start_date` DATE NOT NULL,
	`plan_end_date` DATE 	 NULL,
	`end_date` DATE 	 NULL,
	`maintainer_id` CHAR(36) NOT NULL,
	`firm_id` CHAR(36) NOT NULL,
	`deleted_at` DATETIME(0) 	 NULL,
	CONSTRAINT `PK_Projects` PRIMARY KEY (`id`)
)

;

CREATE TABLE `firms`
(
	`id` CHAR(36) NOT NULL,
	`name` VARCHAR(50) NOT NULL,
	`description` TEXT 	 NULL,
	`email` VARCHAR(30) 	 NULL,
	`tel_number` VARCHAR(15) 	 NULL,
	`deleted_at` DATETIME(0) 	 NULL,
	CONSTRAINT `PK_Firms` PRIMARY KEY (`id`)
)

;

CREATE TABLE `files`
(
	`id` CHAR(36) NOT NULL,
	`name` VARCHAR(50) NOT NULL,
	`upload_date_time` DATETIME(0) NOT NULL,
	`data` BLOB NOT NULL,
	`deleted_at` DATETIME(0) 	 NULL,
	CONSTRAINT `PK_Files` PRIMARY KEY (`id`)
)

;

CREATE TABLE `comments`
(
	`id` CHAR(36) NOT NULL,
	`text` TEXT NOT NULL,
	`post_date_time` DATETIME(0) NOT NULL,
	`user_id` CHAR(36) NOT NULL,
	`task_id` CHAR(36) NOT NULL,
	`deleted_at` DATETIME(0) 	 NULL,
	CONSTRAINT `PK_Comments` PRIMARY KEY (`id`)
)

;

ALTER TABLE `time_records`
 ADD CONSTRAINT `FK_TimeRecords_Tasks`
	FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`) ON DELETE Restrict ON UPDATE Restrict
;

ALTER TABLE `time_records`
 ADD CONSTRAINT `FK_TimeRecords_Users`
	FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE Restrict ON UPDATE Restrict
;

ALTER TABLE `tasks`
 ADD CONSTRAINT `FK_Tasks_Projects`
	FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`) ON DELETE Restrict ON UPDATE Restrict
;

ALTER TABLE `tasks`
 ADD CONSTRAINT `FK_Tasks_Users_Maintainer`
	FOREIGN KEY (`maintainer_id`) REFERENCES `users` (`id`) ON DELETE Restrict ON UPDATE Restrict
;

ALTER TABLE `tasks`
 ADD CONSTRAINT `FK_Tasks_Users_Worker`
	FOREIGN KEY (`worker_id`) REFERENCES `users` (`id`) ON DELETE Restrict ON UPDATE Restrict
;

ALTER TABLE `rel_files`
 ADD CONSTRAINT `FK_Rel_Files_Files`
	FOREIGN KEY (`file_id`) REFERENCES `files` (`id`) ON DELETE Restrict ON UPDATE Restrict
;

ALTER TABLE `rel_files`
 ADD CONSTRAINT `FK_Rel_Files_Projects`
	FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`) ON DELETE Restrict ON UPDATE Restrict
;

ALTER TABLE `rel_files`
 ADD CONSTRAINT `FK_Rel_Files_Tasks`
	FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`) ON DELETE Restrict ON UPDATE Restrict
;

ALTER TABLE `rel_files`
 ADD CONSTRAINT `FK_Rel_Files_Users`
	FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE Restrict ON UPDATE Restrict
;

ALTER TABLE `projects`
 ADD CONSTRAINT `FK_Projects_Firms`
	FOREIGN KEY (`firm_id`) REFERENCES `firms` (`id`) ON DELETE Restrict ON UPDATE Restrict
;

ALTER TABLE `projects`
 ADD CONSTRAINT `FK_Projects_Users`
	FOREIGN KEY (`maintainer_id`) REFERENCES `users` (`id`) ON DELETE Restrict ON UPDATE Restrict
;

ALTER TABLE `comments`
 ADD CONSTRAINT `FK_Comments_Tasks`
	FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`) ON DELETE Restrict ON UPDATE Restrict
;

ALTER TABLE `comments`
 ADD CONSTRAINT `FK_Comments_Users`
	FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE Restrict ON UPDATE Restrict
;

SET FOREIGN_KEY_CHECKS=1 ;
