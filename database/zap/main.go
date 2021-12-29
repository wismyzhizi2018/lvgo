package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	zap.S().Infof("this is %s", "我和我的祖国")
	header := make(map[string]string)
	header["Authorization"] = "ac1bdd4ce3d8d110af3a7b8ef32bdfa5d76afa4b05d7206c71a99ecfeb97ce74"
	url := "http://order.api.nantang-tech.com/order/base/get_user_info"
	req := newget(url, header)
	fmt.Println(req)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json;charset=utf-8"

	res := newget(url, headers)
	fmt.Println(res)
}

func newget(url string, headers map[string]string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
	}
	// add params
	// q := req.URL.Query()
	// if params != nil {
	//  for key, val := range params {
	//      q.Add(key, val)
	//  }
	//  req.URL.RawQuery = q.Encode()
	// }
	// add headers
	if headers != nil {
		for key, val := range headers {
			fmt.Println(key)
			fmt.Println(val)
			req.Header.Add(key, val)
		}
	}
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		zap.S().Infof("this is %s", err)
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		zap.S().Infof("this is %s", err)
	}
	return string(resBody)
}
