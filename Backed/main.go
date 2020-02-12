package main

import "ITBat/router"

func main() {
	// 禁用控制台颜色, 将日志写入文件时不需要控制台颜色
	r := router.InitRoute()
	r.Run(":8080")
}
