package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/CodyGuo/dingtalk"
	"github.com/golang/glog"
	"github.com/gookit/color"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/namsral/flag"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
)

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

var nacos NacosConfig

type NacosConfig struct {
	HOST        string
	NamespaceId string
	Username    string
	Password    string
	Port        uint64
	DataId      string
	Group       string
}

// NewNacosConfig 从 viper 中解析配置信息
func NewNacosConfig(endpoint, namespaceId, accessKey, secretKey, dataId, group string, port uint64) NacosConfig {
	nacos.HOST = endpoint
	nacos.NamespaceId = namespaceId
	nacos.Username = accessKey
	nacos.Password = secretKey
	nacos.DataId = dataId
	nacos.Group = group
	nacos.Port = port
	return nacos
}

func GetNacosConfig() NacosConfig {
	return nacos
}

type MyCircularQueue struct {
	Head, Tail   int
	Queue        []string
	Length, size int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Head:   0,
		Tail:   0,
		Queue:  make([]string, k+1, k+1),
		Length: 0,
		size:   k + 1,
	}
}

func (this *MyCircularQueue) EnQueue(value string) bool {
	if this.IsFull() {
		return false
	}
	this.Queue[this.Tail] = value
	this.Tail = (this.Tail + 1 + this.size) % this.size
	this.Length += 1
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.Head = (this.Head + 1 + this.size) % this.size
	this.Length -= 1
	return true
}

func (this *MyCircularQueue) Front() string {
	if this.IsEmpty() {
		return "-1"
	}
	return this.Queue[this.Head]
}

func (this *MyCircularQueue) Rear() string {
	if this.IsEmpty() {
		return "-1"
	}
	return this.Queue[(this.Tail-1+this.size)%this.size]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Head == this.Tail
}

