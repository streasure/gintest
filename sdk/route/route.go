package route

import (
	// . "sdk/util"

	. "sdk/yx/android"

	"github.com/gin-gonic/gin"
)

type Method func(c *gin.Context)

var RoutesLogin map[string]Method

var RoutesPay map[string]Method

func SetupRouter() *gin.Engine {
	RegistHandle()
	// 初始化 Gin 框架默认实例，该实例包含了路由、中间件以及配置信息
	r := gin.Default()
	r.POST("/Login/:platform", func(c *gin.Context) {
		platform := c.Param("platform")
		RoutesLogin[platform](c)
		// gamereq := make(map[string]string)
		// gamereq["platform"] = c.Param("platform")
		// gamereq["User"] = c.PostForm("user")
		// Post("http://127.0.0.1:11400/ping", gamereq, "application/x-www-form-urlencoded")
		// fmt.Println("result:", gamereq)
		// c.JSON(200, gin.H{
		// 	"user":     gamereq["user"],
		// 	"platform": gamereq["platform"],
		// })
	})

	return r
}

func RegistHandle() {
	RoutesLogin = map[string]Method{
		"yx37login": VerifyToken37,
	}
}
