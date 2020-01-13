/*
@Time : 2020/1/9 4:57 下午
@Author : cc 373471664@qq.com
@Software: GoLand
@func:
*/
package controller

import (
	"github.com/gin-gonic/gin"
	"redisdui/app/module"
	"redisdui/util"
)

type ListINT interface {
	//列表
	List(c *gin.Context)
	//日志
	Listlog(c *gin.Context)
	//添加
	AddList(c *gin.Context)
	//删除
	DelList(c *gin.Context)
	//修改
	EditList(c *gin.Context)
//	发送请求
	SendFa(c *gin.Context)
//	跑
	Run(c *gin.Context)
//	是否启用
	SendPao(c *gin.Context)
}

type ListSTR struct {
	ListModel module.ListModel
}
func (this *ListSTR)List(c *gin.Context)  {
	data,err:=this.ListModel.List()
	if err!=nil {
		util.Echo(false,err.Error(),c)
	}else {
		util.Echo(true,data,c)
	}
}
func (this *ListSTR)AddList(c *gin.Context)  {
	/** 接收 *******/
	var urllist module.UrlListDB
	c.BindJSON(&urllist)
	/** 验证 *******/
	if urllist.Name==""||urllist.Url=="" || urllist.Jiange==0 {
		util.Echo(false,"请填写完整",c)
		return
	}
	is_you:=this.ListModel.Is_You(urllist.Name)
	if is_you {
		util.Echo(false,"已经存在了",c)
		return
	}
	/** 插入 *******/
	err:=this.ListModel.ListSave(urllist)
	if err!=nil {
		util.Echo(false,err.Error(),c)
	}else{
		util.Echo(true,"添加成功",c)
	}
}
func (this *ListSTR)DelList(c *gin.Context)  {
	/** 接收 *******/
	var urllist module.UrlListDB
	c.BindJSON(&urllist)
	if urllist.Name=="" {
		util.Echo(false,"请填写参数",c)
		return
	}
	err:=this.ListModel.Del(urllist.Name)
	if err!=nil {
		util.Echo(false,err.Error(),c)
	}
	util.Echo(true,"删除成功",c)
}
func (this *ListSTR)EditList(c *gin.Context)  {
	/** 接收 *******/
	var urllist module.UrlListDB
	c.BindJSON(&urllist)
	is_you:=this.ListModel.Is_You(urllist.Name)
	if is_you {
		err:=this.ListModel.ListSave(urllist)
		if err!=nil {
			util.Echo(false,err.Error(),c)
		}else{
			util.Echo(true,"修改成功",c)
		}
	}else {
		util.Echo(false,"不存在",c)
	}
}
func (this *ListSTR)SendFa(c *gin.Context)  {
	/** 接收 *******/
	var urllist module.UrlListDB
	c.BindJSON(&urllist)
	is_you:=this.ListModel.Is_You(urllist.Name)
	if is_you {
		err:=this.ListModel.SendUrl(urllist.Name,true)
		if err==false {
			util.Echo(false,"访问失败",c)
		}else{
			util.Echo(true,"请求成功",c)
		}
	}else{
		util.Echo(false,"不存在",c)
	}
}
func (this *ListSTR)Run(c *gin.Context)  {
	this.ListModel.Run()
}

func (this *ListSTR)Listlog(c *gin.Context)  {
	/** 接收 *******/
	var input module.ListInput
	c.BindJSON(&input)
	data,count,err:=this.ListModel.Listlog(input)
	if err!=nil {
		util.Echo(false,err.Error(),c)
	}else {
		data:=util.EchoListSTR{
			Data:  data,
			Count: count,
		}
		util.Echo(true,data,c)
	}
}
func (this *ListSTR)SendPao(c *gin.Context)  {
	/** 接收 *******/
	var urllist module.UrlListDB
	c.BindJSON(&urllist)
	err:=this.ListModel.GaiYong(urllist.Name,urllist.Yong)
	if err!=nil {
		util.Echo(true,err.Error(),c)
	}else{
		util.Echo(false,"更改运行状态成功",c)
	}
}