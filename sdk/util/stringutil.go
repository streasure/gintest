package util

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"
)

func TransMapToUrlParam(params map[string]string) string {
	//paramsArr := make([]string, 0, len(params))
	//for k, v := range params {
	// paramsArr = append(paramsArr, fmt.Sprintf("%s=%s", k, v))
	//}

	//paramStr = strings.Join(paramsArr, "&")

	var uri url.URL
	q := uri.Query()
	for k, v := range params {
		q.Add(k, v)
	}

	return q.Encode()
}

func GetSign(signstr string) string {
	s := fmt.Sprintf("%x", md5.Sum([]byte(signstr)))
	sign := strings.ToLower(s)
	fmt.Println("GetSign sign----", sign)
	return sign
}
