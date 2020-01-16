/*
@Time : 2020/1/9 5:09 下午
@Author : cc 373471664@qq.com
@Software: GoLand
@func:
*/
package module

import (
	"github.com/gomodule/redigo/redis"
	"redisdui/db"
	"redisdui/util"
	"strconv"
	"time"
)

/** list主列表 *******/
type UrlListDB struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	Name    string `json:"name"`
	Jiange  int    `json:"jiange"`
	AddTime string `json:"add_time"`
	XiaTime int64  `json:"xia_time"`
	Sta     int    `json:"sta"`
	Yong    int    `json:"yong"`
	ZhiTime string `json:"zhi_time"`
}

/** 日志 *******/
type UrlListLogDB struct {
	Sta      int    `json:"sta"`
	Url      string `json:"url"`
	Name     string `json:"name"`
	Errmsg   string `json:"msg"`
	AddTime  string `json:"add_time"`
	UnixTime string `json:"unix_time"`
	Isce     bool   `json:"isce"`
}

/** 报价 *******/
type BaoJingDb struct {
	Name    string `json:"name"`
	AddTime string `json:"add_time"`
	IsBao   bool   `json:"is_bao"`
	Url     string `json:"url"`
}

func ListTable() string {
	return "list_"
}
func ListLogTable() string {
	return "listlog_"
}
func ListTableOrder() string {
	return "order_list"
}
func ListLogTableOrder() string {
	return "order_listlog_"
}
func BaoJingTable() string {
	return "listcurlbaojing_"
}

type ListModel struct {
}
type ListInput struct {
	Name     string `json:"name"`
	Current  int    `json:"current"`
	Pagesize int    `json:"pagesize"`
}

func (this *ListModel) Is_You(name string) bool {
	conn := db.Redgo.Get()
	//	判断是否存在
	is_key_exit, _ := redis.Bool(conn.Do("EXISTS", ListTable()+name))
	if is_key_exit {
		return true
	}
	return false
}
func (this *ListModel) ListSave(listDB UrlListDB) error {
	conn := db.Redgo.Get()
	listDB.Yong = 1
	listDB.AddTime = util.SendTime()
	_, err := conn.Do("lpush", ListTableOrder(), listDB.Name)
	_, err = conn.Do("hmset", redis.Args{}.Add(ListTable()+listDB.Name).AddFlat(&listDB)...)
	return err
}
func (this *ListModel) List() (data []UrlListDB, err error) {
	conn := db.Redgo.Get()
	list, err := redis.Strings(conn.Do("lrange", "order_list", 0, -1))
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		for _, v := range list {
			var listdb UrlListDB
			val, err := redis.Values(conn.Do("hgetall", ListTable()+v))
			if err != nil {
				return nil, err
			}
			redis.ScanStruct(val, &listdb)
			data = append(data, listdb)
		}
	}
	return
}
func (this *ListModel) Del(name string) error {
	//list里删除
	_, err := db.Redgo.Get().Do("lrem", ListTableOrder(), 0, name)
	//hash删除
	_, err = db.Redgo.Get().Do("DEL", ListTable()+name)
	//删除log
	util.DelKeys(ListLogTable() + name + "_*")
	//删除log的list
	util.DelKeys(ListLogTableOrder() + name)
	//删除提醒日志
	util.DelKeys(BaoJingTable()+name+"_*")
	return err
}
func (this *ListModel) Dan(name string) (err error, urlist UrlListDB) {
	var listdb UrlListDB
	val, err := redis.Values(db.Redgo.Get().Do("hgetall", ListTable()+name))
	if err != nil {
		return err, listdb
	}
	redis.ScanStruct(val, &listdb)
	return nil, listdb
}

