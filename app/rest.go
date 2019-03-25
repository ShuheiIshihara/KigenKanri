package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// アクセス確認
func Test(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

// 一覧取得
func GetList(c *gin.Context) {
	name := c.Query("name")
	fmt.Printf(name)
	message := SearchKigenKanriList(name)
	c.String(http.StatusOK, message)
}

// １件登録
func AddData(c *gin.Context) {
	name := c.PostForm("name")
	message := "登録予定: " + name
	c.String(http.StatusOK, message)
}

// １件更新
func UpdateData(c *gin.Context) {
	c.String(http.StatusOK, "更新予定")
}

// １件削除
func DeleteData(c *gin.Context) {
	c.String(http.StatusOK, "削除予定")
}
