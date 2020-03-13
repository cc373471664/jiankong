/*
@Time : 2020/1/14 10:28 上午
@Author : cc 373471664@qq.com
@Software: GoLand
@func:
*/
package db

import "github.com/spf13/viper"

// 加载配置
func LoadConfig(fpath string) (c *Config, err error) {
	if fpath == "" {
		fpath = "./config.yaml"
	}
	v := viper.New()
	v.SetConfigFile(fpath)
	v.SetConfigType("yaml")
	if err1 := v.ReadInConfig(); err1 != nil {
		err = err1
		return
	}
	c = &Config{}
	c.Qiwechat.Wecahtid = v.GetString("qiwechat.wecahtid")
	c.Qiwechat.AgentId = v.GetInt("qiwechat.agent_id")
	c.Qiwechat.Secret = v.GetString("qiwechat.secret")
	c.Qiwechat.Ren = v.GetString("qiwechat.ren")
	c.Redis.Path=v.GetString("redis.path")
	c.Redis.Host=v.GetString("redis.host")
	c.Main.Port=v.GetInt("main.port")
	return
}
// Config 配置参数
type Config struct {
	Qiwechat ConfigPei
	Redis RedisPei
	Main Main
}

type ConfigPei struct {
	Wecahtid string `json:"wecahtid"`
	AgentId  int `json:"agent_id"`
	Secret   string `json:"secret"`
	Ren      string `json:"ren"`
}

type RedisPei struct {
	Path string `json:"path"`
	Host string `json:"host"`
}
//主配置
type Main struct {
	Port int `json:"port"`
}
