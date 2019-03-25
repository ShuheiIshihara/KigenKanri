package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB接続情報の構造体
type DbConnection struct {
	Dbms     string `json:"dbms"`
	User     string `json:"user"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
	DbName   string `json:"dbName"`
}

// DB接続
func InitDbConnection() *gorm.DB {

	// 接続情報(json)を読み込む
	raw, err := ioutil.ReadFile("./resource/dbConnection.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var dbConn DbConnection

	json.Unmarshal(raw, &dbConn)

	DBMS := dbConn.Dbms
	USER := dbConn.User
	PASS := dbConn.Password
	PROTOCOL := dbConn.Protocol
	DBNAME := dbConn.DbName

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
	// defer db.Close()
}
