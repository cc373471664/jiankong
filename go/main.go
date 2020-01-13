/*
@Time : 2020/1/4 3:08 下午
@Author : cc 373471664@qq.com
@Software: GoLand
@func:
*/
package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"redisdui/app"
	"redisdui/app/module"
	"redisdui/config"
	"time"
)

func main()  {
	//配置文档
	router := gin.Default()
	//全局跨域1
	router.Use(config.Logger())
	app.DefaultRoute(router)
	go func() {
		for{
			t:=module.ListModel{}
			t.Run()
			time.Sleep(1*time.Second)
		}
	}()
	router.Run()
}
