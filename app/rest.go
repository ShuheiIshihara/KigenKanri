package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func GetList(c *gin.Context) {
	name := c.Query("name")
	fmt.Printf(name)
	message := "一覧取得予定: " + name
	c.String(http.StatusOK, message)
}

func AddData(c *gin.Context) {
	name := c.PostForm("name")
	message := "登録予定: " + name
	c.String(http.StatusOK, message)
}

func UpdateData(c *gin.Context) {
	c.String(http.StatusOK, "更新予定")
}

func DeleteData(c *gin.Context) {
	c.String(http.StatusOK, "削除予定")
}
