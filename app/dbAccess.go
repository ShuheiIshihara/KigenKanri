package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/ShuheiIshihara/KigenKanri/model"
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
}

// 一覧取得
func SearchKigenKanriListQuery(db *gorm.DB) {

	usebyDateInfoList := []model.TUsebyDateInfo{}

	db.LogMode(true)

	// 全件取得
	if err := db.First(&usebyDateInfoList).Error; err != nil {
		panic(err)
	}

	db.Select("ID").Find(&usebyDateInfoList)
	fmt.Println(usebyDateInfoList)

	// for count, record := range usebyDateInfoList{
	// 	fmt.Println(count)
	// 	fmt.Println("id:" + fmt.Sprint(record.Id))
	// 	// fmt.Println("GoodsId:" + fmt.Sprint(record.GOODS_ID))
	// 	// fmt.Println("UserKbn:" + fmt.Sprint(record.USEBY_KBN))
	// 	// fmt.Println("LimitDate:" + fmt.Sprint(record.LIMIT_DATE))
	// 	}

}