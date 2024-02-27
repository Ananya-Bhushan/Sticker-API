CREATE DATABASE  IF NOT EXISTS `dummy` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `dummy`;
DROP TABLE IF EXISTS `dummy`;
CREATE TABLE `trendingStickers` (
                           `id` int NOT NULL,
                           `Name` varchar(25) NOT NULL,
                           'Priority' int NOT NULL,
                           'StartTime' time(0) NOT NULL,
                           'EndTime' time(0) NOT NULL
                           PRIMARY KEY (`id`)
)
-- LOCK TABLE `trendingStickers` WRITE;
INSERT INTO `trendingStickers` VALUES(1,"animated stickers",57,"9:00:00","17:00:00")
INSERT INTO `trendingStickers` VALUES(2,"friends stickers",58,"8:30:00","18:00:00")

-- UNLOCK TABLE;



