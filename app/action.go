package app

import (
	"fmt"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SearchKigenKanriList(keyWord string) (result string) {
	result = "一覧取得予定: " + keyWord
	fmt.Println(result)

	var db *gorm.DB = InitDbConnection()
	SearchKigenKanriListQuery(db)

	defer db.Close()
	return result
}
