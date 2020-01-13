/*
@Time : 2020/1/9 5:50 下午
@Author : cc 373471664@qq.com
@Software: GoLand
@func:
*/
package util

import "github.com/gin-gonic/gin"

type EchoListSTR struct {
	Data interface{} `json:"data"`
	Count int `json:"count"`
}
func Echo(sta bool,data interface{},c *gin.Context)  {
	if sta {
		c.JSON(200,gin.H{"data":data,"sta":1})
	}else{
		c.JSON(200,gin.H{"data":data,"sta":0})
	}
}
func Echolist(sta bool,list EchoListSTR,c *gin.Context)  {
	if sta {
		c.JSON(200,gin.H{"data":list,"sta":1})
	}else{
		c.JSON(200,gin.H{"data":list,"sta":0})
	}
}