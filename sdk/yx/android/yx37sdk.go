package android

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	. "sdk/util"

	. "sdk/common"

	"github.com/gin-gonic/gin"
	// "net/http"
	// "strconv"
	// "time"
)

type VerifyTokenReq struct {
	Pid   int    `json:"pid"`   //联运商ID
	Gid   int    `json:"gid"`   //游戏ID
	Time  int    `json:"time"`  //Unix时间戳 注：以服务端时间为准
	Sign  string `json:"sign"`  //校验令牌 strtolower( md5( "{$gid}{$time}{$key}" ) )
	Token string `json:"token"` //37手游生成的加密校验令牌
}

type VerifyTokenResp struct {
	State int `json:"state"` //1:成功 0:失败
	/*
		uid	Int	用户UID
		disname	String	显示名称
		is_adult	Int	0 未实名，1已实名； 注：pid为1或46才有效
		is_adult_age	Int	0 未认证，1认证未成年，2认证已成年； 注：pid为1或46才有效
		is_bind_phone	Int	0 未绑定，1已绑定
	*/
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

func VerifyToken37(c *gin.Context) {
	url := "http://vt-api.37.com.cn/verify/token/"
	// err := req.ParseForm()
	// if err != nil {
	// 	return
	// }
	// token := GetRequestString("user", req)
	token := c.PostForm("token")
	gid := strconv.Itoa(GID)
	nowtime := int(time.Now().Unix())
	nowstr := strconv.Itoa(nowtime)
	//校验令牌 strtolower( md5( "{$gid}{$time}{$key}" ) )
	//signstr := gid + nowstr + COMMONKEY
	//s := fmt.Sprintf("%x", md5.Sum([]byte(signstr)))
	sign := GetSign(gid + nowstr + COMMONKEY)
	fmt.Println("time:", nowtime, "sign:", sign)
	// reqs := make(map[string]string)
	// reqs["pid"] = strconv.Itoa(PID)
	// reqs["gid"] = strconv.Itoa(GID)
	// reqs["time"] = nowstr
	// reqs["sign"] = sign
	// reqs["token"] = token

	vrifydata := &VerifyTokenReq{
		Pid:   PID,
		Gid:   GID,
		Time:  nowtime,
		Sign:  sign,
		Token: token,
	}
	result := PostByJson(url, vrifydata, SDK_CONTENT_TYPE)
	//result := Post(url, reqs, "application/json")
	//resp.Write(result)
	verifyresp := &VerifyTokenResp{}
	json.Unmarshal(result, verifyresp)
	if verifyresp.State == 0 {
		fmt.Println(string(result))
	}
	gamereq := make(map[string]string)
	gamereq["platform"] = c.Param("platform")
	gamereq["User"] = c.PostForm("user")
	Post("http://127.0.0.1:11400/ping", gamereq, GAME_CONTENT_TYPE)
	fmt.Println("result:", gamereq)
	// gamereq := make(map[string]string)
	// gamereq["Uid"] = "123"
	// gamereq["Platform"] = PLATFORM
	// //向游戏服发送请求
	// url = FRONTURL + "/GetToken?"
	// param := TransMapToUrlParam(gamereq)
	// result = Post(url+param, nil, GAME_CONTENT_TYPE)
	//result = Post(url, vrifydata)
	c.JSON(200, gin.H{
		"user":     gamereq["User"],
		"platform": gamereq["platform"],
	})
}
