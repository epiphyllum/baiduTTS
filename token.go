/*
访问如下网址， 可用浏览器测试

// appKey = Va5yQRHl********LT0vuXV4
// appSecret = 0rDSjzQ20XUj5i********PQSzr5pVw2

https://openapi.baidu.com/oauth/2.0/token?grant_type=client_credentials&client_id=Va5yQRHl********LT0vuXV4&client_secret=0rDSjzQ20XUj5i********PQSzr5pVw2

可以获取如下结果

{
    "access_token": "1.a6b7dbd428f731035f771b8d********.86400.1292922000-2346678-124328",
    "expires_in": 2592000,
    "refresh_token": "2.385d55f8615fdfd9edb7c4b********.604800.1293440400-2346678-124328",
    "scope": "public audio_tts_post ...",
    "session_key": "ANXxSNjwQDugf8615Onqeik********CdlLxn",
    "session_secret": "248APxvxjCZ0VEC********aK4oZExMB",
}

scope中含有audio_tts_post 表示有语音合成能力，没有该audio_tts_post 的token调用接口会返回502错误。*/

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type TTSApi struct {
	Token   string `json:"access_token"`
	Expires int64  `json:"expires_in"`
	Scope   string `json:"scope"`
}

func getToken(apiKey, secretKey string) (token string, expires int64, err error) {
	url1 := fmt.Sprintf("https://openapi.baidu.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s",
		apiKey, secretKey)
	r, err := http.Get(url1)
	if err != nil {
		return
	}
	defer r.Body.Close()
	buf := bytes.NewBufferString("")
	rd1 := bufio.NewReader(r.Body)
	rd1.WriteTo(buf)
	res := new(TTSApi)
	err = json.Unmarshal(buf.Bytes(), res)
	if err != nil {
		return
	}
	//log.Println(*res)
	if strings.Contains(res.Scope, "audio_tts_post") {
		token = res.Token
		expires = res.Expires
		err = nil
	} else {
		err = errors.New("no tts")
	}
	return
}

type AppIDs struct {
	AppID     string
	ApiKey    string
	SecretKey string
}
