package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//2. 绑定路由规则,api
	r.POST("/ping", func(c *gin.Context) {
		// c.Header("Access-Control-Allow-Origin", "*")
		////第二个参数是默认数值
		uid := c.PostForm("User")
		platform := c.PostForm("platform")
		fmt.Println("uid", uid, "platform", platform)
		c.String(http.StatusOK, fmt.Sprintf("hello: %s", uid))
	})
	r.Run(":11400")
}
