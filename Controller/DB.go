package Controller

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var db *sql.DB
var rdb *redis.Client

func init() {
	//连接MySql数据库
	dsn := "StartPage:liyuqi521@tcp(www.wonend.cn:3306)/startpage"
	dbs, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panic(err.Error())
	}
	db = dbs
	fmt.Println("Mysql Connect Success!")
	//连接Redis数据库
	rdb = redis.NewClient(&redis.Options{
		Addr:     "www.wonend.cn:6379",
		Password: "",
		DB:       0,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Redis Connect Success!")
}

//DB Return Mysql DataBase Pointer To Control
func DB() *sql.DB {
	return db
}

//RDB Return Redis Database Client To Control
func RDB() *redis.Client {
	return rdb
}

func DoLog(id int, ip string, message string) {
	template := `insert Into Log Set UserId = ?,UserIp = ?,Message=?"`
	_, err := DB().Query(template, id, ip, message)
	if err != nil {
		log.Println("[Mysql]Log make default")
	}
}
