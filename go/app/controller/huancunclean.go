/*
@Time : 2020/1/14 10:26 上午
@Author : cc 373471664@qq.com
@Software: GoLand
@func:
*/
package controller

import (
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"redisdui/app/module"
	"redisdui/db"
	"redisdui/util"
)

type HuancuncleanINT interface {
	Clean(c *gin.Context)
//	读取缓存大小
	DuHuan(c *gin.Context)
}
type HuancuncleanSTR struct {

}

func (this *HuancuncleanSTR)Clean(c *gin.Context)  {
	err:=util.DelKeys(module.ListLogTableOrder())
	err=util.DelKeys(module.ListLogTable())
	_,err=db.Redgo.Get().Do("BGREWRITEAOF")
	if err!=nil {
		c.JSON(200,gin.H{"err":"清理失败:"+err.Error()})
	}else{
		c.JSON(200,gin.H{"err":""})
	}
}

func (this *HuancuncleanSTR)DuHuan(c *gin.Context)  {
	t, err := db.LoadConfig("")
	if err != nil {
		c.JSON(200,gin.H{"size":0,"err":"文件打开错误:"+err.Error()})
		return
	}
	c.JSON(200,gin.H{"size":getFileSize(t.Redis.Path)/1024/1024})
}
func getFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}
