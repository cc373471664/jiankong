/*
@Time : 2020/1/4 3:42 下午
@Author : cc 373471664@qq.com
@Software: GoLand
@func:
*/
package db

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var Redgo *redis.Pool

func init()  {
	Redgo=newPool()
}
func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func () (redis.Conn, error) {
			c,err:=redis.Dial("tcp", "localhost:6379")
			if err!=nil {
				log.Fatal("redis连接失败",err.Error())
			}
			return c, nil
		},
	}
}