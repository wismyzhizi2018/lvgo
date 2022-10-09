package main

import (
	"fmt"
	"github.com/namsral/flag"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

var list = flag.String("list", "sh000001,sz000423,sh601857,sh600884,sh601162,sz002409,sz000671,sz002547,sh600030", "please input search number")



func main() {
    flag.Parse()
    for {
        logger, _ := zap.NewDevelopment()
        zap.ReplaceGlobals(logger)
        zap.S().Infof("listen is %s", *list)
        header := make(map[string]string)
        header["Referer"] = "https://finance.sina.com.cn"
        url := "http://hq.sinajs.cn/list=" + *list
        req := newget(url, header)
        fmt.Println(req)
        time.Sleep(30 * time.Second)
    }
}

func newget(url string, headers map[string]string) string {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        zap.S().Infof("this is %s", err)
    }
    if headers != nil {
        for key, val := range headers {
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