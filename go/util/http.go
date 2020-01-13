/*
@Time : 2020/1/10 1:49 下午
@Author : cc 373471664@qq.com
@Software: GoLand
@func:
*/
package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func SendHttpGet(url string) (body []byte,code int, err error){
	resp, err :=   http.Get(url)
	if err != nil {
		return nil,0,err
	}
	defer resp.Body.Close()
	code=resp.StatusCode
	body, err = ioutil.ReadAll(resp.Body)
	return
}
//post json
func PostJson(fa interface{},url string)  (bool,[]byte) {
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(fa)
	req, err := http.NewRequest("POST", url, requestBody)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false,nil
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return true,body
}