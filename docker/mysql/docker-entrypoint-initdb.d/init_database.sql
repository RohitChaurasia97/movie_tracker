CREATE DATABASE IF NOT EXISTS `local` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
use local;
CREATE TABLE `Movie` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Title` varchar(100) DEFAULT NULL,
  `Completed` tinyint(1) DEFAULT NULL,
  `CreatedAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `UpdatedAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SELECT * FROM local.Movie;

