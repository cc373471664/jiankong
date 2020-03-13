package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)
var (
	CorpId         = ""   //企业微信 corpid
	TokenCache *cache.Cache
)
//actoken的结构体
type AccessToken  struct {
	Errcode string `json:"errcode"`
	Errmsg string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn string `json:"expires_in"`
}
//接收回折消息的结构体
type SendMsgHui struct {
	Errcode string `json:"errcode"`
	Errmsg string `json:"errmsg"`
}
//初始化缓存
func init()  {
	//创建一个60分钟过期一次,80分钟清理一次的cache
	TokenCache = cache.New(60*time.Minute, 80*time.Minute)

}
//查询cache是否有值
func QuAcToken(Corpid string,EncodingAESKey string) string {
	//先去缓存取有没有token
	foot,bool:=TokenCache.Get(EncodingAESKey)
	if bool==false  {
		gettoken_url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + Corpid + "&corpsecret=" + EncodingAESKey
		client := &http.Client{}
		req, _ := client.Get(gettoken_url)
		defer req.Body.Close()
		body, _ := ioutil.ReadAll(req.Body)
		var json_str AccessToken
		json.Unmarshal([]byte(body), &json_str)
		if(json_str.Errmsg=="ok"){
			TokenCache.Set(EncodingAESKey,json_str.AccessToken,cache.DefaultExpiration)
			return json_str.AccessToken
		}else{
		}
	}
	return foot.(string)
}

/**
发送应用消息
*/
//文本卡片消息结构体
type MESSAGES struct {
	Touser string `json:"touser"`
	Toparty string `json:"toparty"`
	Totag string `json:"totag"`
	Msgtype string `json:"msgtype"`
	Agentid int `json:"agentid"`
	Textcard Textcard `json:"textcard"`
	Enable_id_trans int `json:"enable_id_trans"`
	Enable_duplicate_check int `json:"enable_duplicate_check"`
}
type Textcard struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Url string `json:"url"`
	Btntxt string `json:"btntxt"`
}
//接收数据
type Messagesshou struct {
	ErrorMsg string `json:"errmsg"`
	ErrorCode int `json:"errcode"`
}
//拼凑卡片数据结构体
func messages(touser string,toparty string,agentid int,totag string,Textcard Textcard) string{
	msg := MESSAGES{
		Touser: touser,
		Toparty: toparty,
		Totag:totag,
		Msgtype: "textcard",
		Agentid: agentid,
		Textcard:Textcard,
		Enable_id_trans:0,
		Enable_duplicate_check:0,
	}
	sed_msg, _ := json.Marshal(msg)
	return string(sed_msg)
}

//发送
func SendMsg(Corpid string,EncodingAESKey string,agentid int,touser string,toparty string,totag string,Textcard Textcard) bool  {
	//取token
	AppActoken:=QuAcToken(Corpid,EncodingAESKey)
	send_url  := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token="+AppActoken+""
	client := &http.Client{}
	msg := strings.Replace(messages(touser,toparty,agentid,totag,Textcard),"\\\\","\\",-1)
	fmt.Println(msg,"ss")
	req, _ := http.NewRequest("POST", send_url, bytes.NewBuffer([]byte(msg)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset","UTF-8")
	resp, err := client.Do(req)
	if err!=nil {
		fmt.Println("企业微信推送接口调用错误"+err.Error())
		return false
	}else{
		body, _ := ioutil.ReadAll(resp.Body)
		var message Messagesshou
		json.Unmarshal(body,&message)
		if message.ErrorCode!=0 {
			fmt.Println("企业微信推送接口调用错误"+strconv.Itoa(message.ErrorCode)+":"+message.ErrorMsg)
			return false
		}else{
			return true
		}
	}
	defer resp.Body.Close()
	return false
}

//json序列化(禁止 html 符号转义)
func encodeJson(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
//string 类型转 int
func StringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("agent 类型转换失败, 请检查配置文件中 agentid 配置是否为纯数字(%v)", err)
		return 0
	}
	return n
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
//AES解密
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("aes解密失败: %v", err)
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

