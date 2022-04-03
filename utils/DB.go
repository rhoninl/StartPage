package utils

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

func DoLog(id int64, ip string, message string) {
	template := `insert Into startpage.Log Set UserId = ?,UserIp = ?,Message=?`
	result, err := DB().Query(template, id, ip, message)
	if err != nil {
		log.Println("[Mysql]Log make default")
	}
	result.Close()
}

func GetUserIdByUserName(username string) (int64, error) {
	template := `Select UserId From startpage.User Where UserName=? limit 1`
	rows, err := DB().Query(template, username)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		var userId int64
		rows.Scan(&userId)
		return userId, nil
	}
	return 0, fmt.Errorf("account Not Found %s", err.Error())
}

func UpdateSetting(setting UserSetting) bool {
	template := `Update startpage.Setting Set BackGroundUrl = ? Where UserId = ?`
	result, err := DB().Exec(template, setting.ImgUrl, setting.Id)
	if err != nil {
		return false
	}
	rows, _ := result.RowsAffected()
	return rows == 1
}

func InitSetting(id int64) bool {
	template := `Insert into startpage.Setting Set UserId = ?,BackGroundUrl = ''`
	rows, err := DB().Query(template, id)
	if err != nil {
		return false
	}
	defer rows.Close()
	return rows.Next()
}
