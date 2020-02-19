# golang_mysql
环境：Mysql版本: 5.7.28-log Homebrew;Golang版本 go1.13.8 darwin/amd64
/**
使用说明：
本项目主要为了完善一个通用的golang-mysql操作库，网上查了一些资料，发现不是很通用，结合网上的资料整理分享给大家使用

1.0可以支持多库


一、配置数据库连接
打开项目文件 src/db/db.go 在initDbPoolConfig()中配置自己的数据库连接
二、运行main 可以看到创建测试表，增删改查操作

可以根据自己的项目，扩展自己的需求

环境
mysql 5.7
golang 1.3
IDE:GoLand
*/