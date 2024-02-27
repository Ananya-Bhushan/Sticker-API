DROP TABLE IF EXISTS `dummy`;

CREATE TABLE `trendingstickers` (
                           id int NOT NULL,
                           Name varchar(25) NOT NULL,
                           Priority int NOT NULL,
                           StartTime time NOT NULL,
                           EndTime time NOT NULL,
                           PRIMARY KEY (id)
)
