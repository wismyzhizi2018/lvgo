package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/CodyGuo/dingtalk"
	"github.com/golang/glog"
	"github.com/gookit/color"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/namsral/flag"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"order/config"
	"os"
	"path"
	"strings"
	"time"
)

var httpClient http.Client
var retry int

func init() {
	httpClient = http.Client{
		Timeout: 10 * time.Second,
	}
	retry = 3
}

type YamlConfig struct {
	Version int      `yaml:"version"`
	Ding    DingApp  `yaml:"ding"`
	Config  []Config `yaml:"config"`
}
type DingApp struct {
	Webhook   string `yaml:"webhook"`
	Secretkey string `yaml:"secretkey"`
}
type Config struct {
	Name     string `yaml:"name"`
	Time     int64  `yaml:"time"`
	Path     string `yaml:"path"`
	Desc     string `yaml:"desc"`
	LastTime string `yaml:"last"`
	NowTime  string `yaml:"now"`
}

type DingConfig struct {
	Msgtype  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
}
type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
var endpoint = flag.String("endpoint", "<point>", "nacos endpoint")
var namespaceId = flag.String("namespace_id", "<namespace_id>", "nacos namespace Id")
var accessKey = flag.String("access_key", "nacos", "nacos access key")
var secretKey = flag.String("secret_key", "nacos", "nacos secret key")
var dataId = flag.String("data_id", "order-config.yaml", "nacos secret key")
var group = flag.String("group", "dev", "nacos secret key")
var port = flag.Uint64("port", 8848, "nacos port")
var c YamlConfig

func main() {
	flag.Parse()
	config.NewNacosConfig(*endpoint, *namespaceId, *accessKey, *secretKey, *dataId, *group, *port)
	if *namespaceId != "<namespace_id>" {
		InitNACOS()
	} else {
		fmt.Println(*namespaceId)
		v := viper.New()
		path, _ := os.Getwd()
		v.AddConfigPath(path)
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := v.Unmarshal(&c); err != nil {
			panic(err)
		}
	}

	cron2 := cron.New() //创建一个cron实例
	//执行定时任务（每5秒执行一次）
	err := cron2.AddFunc("*/5 * * * * *", func() {
		//
		var content = []Config{}
		for _, item := range c.Config {
			if ok, _ := Exists(item.Path); ok {
				nowTime := time.Now().Unix()
				lastTime, _ := FileLastUpdateTime(item.Path)
				diffTime := nowTime - lastTime
				if diffTime > item.Time {
					fmt.Printf("%s大于设置时间%v秒\n", item.Desc, item.Time)
					//需要获取推送的订单信息
					item.LastTime = timeToData(lastTime)
					item.NowTime = timeToData(nowTime)
					content = append(content, item)
				}
			} else {
				fmt.Printf("文件%s不存在\n", item.Path)
			}
		}
		if len(content) > 0 {
			header := make(map[string]string)
			header["Content-Type"] = "application/json;charset=utf-8"
			var bt bytes.Buffer
			for _, v := range content {
				txt := ""
				txt += "------------------------\n\n"
				txt += fmt.Sprintf(" **监控说明：** %s\n\n", v.Desc)
				txt += fmt.Sprintf(" **监控名称：** %s\n\n", v.Name)
				txt += fmt.Sprintf(" **监控时差：** %v秒\n\n", v.Time)
				txt += fmt.Sprintf(" **监控文件：** %s\n\n", v.Path)
				txt += fmt.Sprintf(" **文件时间：** %s\n\n", v.LastTime)
				txt += fmt.Sprintf(" **当前时间：** %s\n\n", v.NowTime)

				bt.WriteString(txt)
			}

			secret := c.Ding.Secretkey
			webHook := c.Ding.Webhook
			// markdown类型
			dt := dingtalk.New(webHook, dingtalk.WithSecret(secret))
			markdownTitle := "日志监控通知"
			markdownText := bt.String()
			if err := dt.RobotSendMarkdown(markdownTitle, markdownText); err != nil {
				glog.Fatal(err)
			}
			fmt.Println(dt)
		}
		fmt.Println(content)
	})
	if err != nil {
		fmt.Println(err)
	}
	//启动/关闭
	cron2.Start()
	defer cron2.Stop()
	select {
	//查询语句，保持程序运行，在这里等同于for{}
	}

}
func EncodeSHA256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	sum := h.Sum(nil)
	message1 := base64.StdEncoding.EncodeToString(sum)

	message2 := base64.URLEncoding.EncodeToString([]byte(message1))
	//uv := url.Values{}
	//uv.Add("0", message1)
	//message2 := uv.Encode()[2:]
	return message2
}
func hmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
func timeToData(timestamp int64) string {
	timeFormat := "2006-01-02 15:04:05"
	// 时间戳转日期
	t3 := time.Unix(timestamp, 0)
	return t3.Format(timeFormat)
}

