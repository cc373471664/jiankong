/*
@Time : 2020/1/13 2:08 下午
@Author : cc 373471664@qq.com
@Software: GoLand
@func:
*/
package controller

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"redisdui/app/module"
	"redisdui/db"
	"redisdui/util"
	"strings"
)

type QiwechatStr struct {
}
func (this *QiwechatStr) Tuisong() {
	conn := db.Redgo.Get()
	/** 查询所有需要发送的 *******/
	t, err := redis.Strings(conn.Do("keys", module.BaoJingTable()+"*"))
	if err != nil {
		fmt.Println(err.Error())
	}
	if len(t) > 0 {
		for _, v := range t {
			p, _ := redis.Values(conn.Do("hgetall", v))
			var baojin module.BaoJingDb
			redis.ScanStruct(p, &baojin)
			err := this.FaRen(baojin)
			if err == nil {
				conn.Do("DEL", v)
			} else {
				fmt.Println("redis处理发送的时候错误:", err.Error())
			}
		}
	}
}
func (this *QiwechatStr) FaRen(jingDb module.BaoJingDb) (err error) {
	/** 读取根目录ren.txt文件读取人 *******/
	t, err := db.LoadConfig("")
	if err != nil {
		fmt.Println("配置文件打开错误:", err.Error())
	}
	s := strings.Split(t.Qiwechat.Ren, ",")
	for _, v := range s {
		float:=util.SendMsg(t.Qiwechat.Wecahtid, t.Qiwechat.Secret, t.Qiwechat.AgentId, v, "", "",
			util.Textcard{
				Url: jingDb.Url,
				Description:jingDb.Name,
				Title:"服务掉线"+jingDb.AddTime,
				Btntxt:"查看",
			})
		if float==false{
			return errors.New("发送企业微信错误")
		}
	}
	return
}