//发送请求,isce标识是否是测试,不是测试的不进去推送
func (this *ListModel) SendUrl(name string, isce bool) (sta bool) {
	errmsg := ""
	conn := db.Redgo.Get()
	err, urllist := this.Dan(name)
	_, code, err := util.SendHttpGet(urllist.Url)
	if err != nil {
		errmsg = err.Error()
	}
	s := strconv.FormatInt(time.Now().Unix(), 10)
	urllistlogdb := UrlListLogDB{
		Url:      urllist.Url,
		Name:     urllist.Name,
		Sta:      code,
		Errmsg:   errmsg,
		AddTime:  util.SendTime(),
		Isce:     isce,
		UnixTime: s,
	}
	/** 插入日志 *******/
	conn.Do("hmset", redis.Args{}.Add(ListLogTable()+name+"_"+s).AddFlat(&urllistlogdb)...)
	conn.Do("lpush", ListLogTableOrder()+urllistlogdb.Name, urllistlogdb.UnixTime)
	//加入过期
	conn.Do("expire",ListLogTable()+name+"_"+s,60*60*24)
	conn.Do("expire",ListLogTableOrder()+urllistlogdb.Name,60*60*24)

	//更新下次时间戳
	urllist.XiaTime = time.Now().Unix() + int64(urllist.Jiange)
	urllist.ZhiTime = util.SendTime()
	urllist.Sta = code

	/** 进行是否3次失败判断,超过3次停止发,并且记录报警日志,不记录测试 *******/
	if isce == false {
		if urllistlogdb.Sta != 200 {
			zui, _ := redis.Strings(conn.Do("lrange", ListLogTableOrder()+name, 0, 2))
			if len(zui) == 3 {
				biaozhi := 0
				var listdb UrlListDB
				for _, v := range zui {
					val, _ := redis.Values(conn.Do("hgetall", ListLogTable()+name+"_"+v))
					redis.ScanStruct(val, &listdb)
					if listdb.Sta != 200 {
						biaozhi++
					}
				}

				if biaozhi == 3 {
					/** 如果队列里还有没发同个提醒就不发 *******/
					biaozhi_ce := 0
					t, _ := redis.Strings(conn.Do("keys", BaoJingTable()+name+"_*"))
					if len(t) > 0 {
						for _, v := range t {
							val, _ := redis.Values(conn.Do("hgetall", v))
							var baojing BaoJingDb
							redis.ScanStruct(val, &baojing)
							if baojing.IsBao == false {
								biaozhi_ce = 1
							}
						}
					}
					if biaozhi_ce == 0 {
						/** 插入报警日志 *******/
						conn.Do("hmset", redis.Args{}.Add(BaoJingTable()+name+"_"+s).AddFlat(&BaoJingDb{Name: name, AddTime: urllist.ZhiTime, IsBao: false, Url: urllist.Url})...)

					}
					urllist.Yong = 0
				}
			}
		}
	}
	conn.Do("hmset", redis.Args{}.Add(ListTable()+name).AddFlat(&urllist)...)
	if code == 200 {
		return true
	} else {
		return false
	}
}

func (this *ListModel) Run() {
	/** 查询所有任务 *******/
	data, _ := this.List()
	for _, v := range data {
		if v.XiaTime <= time.Now().Unix() && v.Yong == 1 {
			go this.SendUrl(v.Name, false)
		}
	}
}
func (this *ListModel) Listlog(input ListInput) (data []UrlListLogDB, count int, err error) {
	conn := db.Redgo.Get()
	curr, page := util.RedisPage(input.Current, input.Pagesize)
	list, err := redis.Strings(conn.Do("lrange", ListLogTableOrder()+input.Name, curr, page))
	count, err = redis.Int(conn.Do("llen", ListLogTableOrder()+input.Name))
	if err != nil {
		return nil, 0, err
	}
	if len(list) > 0 {
		for _, v := range list {
			var listlogdb UrlListLogDB
			val, err := redis.Values(conn.Do("hgetall", ListLogTable()+input.Name+"_"+v))
			if err != nil {
				return nil, 0, err
			}
			redis.ScanStruct(val, &listlogdb)
			data = append(data, listlogdb)
		}
	}
	return
}
func (this *ListModel) GaiYong(name string, yong int) error {
	_, err := db.Redgo.Get().Do("hset", ListTable()+name, "Yong", yong)
	this.SendUrl(name, true)
	return err
}
