package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SearchKigenKanriList(keyWord string) (result string) {
	result = "一覧取得予定: " + keyWord
	var conn *gorm.DB = InitDbConnection()

	defer conn.Close()
	return result
}
