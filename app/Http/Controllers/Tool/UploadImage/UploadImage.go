package UploadImage

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"order/app/Http/Models/Order"
)

func Test(ctx *gin.Context) {
	method := ctx.Request.Method
	body := ctx.Request.Body
	header := ctx.Request.Header["Content-Type"]
	//
	//fmt.Println(ctx.Request.Form)
	//
	//fmt.Println(body)
	orderCode := ctx.Query("order_code")
	user, err := Order.GetOrder(orderCode)
	//fmt.Println(orderCode)
	//fmt.Println(Order.GetOrder(orderCode))
	if err != nil {
		ctx.JSON(200, gin.H{"error": err})
	}
	//fmt.Printf("\n %c[0;48;32m%s%c[0m\n\n", 0x1B, "["+time.Now().Format("2006-01-02 15:04:05")+"]"+orderCode, 0x1B)
	color.Debug.Println("Debug message : " + orderCode)
	id := "1008611" //
	nickname := "456"
	// 接口返回
	back := gin.H{
		"method":   method,
		"body":     body,
		"user":     user,
		"header":   header,
		"id":       id,
		"nickname": nickname,
		//"latency": ctx.Get("state_latency"),
		//"img": img,
		//"excel_name": excelName,
		//"excel_state": excelState,
		////"timezone": Common.ServerInfo["timezone"],
		//"date": Common.GetTimeDate("Y-m-d H:i:s"),

	}
	ctx.JSON(200, back)
}
func signTopRequest(params map[string]string, secret string, signMethod string) {

}
