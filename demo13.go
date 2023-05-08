/**
 * @author  simple
 * @date  2023/5/8 11:52
 * @version 1.0
 */

package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

func main() {
	//username:password@protocol(address)/dbname?param=value
	dsn := "root:root@tcp(127.0.0.1:3306)/go_test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		println(err)
	}
	db.Ping()
	println("连接成功")
	err = initDB(dsn)
	if err != nil {
		println(err)
	}

	router := gin.Default()
	router.GET("/insert", insertHandler)
	router.GET("/delete", deleteHandler)
	router.GET("/update", updateHandler)
	router.GET("/select", selectHandler)
	router.Run(":8000")

}
func selectHandler(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	course := Query(dbs, id)
	ctx.JSON(200, course)

}

func updateHandler(ctx *gin.Context) {
	name := ctx.Query("name")
	id, _ := strconv.Atoi(ctx.Query("id"))
	age, _ := strconv.Atoi(ctx.Query("age"))
	course := Course{
		Cid:   id,
		Cname: name,
		Cage:  age,
	}
	update(dbs, course)
	ctx.JSON(200, course)
}

func deleteHandler(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	delete(dbs, id)
	ctx.JSON(200, id)
}

func insertHandler(ctx *gin.Context) {
	name := ctx.Query("name")
	id, _ := strconv.Atoi(ctx.Query("id"))
	age, _ := strconv.Atoi(ctx.Query("age"))

	course := Course{
		Cid:   id,
		Cname: name,
		Cage:  age,
	}
	insert(dbs, course)
	ctx.JSON(200, course)
}

// 封装一个方法:
var dbs *sql.DB //默认给一个数据库连接池

func initDB(str string) (err error) {
	dbs, err = sql.Open("mysql", str)
	if err != nil {
		println(err)
		return err
	}
	//最大连接数
	dbs.SetMaxOpenConns(20)
	//最大空闲数
	dbs.SetMaxIdleConns(9)
	return nil
}

// 增
type Course struct {
	Cid   int
	Cname string
	Cage  int
}

func insert(db *sql.DB, course Course) {
	//开启事务
	db.Begin()
	//预编译的方式
	stmt, _ := db.Prepare("insert into user (id,name,age) values (?,?,?)")
	result, _ := stmt.Exec(course.Cid, course.Cname, course.Cage)
	println(result)
	stmt.Close()

}

func delete(db *sql.DB, id int) {
	//开启事务
	db.Begin()
	//预编译的方式
	stmt, _ := db.Prepare("delete from user where	id = ?")
	result, _ := stmt.Exec(id)
	println(result)
	stmt.Close()

}

func update(db *sql.DB, course Course) {
	//开启事务
	db.Begin()
	//预编译的方式
	stmt, _ := db.Prepare("update user set id=? , name = ?, age=? where id =?")
	result, _ := stmt.Exec(course.Cid, course.Cname, course.Cage, course.Cid)
	println(result)
	stmt.Close()

}

func Query(db *sql.DB, id int) (course Course) {
	var (
		name = ""
		age  = 0
	)
	row := db.QueryRow("select id,name,age from user where id = ?", id).Scan(&id, &name, &age)
	if row != nil {
		log.Println(row)
	} else {
		fmt.Printf("user name :%s , age :%d", name, age)
	}
	course = Course{
		Cid:   id,
		Cname: name,
		Cage:  age,
	}
	return course
}
