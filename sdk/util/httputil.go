package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var client = http.Client{
	Timeout: 10 * time.Second,
}

func GetRequestString(name string, req *http.Request) string {
	v, _ := req.Form[name]
	if v != nil {
		return v[0]
	}

	return ""
}

func GetRequestInt(name string, req *http.Request) int {
	v, _ := req.Form[name]
	if v != nil {
		number, _ := strconv.Atoi(v[0])
		return number
	}
	return -1
}

func HttpGetRequest(url string, result interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)

	return err
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data map[string]string, content string) []byte {
	// 超时时间：5秒
	// jsonStr, err := json.Marshal(data)
	// if nil != err {
	// 	fmt.Println("json.Marshal error:", err)
	// 	return []byte{}
	// }
	str := TransMapToUrlParam(map[string]string(data))
	fmt.Println("message", str)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Post(url, content, strings.NewReader(str))
	if err != nil {
		//panic(err)
		fmt.Println("client.Post error:", err)
		return []byte{}
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return result
}

func PostByJson(url string, data interface{}, content string) []byte {
	// 超时时间：5秒
	jsonStr, err := json.Marshal(data)
	if nil != err {
		fmt.Println("json.Marshal error:", err)
		return []byte{}
	}
	fmt.Println("message", string(jsonStr))

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Post(url, content, bytes.NewReader(jsonStr))
	if err != nil {
		//panic(err)
		fmt.Println("client.Post error:", err)
		return []byte{}
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return result
}

//Get get请求url参数
func HttpGet(req *http.Request) map[string]string {
	var result = make(map[string]string)
	keys := req.URL.Query()
	for k, v := range keys {
		result[k] = v[0]
	}

	return result
}

//PostForm 获取postform形式的参数
func PostForm(req *http.Request) map[string]string {
	//body, _ := ioutil.ReadAll(req.Body)
	var result = make(map[string]string)
	req.ParseForm()
	for k, v := range req.PostForm {
		if len(v) < 1 {
			continue
		}

		result[k] = v[0]
	}
	return result
}

//PostJson 获取post json参数
func PostJson(req *http.Request, obj interface{}) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, obj)
	if err != nil {
		return err
	}

	return nil
}
