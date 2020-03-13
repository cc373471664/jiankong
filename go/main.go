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
	"github.com/robfig/cron"
	"redisdui/app"
	"redisdui/app/controller"
	"redisdui/app/module"
	"redisdui/config"
	"redisdui/db"
	"strconv"
)

func main()  {
	//配置文档
	router := gin.Default()
	//全局跨域1
	router.Use(config.Logger())
	app.DefaultRoute(router)
	//	定时任务
	c := cron.New()
	c.AddFunc("*/5 * * * *", func() {
		t:=controller.QiwechatStr{}
		t.Tuisong()
	})
	c.AddFunc("* * * * *", func() {
		t:=module.ListModel{}
		t.Run()
	})
	c.Start()
	/** 读取根目录ren.txt文件读取人 *******/
	t, err := db.LoadConfig("")
	if err!=nil {
		panic("配置文件打开错误:"+err.Error())
	}
	router.Run(":"+strconv.Itoa(t.Main.Port))
}
