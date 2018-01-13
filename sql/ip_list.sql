CREATE TABLE `ip_list` {
    `id` int NOT NULL AUTO_INCREMENT COMMENT "每条记录的唯一标识",
    `ip` varchar(16) NOT NULL COMMENT "抓取的代理地址",
    `port` int NOT NULL COMMENT "代理地址端口",
    `type` varchar(8) NOT NULL COMMENT "类型(http/https)",
    `from` varchar(16) NOT NULL COMMENT "来源站",
    `capture_time` datetime NOT NULL COMMENT "爬取时间",
    `raw_time` datetime NOT NULL COMMENT "源站爬取时间",
    `region` datetime  COMMENT "地区",
    `last_verify_time` datetime NOT NULL COMMENT "最后验证时间",
    `create_time` datetime NOT NULL COMMENT "创建时间",
    PRIMARY KEY (`id`)
}ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;