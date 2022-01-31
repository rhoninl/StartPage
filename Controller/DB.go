package Controller

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func init() {
	dsn := "root:dmutreehole@tcp(www.wonend.cn:3306)/startpage"
	dbs, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panic(err.Error())
	}
	db = dbs
	fmt.Println(db.Ping())
}

func DB() *sql.DB {
	return db
}
