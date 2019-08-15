package app

import (
	"fmt"
	
	"database/sql"
   _ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

func SearchKigenKanriList(keyWord string) (result string) {
	result = "一覧取得予定: " + keyWord
	fmt.Println(result)

	var db *sql.DB = InitDbConnection()
	SearchKigenKanriListQuery(db)

	defer db.Close()
	return result
}
