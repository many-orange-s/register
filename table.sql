CREATE TABLE IF NOT EXISTS %s (
    `id` int(10) NOT NULL AUTO_INCREMENT COMMENT'主键',
    `name` varchar(50) DEFAULT NULL COMMENT'姓名',
    `gender` varchar(50) DEFAULT NULL COMMENT'性别',
    `grade` varchar(50) DEFAULT NULL COMMENT'年级',
    `birth` varchar(100) DEFAULT NULL COMMENT'生日',
    `telephone` varchar(100) DEFAULT NULL COMMENT'电话号码',
    `group_name` varchar(100) DEFAULT NULL COMMENT'分组',
    PRIMARY KEY (`id`)
)