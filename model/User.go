/**
 * @author  simple
 * @date  2023/5/7 18:47
 * @version 1.0
 */

package model

import "time"

type User struct {
	UserName string `form:"name"`
	Password string
	Hobby    []string
}

type Person struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
