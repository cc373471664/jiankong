/*
@Time : 2020/1/9 9:01 上午
@Author : cc 373471664@qq.com
@Software: GoLand
@func:
*/
package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

var HbMq *gorm.DB
var err error
func InitDB()  {
	HbMq,err= gorm.Open("mysql", "root:root@/gigi?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		//fmt.Println(err)
		fmt.Println("DB初始化错误:",err)
		return
	}
	HbMq.DB().SetMaxIdleConns(10)
	HbMq.DB().SetMaxOpenConns(100)
	HbMq.DB().SetConnMaxLifetime(time.Second)
	HbMq.LogMode(true)
	HbMq.SingularTable(true)
	fmt.Println("hbdb数据库初始化成功")

}
