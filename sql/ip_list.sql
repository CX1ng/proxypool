CREATE TABLE `ip_list` (
    `ip` varchar(16) NOT NULL COMMENT "抓取的代理地址",
    `port` int NOT NULL COMMENT "代理地址端口",
    `type` varchar(8) NOT NULL COMMENT "类型(http/https)",
    `origin` varchar(16) NOT NULL COMMENT "来源站",
    `raw_time` varchar(32) NOT NULL COMMENT "源站爬取时间",
    `region` varchar(64)  COMMENT "地区",
    `capture_time` datetime NOT NULL COMMENT "爬取时间",
    `last_verify_time` datetime NOT NULL COMMENT "最后验证时间",
    `create_time` datetime NOT NULL COMMENT "创建时间",
    PRIMARY KEY (`ip`,`port`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;