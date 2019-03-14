package app

import (
	"github.com/gin-gonic/gin"
)

func Route() {

	router := gin.Default()
	router.GET("/", Test)                  // 起動確認用
	router.GET("/getList", GetList)        // 一覧取得
	router.POST("/addData", AddData)       // 新規登録
	router.POST("/updateData", UpdateData) // 1件更新
	router.POST("/deleteData", DeleteData) // 1件削除

	// アプリ起動
	router.Run(":8081")
}
