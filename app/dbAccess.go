package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"database/sql"
   _ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"

	// "github.com/ShuheiIshihara/KigenKanri/model"
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
func InitDbConnection() *sql.DB {

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
	db, err := sql.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// 一覧取得
func SearchKigenKanriListQuery(db *sql.DB) {

	// var usebyDateInfoList model.Test

	rows, err := db.Query("Select id, GOODS_ID from tests ")
	if err != nil {
		panic(err)
	}
	
	var id int
	var goods_id int
	for rows.Next() {
		err := rows.Scan(&id, &goods_id)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, goods_id)
	}
	
}