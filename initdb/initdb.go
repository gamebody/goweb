package initdb

import (
	"database/sql"
	"fmt"
)

// Db asd
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:123456@tcp(localhost:6033)/blog")
	if err != nil {
		fmt.Println("连接失败了")
		fmt.Println(err)
	}
	fmt.Println("数据库连接成功")
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}
