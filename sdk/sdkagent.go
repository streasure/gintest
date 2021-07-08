package main

import (
	. "sdk/route"
)

func main() {
	// 设置路由信息
	r := SetupRouter()
	// 启动服务器并监听 8080 端口
	r.Run(":18000")
}
