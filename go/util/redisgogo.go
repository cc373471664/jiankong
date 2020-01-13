/*
@Time : 2020/1/13 9:23 上午
@Author : cc 373471664@qq.com
@Software: GoLand
@func:
*/
package util

import (
	"github.com/gomodule/redigo/redis"
	"redisdui/db"
)

func DelKeys(name string) (err error)  {
	conn:=db.Redgo.Get()
	list,err:=redis.Strings(conn.Do("keys",name+"*"))
	if err!=nil {
		return
	}
	if len(list)>0 {
		for _,v:=range list{
			_,err=db.Redgo.Get().Do("DEL",v)
		}
	}
	return
}
func RedisPage(current,pagesize int) (int,int)  {
	if current==1 {
		return 0,pagesize
	}else{
		return (current*pagesize)-pagesize+1,current*pagesize
	}
}