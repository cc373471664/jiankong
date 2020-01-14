/*
@Time : 2020/1/4 3:27 下午
@Author : cc 373471664@qq.com
@Software: GoLand
@func:
*/
package app

import (
	"github.com/gin-gonic/gin"
	"redisdui/app/controller"
)

func DefaultRoute(router *gin.Engine) {
	router.GET("/ceshi", func(context *gin.Context) {
		context.JSON(200, gin.H{"status": "OK"})
	})
	/** 列表的增删改查 *******/
	router.POST("/urllist",controller.ListINT(&controller.ListSTR{}).List)
	router.POST("/urllistlog",controller.ListINT(&controller.ListSTR{}).Listlog)
	router.POST("/addlist",controller.ListINT(&controller.ListSTR{}).AddList)
	router.POST("/dellist",controller.ListINT(&controller.ListSTR{}).DelList)
	router.POST("/editlist",controller.ListINT(&controller.ListSTR{}).EditList)
	/** 测试发送请求 *******/
	router.POST("/sendfa",controller.ListINT(&controller.ListSTR{}).SendFa)
	/** 点击启动或关闭 *******/
	router.POST("/send_pao",controller.ListINT(&controller.ListSTR{}).SendPao)
	/** 测试失败企业微信推送 *******/
	router.POST("/errqi",controller.ListINT(&controller.ListSTR{}).SendWechat)
	/** websocket *******/
	router.GET("/socketlist",controller.ListINT(&controller.ListSTR{}).SendSocketList)
}
