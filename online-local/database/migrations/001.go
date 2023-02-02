package migrations

import "database/sql"

func Init(db *sql.DB) {
	var err error
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `user` (`id` binary(16) NOT NULL,`login` varchar(255) NOT NULL,`password_hash` varchar(255) NOT NULL,`name` varchar(255) NOT NULL,`enabled` tinyint(1) NOT NULL,`created_at` datetime DEFAULT NULL,`updated_at` datetime DEFAULT NULL,`_auth_fail_count` int(11) NOT NULL,`_last_auth_fail_date` datetime NOT NULL,PRIMARY KEY (`id`),UNIQUE KEY `login` (`login`)) ENGINE=InnoDB DEFAULT CHARSET=utf8")
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `competitor` (`id` binary(16) NOT NULL,`identification` binary(16) NOT NULL,`name` varchar(255) NOT NULL,`organization` varchar(255) DEFAULT NULL,`region` varchar(255) DEFAULT NULL,`country` varchar(255) DEFAULT NULL,`birthday` date DEFAULT NULL,PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8")
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `race` (`id` binary(16) NOT NULL,`owner_id` binary(16) DEFAULT NULL,`title` varchar(1024) NOT NULL,`description` varchar(1024) NOT NULL,PRIMARY KEY (`id`),KEY `FK_race_user` (`owner_id`),CONSTRAINT `FK_race_user` FOREIGN KEY (`owner_id`) REFERENCES `user` (`id`) ON DELETE SET NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8")
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `group` (`id` binary(16) NOT NULL,`race_id` binary(16) NOT NULL,`name` varchar(255) NOT NULL,PRIMARY KEY (`id`),KEY `FK_group_race` (`race_id`),CONSTRAINT `FK_group_race` FOREIGN KEY (`race_id`) REFERENCES `race` (`id`) ON DELETE CASCADE) ENGINE=InnoDB DEFAULT CHARSET=utf8")
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `course` (`id` binary(16) NOT NULL,`race_id` binary(16) NOT NULL,`name` varchar(255) NOT NULL,`chunks` json NOT NULL,PRIMARY KEY (`id`),KEY `FK_course_race` (`race_id`),CONSTRAINT `FK_course_race` FOREIGN KEY (`race_id`) REFERENCES `race` (`id`) ON DELETE CASCADE) ENGINE=InnoDB DEFAULT CHARSET=utf8")
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `result` (`id` binary(16) NOT NULL,`bib` varchar(50) NOT NULL,`competitor_id` binary(16) DEFAULT NULL,`race_id` binary(16) NOT NULL,`group_id` binary(16) DEFAULT NULL,`course_id` binary(16) DEFAULT NULL,`qual` varchar(50) DEFAULT NULL,`chunks` json NOT NULL,PRIMARY KEY (`id`),KEY `FK_result_race` (`race_id`),KEY `FK_result_competitor` (`competitor_id`),KEY `FK_result_course` (`course_id`),KEY `FK_result_group` (`group_id`),CONSTRAINT `FK_result_competitor` FOREIGN KEY (`competitor_id`) REFERENCES `competitor` (`id`) ON DELETE SET NULL,CONSTRAINT `FK_result_course` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`) ON DELETE SET NULL,CONSTRAINT `FK_result_group` FOREIGN KEY (`group_id`) REFERENCES `group` (`id`) ON DELETE SET NULL,CONSTRAINT `FK_result_race` FOREIGN KEY (`race_id`) REFERENCES `race` (`id`) ON DELETE CASCADE) ENGINE=InnoDB DEFAULT CHARSET=utf8")
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `chip` (`name` varchar(50) NOT NULL, `identification` varchar(50) NOT NULL, `race_id` binary(16) NOT NULL, `result_id` binary(16) NOT NULL, KEY `FK_chip_race` (`race_id`), KEY `FK_chip_result` (`result_id`), CONSTRAINT `FK_chip_race` FOREIGN KEY (`race_id`) REFERENCES `race` (`id`) ON DELETE CASCADE, CONSTRAINT `FK_chip_result` FOREIGN KEY (`result_id`) REFERENCES `result` (`id`) ON DELETE CASCADE) ENGINE=InnoDB DEFAULT CHARSET=utf8")
	if err != nil {
		panic(err)
	}
}
