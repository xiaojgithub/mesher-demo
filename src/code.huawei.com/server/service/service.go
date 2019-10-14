package service

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/http"
)

//mesher listen 127.0.0.1:30101
var proxy = gorequest.New().Proxy("http://127.0.0.1:30101")

//use proxy way to connect provider demo-mesher-server's api /demo/hello
func Greeting() ([]byte, error) {
	resp, body, errs := proxy.Get("http://demo-mesher-server:8090/demo/hello").EndBytes()
	if errs != nil {
		return nil, fmt.Errorf(fmt.Sprintf("do request catch a err:%#v", errs))
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("request status not ok, %d", resp.StatusCode)
	}
	return body, nil
}
