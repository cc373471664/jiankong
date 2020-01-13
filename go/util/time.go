/*
@Time : 2020/1/10 10:58 上午
@Author : cc 373471664@qq.com
@Software: GoLand
@func:
*/
package util

import "time"

func SendTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
