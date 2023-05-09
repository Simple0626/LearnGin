/**
 * @author  simple
 * @date  2023/5/9 19:23
 * @version 1.0
 */

package main

import (
	"LearnGin/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("数据库连接失败")
	}
	/*
		自动迁移模式将保持更新到最新。

		警告：自动迁移仅仅会创建表，缺少列和索引，并且不会改变现有列的类型或删除未使用的列以保护数据。
	*/
	db.AutoMigrate(&model.Person{})
	defer db.Close()
}