func FileLastUpdateTime(path string) (int64, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return time.Now().Unix(), err
	}
	fi, err := os.Stat(path)
	if err != nil {
		return time.Now().Unix(), err
	}
	return fi.ModTime().Unix(), nil
}
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// InitNACOS
// @title  从配置中心读取配置文件
// @description  从配置中心读取配置文件
func InitNACOS() {
	mainDirectory, _ := os.Getwd()
	logFilePath := mainDirectory + "/tmp/nacos/log/"
	logFileName := "nacos.log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	//写入文件
	_, err := os.Stat(fileName)
	if !(err == nil || os.IsExist(err)) {
		var err error
		//目录不存在则创建
		if _, err = os.Stat(logFilePath); err != nil {
			if err = os.MkdirAll(logFilePath, 0777); err != nil { //这里如果是0711权限 可能会导致其它线程，读取文件夹内内容出错
				color.Danger.Println("Create log dir err :", err)
			}
		}
		//创建文件
		if _, err = os.Create(fileName); err != nil {
			color.Danger.Println("Create log file err :", err)
		}
	}
	nacosConf := config.GetNacosConfig()

	sc := []constant.ServerConfig{
		{
			IpAddr: nacosConf.Endpoint,
			Port:   nacosConf.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         nacosConf.NamespaceId, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5 * 1000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		AccessKey:           nacosConf.AccessKey,
		SecretKey:           nacosConf.SecretKey,
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
		ListenInterval:      30 * 1000,
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		color.Danger.Println("nacos read Error = ", err.Error(), "运行中断")
		fmt.Println(err.Error())
		os.Exit(200)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: nacosConf.DataId,
		Group:  nacosConf.Group,
	})

	if err != nil {
		color.Danger.Println("env read Error = ", err.Error(), "运行中断")
		fmt.Println(err.Error())
		os.Exit(200)
	}
	color.Info.Println(content) //字符串 - yaml
	color.Debug.Println("使用NACOS加载配置文件")
	viper.SetConfigType("yaml")
	//读取
	if err := viper.ReadConfig(bytes.NewBuffer([]byte(content))); err != nil {
		color.Danger.Println("env read Error = ", err.Error(), "运行中断")
		fmt.Println(err.Error())
		os.Exit(200)
	}

	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}
}

//
//  Curl 发送http请求
//  @Description: 发送http请求
//  @param method 请求方法
//  @param url
//  @param query 请求参数 支持（string,map[string]string）
//  @param headers 请求头
//  @return []byte
//  @return error
//
func Curl(method, url string, query interface{}, headers map[string]string) ([]byte, error) {
	methodList := map[string]interface{}{"post": nil, "get": nil}
	s := strings.ToLower(method)
	if _, ok := methodList[s]; !ok {
		return nil, errors.New("请求方法必须是post/get")
	}
	data := ""
	if query != nil {
		if params, ok := query.(map[string]string); ok {
			i := 0
			for k, v := range params {
				if i > 0 {
					data += "&"
				}
				data += fmt.Sprintf("%s=%s", k, v)
				i++
			}
		} else {
			data, ok = query.(string)
			if !ok {
				return nil, errors.New("参数类型不合法")
			}
		}
	}
	request, err := http.NewRequest(strings.ToUpper(method), url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer request.Body.Close()
	hasUserAgent := false
	if headers != nil {
		for hName, hVal := range headers {
			if strings.ToLower(hName) == "user-agent" {
				hasUserAgent = true
			}
			request.Header.Add(hName, hVal)
		}
	}
	if !hasUserAgent {
		request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")
	}
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		if retry > 0 {
			fmt.Printf("第%d次尝试\n", retry)
			retry = retry - 1
			return Curl(method, url, query, headers)
		}
	}
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return res, nil
}
