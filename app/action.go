package app

import (
	"fmt"
	
	"database/sql"
   _ "github.com/go-sql-driver/mysql"
)

// 一覧取得
func SearchKigenKanriList(keyWord string) (result string) {
	result = "一覧取得予定: " + keyWord
	fmt.Println(result)

	var db *sql.DB = InitDbConnection()
	SearchKigenKanriListQuery(db)

	defer db.Close()
	return result
}

// 新規登録
func RegisterData(keyWord string) (result string) {
	return ""
}

// １件更新
func UpdateData(keyWord string) (result string) {
	return ""
}

// 1件削除
func DeleteData(keyWord string) (result string) {
	return ""
}