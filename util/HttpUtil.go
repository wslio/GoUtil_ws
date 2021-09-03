package util

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

func HttpGet(uu string, params url.Values) string {
	reqUrl, _ := url.ParseRequestURI(uu)
	reqUrl.RawQuery = params.Encode()
	resp, err := http.Get(reqUrl.String())
	HandleError(err, "get fail ", false)
	tmp := HandleResp(resp)
	return tmp
}

func HttpPost(uu string, data string) string {
	c := http.Client{}
	c.Timeout = 15 * time.Second
	resp, err := c.Post(uu, "application/x-www-form-urlencoded", strings.NewReader(url.Values{"jsonrequest": {data}}.Encode()))
	HandleError(err, uu + " POST请求失败", false)
	tmp := HandleResp(resp)
	return tmp
}