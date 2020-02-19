package main

import (
	"fmt"
	"models"
)

func createTableDemo() {
	demoSql := "CREATE TABLE `demo` ("
	demoSql += "  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,"
	demoSql += "  `name` varchar(255) DEFAULT '',"
	demoSql += "  `age` int(10) DEFAULT '0',"
	demoSql += "  PRIMARY KEY (`id`),"
	demoSql += "  KEY `s1` (`name`) USING BTREE"
	demoSql += ") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"
	if models.InitModel("demo", "demo").CheckTableIsExists() == false {
		models.InitModel("demo", "demo").ExecuteSql(demoSql)
		fmt.Println("测试表创建成功")
	} else {
		fmt.Println("测试表已经存在")
	}
}

func insertDemo() {
	createMap := make(map[string]interface{})
	createMap["name"] = "测试姓名"
	createMap["age"] = "11"
	models.InitModel("demo", "demo").Insert(createMap)
}

func updateDemo() {
	likeMap := make(map[string]interface{})
	whereMap := make(map[string]interface{})
	updateDatas := make(map[string]interface{})

	whereMap["name"] = "测试姓名"

	updateDatas["name"] = "修改过的姓名"

	models.InitModel("demo", "demo").UpdateByWhereLike(updateDatas, whereMap, likeMap)
}

func getOneDemo() {
	likeMap := make(map[string]interface{})
	whereMap := make(map[string]interface{})

	whereMap["name"] = "测试姓名"
	row := models.InitModel("demo", "demo").GetOneByWhereLike("*", whereMap, likeMap)
	for db_field, db_value := range row {
		fmt.Println(db_field, db_value)
	}
}

func getAllDemo() {
	likeMap := make(map[string]interface{})
	whereMap := make(map[string]interface{})

	whereMap["name"] = "测试姓名"
	rows := models.InitModel("demo", "demo").GetAllByWhereLike("*", whereMap, likeMap)
	for idx, row := range rows {
		fmt.Println(idx, row)
	}
}

func delDemo() {
	likeMap := make(map[string]interface{})
	whereMap := make(map[string]interface{})
	whereMap["name"] = "修改过的姓名"
	models.InitModel("demo", "demo").Delete(whereMap, likeMap)
}

func main() {
	createTableDemo()
	insertDemo()
	getOneDemo()
	getAllDemo()
	updateDemo()
	delDemo()
}
