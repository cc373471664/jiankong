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
/**
  @author 373471664@qq.com cc 2020-01-13 14:27:35
  @action 删除指定开头的key @param @return
*/
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
/**
  @author 373471664@qq.com cc 2020-01-13 14:27:51
  @action 分页 @param @return
*/
func RedisPage(current,pagesize int) (int,int)  {
	if current==1 {
		return 0,pagesize
	}else{
		return (current-1)*pagesize,current*pagesize-1
	}
}