func (this *MyCircularQueue) IsFull() bool {
	if (this.Tail+1)%this.size == this.Head {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsExists(key string) bool {
	for _, v := range this.Queue {
		// fmt.Println(v)
		if v == key {
			return true
		}
	}
	return false
}

func (this *MyCircularQueue) push(key string) bool {
	if this.IsFull() {
		this.DeQueue()
		return this.EnQueue(key)
	} else {
		return this.EnQueue(key)
	}
	return true
}

//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
var endpoint = flag.String("host", "config.nantang-tech.com", "nacos host")

var (
	namespaceId = flag.String("namespace_id", "<namespace_id>", "nacos namespace Id")
	accessKey   = flag.String("username", "", "nacos access username")
	secretKey   = flag.String("password", "", "nacos secret password")
	dataId      = flag.String("data_id", "order-config.yaml", "nacos secret key")
	group       = flag.String("group", "production", "nacos secret key")
	port        = flag.Uint64("port", 80, "nacos port")
	c           YamlConfig
)

func main() {
	flag.Parse()
	NewNacosConfig(*endpoint, *namespaceId, *accessKey, *secretKey, *dataId, *group, *port)
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

	cron2 := cron.New() // 创建一个cron实例
	queueObject := Constructor(10)
	// 执行定时任务（每5秒执行一次）
	err := cron2.AddFunc("*/5 * * * * *", func() {
		//
		var content []Config
		for _, item := range c.Config {
			if ok, _ := Exists(item.Path); ok {
				nowTime := time.Now().Unix()
				lastTime, _ := FileLastUpdateTime(item.Path)
				diffTime := nowTime - lastTime
				if diffTime > item.Time {
					fmt.Printf("%s大于设置时间%v秒\n", item.Path, item.Time)
					// 需要获取推送的订单信息
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
			var cachekey string
			for _, v := range content {
				txt := ""
				txt += "------------------------\n\n"
				txt += fmt.Sprintf(" **监控说明：** %s\n\n", v.Desc)
				txt += fmt.Sprintf(" **监控名称：** %s\n\n", v.Name)
				txt += fmt.Sprintf(" **监控时差：** %v秒\n\n", v.Time)
				txt += fmt.Sprintf(" **监控文件：** %s\n\n", v.Path)
				txt += fmt.Sprintf(" **文件时间：** %s\n\n", v.LastTime)
				txt += fmt.Sprintf(" **当前时间：** %s\n\n", v.NowTime)
				cachekey += fmt.Sprintf(" **监控时差：** %v秒\n\n", v.Time) + fmt.Sprintf(" **监控名称：** %s\n\n", v.Name) + fmt.Sprintf(" **文件时间：** %s\n\n", v.LastTime) + fmt.Sprintf(" **监控文件：** %s\n\n", v.Path) + fmt.Sprintf(" **当前时间：** %v\n\n", time.Now().Day())
				// fmt.Println(cachekey)
				bt.WriteString(txt)
			}
			key := hmacSha256(cachekey, c.Ding.Secretkey)
			if queueObject.IsExists(key) == false {
				secret := c.Ding.Secretkey
				webHook := c.Ding.Webhook
				// markdown类型
				dt := dingtalk.New(webHook, dingtalk.WithSecret(secret))
				markdownTitle := "日志监控通知"
				markdownText := bt.String()
				if err := dt.RobotSendMarkdown(markdownTitle, markdownText); err != nil {
					glog.Fatal(err)
				}
				queueObject.push(key)
			}
		}
	})
	if err != nil {
		fmt.Println(err)
	}
	// 启动/关闭
	cron2.Start()
	defer cron2.Stop()
	select {
	// 查询语句，保持程序运行，在这里等同于for{}
	}
}

//
// timeToData  时间戳转换时间
//  @param path 文件路径
func timeToData(timestamp int64) string {
	timeFormat := "2006-01-02 15:04:05"
	// 时间戳转日期
	t3 := time.Unix(timestamp, 0)
	return t3.Format(timeFormat)
}

//
// FileLastUpdateTime  获取文件最后时间
//  @param path 文件路径
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

//
// Exists  判断文件是否存在
//  @param path 文件路径
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

//
// InitNACOS
// @title  从配置中心读取配置文件
// @description  从配置中心读取配置文件
func InitNACOS() {
	mainDirectory, _ := os.Getwd()
	logFilePath := mainDirectory + "/tmp/nacos/log/"
	logFileName := "nacos.log"
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	// 写入文件
	_, err := os.Stat(fileName)
	if !(err == nil || os.IsExist(err)) {
		var err error
		// 目录不存在则创建
		if _, err = os.Stat(logFilePath); err != nil {
			if err = os.MkdirAll(logFilePath, 0777); err != nil { // 这里如果是0711权限 可能会导致其它线程，读取文件夹内内容出错
				color.Danger.Println("Create log dir err :", err)
			}
		}
		// 创建文件
		if _, err = os.Create(fileName); err != nil {
			color.Danger.Println("Create log file err :", err)
		}
	}
	nacosConf := GetNacosConfig()
	fmt.Println(nacosConf)
	// 创建serverConfig的另一种方式
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			nacosConf.HOST,
			nacosConf.Port,
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos"),
		),
		//*constant.NewServerConfig(
		//	"console2.nacos.io",
		//	80,
		//	constant.WithScheme("http"),
		//	constant.WithContextPath("/nacos")
		//),
	}
	// 创建clientConfig的另一种方式
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId("7bbe2bee-9157-43e1-9b9c-2e19b0a102a6"), // 当namespace是public时，此处填空字符串。
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithRotateTime("1h"),
		constant.WithMaxAge(3),
		constant.WithLogLevel("debug"),
		constant.WithUsername(nacosConf.Username),
		constant.WithPassword(nacosConf.Password),
	)

	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		color.Danger.Println("nacos read Error = ", err.Error(), "运行中断")
		fmt.Println(err.Error())
		os.Exit(200)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: nacosConf.DataId,
		Group:  nacosConf.Group,
	})
	fmt.Println(content)
	fmt.Println(nacosConf.Group)
	if err != nil {
		color.Danger.Println("env read Error = ", err.Error(), "运行中断")
		fmt.Println(err.Error())
		os.Exit(200)
	}
	color.Info.Println(content) // 字符串 - yaml
	color.Debug.Println("使用NACOS加载配置文件")
	viper.SetConfigType("yaml")
	// 读取
	if err := viper.ReadConfig(bytes.NewBuffer([]byte(content))); err != nil {
		color.Danger.Println("env read Error = ", err.Error(), "运行中断")
		fmt.Println(err.Error())
		os.Exit(200)
	}

	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}
	go func() {
		for {
			err := configClient.ListenConfig(vo.ConfigParam{
				DataId: nacosConf.DataId,
				Group:  nacosConf.Group,
				OnChange: func(namespace, group, dataId, data string) {
					// fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
					content, err := configClient.GetConfig(vo.ConfigParam{
						DataId: nacosConf.DataId,
						Group:  nacosConf.Group,
					})
					if err != nil {
						color.Danger.Println("env read Error = ", err.Error(), "运行中断")
						fmt.Println(err.Error())
						os.Exit(200)
					}
					color.Info.Println(content) // 字符串 - yaml
					color.Debug.Println("使用NACOS加载配置文件")
					viper.SetConfigType("yaml")
					// 读取
					if err := viper.ReadConfig(bytes.NewBuffer([]byte(content))); err != nil {
						color.Danger.Println("env read Error = ", err.Error(), "运行中断")
						fmt.Println(err.Error())
						os.Exit(200)
					}

					if err := viper.Unmarshal(&c); err != nil {
						panic(err)
					}
				},
			})
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
}

func hmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
