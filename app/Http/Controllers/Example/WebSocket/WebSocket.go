package WebSocket

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"order/app/Common"
	"order/app/Http/Models/Kit"
	"strings"
	"time"
)

var ch = make(chan string)

// WebSocket参数
var upGrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	CheckOrigin: func(r *http.Request) bool { // 取消ws跨域校验
		return true
	},
}

func getInputSay(ch chan string) {
	var ctx = context.Background()
	pp := Kit.RDB.PSubscribe(ctx, "__keyevent@0__:expired")
	defer pp.Close()
	for msg := range pp.Channel() {
		fmt.Printf("channel=%s message=%s\n", msg.Channel, msg.Payload)
		ch <- msg.Payload
	}
}

func getOutPutSay(ch chan string) {

}

// Ping1 处理WebSocket消息
// ws:// wss://
// 参考：https://blog.csdn.net/qq_17612199/article/details/79601318
func Ping1(ctx *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for { // 防止gin通过协程调用该handler函数，一旦退出函数，ws会被主动销毁

		// 读取数据
		mt, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}

		// 处理消息
		newMsg1 := string(msg) + "@date1=" + Common.GetTimeDate("YmdHis")
		msg = []byte(newMsg1)

		// 写入（发送）ws数据
		err = ws.WriteMessage(mt, msg)
		if err != nil {
			break
		}

		time.Sleep(3 * time.Second)

		newMsg2 := string(msg) + "@date2=" + Common.GetTimeDate("YmdHis")
		msg = []byte(newMsg2)

		err = ws.WriteMessage(mt, msg)
		if err != nil {
			break
		}

		time.Sleep(3 * time.Second)

		newMsg3 := string(msg) + "@date3=" + Common.GetTimeDate("YmdHis")
		msg = []byte(newMsg3)

		err = ws.WriteMessage(mt, msg)
		if err != nil {
			break
		}
		//ch1 := make(chan string)
		go func() {
			for {
				message := <-ch
				substr := "go_database_jwt_cache:"
				fmt.Println(`Message: ` + message)
				index := strings.Index(message, substr)
				if index > -1 {
					userSlice := strings.Split(message, substr)
					userStr := strings.Join(userSlice, "")
					newMsg4 := string(userStr) + "@date=" + Common.GetTimeDate("YmdHis")
					msg = []byte(newMsg4)

					err = ws.WriteMessage(mt, msg)
					if err != nil {
						break
					}
				}
			}
		}()
		go func() {
			var ctx = context.Background()
			pp := Kit.RDB.PSubscribe(ctx, "__keyevent@0__:expired")
			defer pp.Close()
			for msg := range pp.Channel() {
				fmt.Printf("channel=%s message=%s\n", msg.Channel, msg.Payload)
				ch <- msg.Payload
			}
		}()
	}
}
