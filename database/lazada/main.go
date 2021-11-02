package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"github.com/wjpxxx/lazadago"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	"github.com/wjpxxx/letgo/lib"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	util "order/app/Helpers"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	ProDB       *gorm.DB
	ProLazadaDB *gorm.DB
)

type MysqlConfig struct {
	Host     string `mapstructure:"host" orm:"host"`
	Port     int    `mapstructure:"port" orm:"port"`
	Name     string `mapstructure:"db" orm:"db"`
	User     string `mapstructure:"user" orm:"user"`
	Password string `mapstructure:"password" orm:"password"`
}

type Order_Main struct {
	ID                      int            `orm:"id"`
	OrderCode               string         `orm:"order_code"`
	OrderStatus             int            `orm:"order_status"`
	ExceptionType           int            `orm:"exception_type"`
	StoreOrderCode          string         `orm:"store_order_code"`
	Platform                string         `orm:"platform"`
	ShopName                string         `orm:"shop_name"`
	ShipFirstName           string         `orm:"ship_first_name"`
	ShipLastName            string         `orm:"ship_last_name"`
	ShipCompany             string         `orm:"ship_company"`
	ShipStreet1             string         `orm:"ship_street1"`
	ShipStreet2             string         `orm:"ship_street2"`
	ShipCity                string         `orm:"ship_city"`
	ShipState               string         `orm:"ship_state"`
	ShipZip                 string         `orm:"ship_zip"`
	ShipCountry             string         `orm:"ship_country"`
	ShipCountryCode         string         `orm:"ship_country_code"`
	ShipPhone               string         `orm:"ship_phone"`
	ShipFax                 string         `orm:"ship_fax"`
	ShipRemark              string         `orm:"ship_remark"`
	OrderUserEmail          string         `orm:"order_user_email"`
	OrdersUserID            string         `orm:"orders_user_id"`
	Currency                string         `orm:"currency"`
	CurrenciesID            int            `orm:"currencies_id"`
	CurrencyRate            string         `orm:"currency_rate"`
	ShippingMethod          string         `orm:"shipping_method"`
	ShippingMethodOrig      string         `orm:"shipping_method_orig"`
	ShippingMethodName      string         `orm:"shipping_method_name"`
	GrandTotal              string         `orm:"grand_total"`
	TrackNumber             string         `orm:"track_number"`
	FollowNumber            string         `orm:"follow_number"`
	DatePayment             string         `orm:"date_payment"`
	TransactionNumber       string         `orm:"transaction_number"`
	TransactionFee          string         `orm:"transaction_fee"`
	UpdateTrackTime         string         `orm:"update_track_time"`
	UpdateFollowTime        string         `orm:"update_follow_time"`
	ShippingCost            string         `orm:"shipping_cost"`
	TaxesNumber             string         `orm:"taxes_number"`
	Tax                     string         `orm:"tax"`
	OrderDiscount           string         `orm:"order_discount"`
	Insurance               string         `orm:"insurance"`
	OrderWeight             string         `orm:"order_weight"`
	OrderCretateType        string         `orm:"order_cretate_type"`
	OrderType               string         `orm:"order_type"`
	StoreCreatedAt          string         `orm:"store_created_at"`
	SubOrdersCode           string         `orm:"sub_orders_code"`
	UpdatedAt               string         `orm:"updated_at"`
	CreatedAt               string         `orm:"created_at"`
	WarehouseID             string         `orm:"warehouse_id"`
	OrderListType           int            `orm:"order_list_type"`
	OrderBatchNo            string         `orm:"order_batch_no"`
	ShopID                  int            `orm:"shop_id"`
	Subtotal                string         `orm:"subtotal"`
	ShippingFree            string         `orm:"shipping_free"`
	OrderCost               string         `orm:"order_cost"`
	OrderProfit             string         `orm:"order_profit"`
	ShipDate                string         `orm:"ship_date"`
	Commission              string         `orm:"commission"`
	InterceptState          int            `orm:"intercept_state"`
	SendOrderTime           string         `orm:"send_order_time"`
	SendOrderUser           string         `orm:"send_order_user"`
	ProviderType            int            `orm:"provider_type"`
	OrderPackCost           string         `orm:"order_pack_cost"`
	PaymentFixCost          string         `orm:"payment_fix_cost"`
	IsFbaBehalf             int            `orm:"is_fba_behalf"`
	FulfillType             int            `orm:"fulfill_type"`
	FulfillException        int            `orm:"fulfill_exception"`
	IossNumber              string         `orm:"ioss_number"`
	WarehouseShopID         int            `orm:"warehouse_shop_id"`
	SalesRecordNumber       int            `orm:"sales_record_number"`
	TicketCode              string         `orm:"ticket_code"`
	OverseasWarehouseStatus string         `orm:"overseas_warehouse_status"`
	PayPalID                string         `orm:"pay_pal_id"`
	IsUpdate                int            `orm:"is_update"`
	ProductCost             string         `orm:"product_cost"`
	HeadCost                string         `orm:"head_cost"`
	CarrierCode             string         `orm:"carrier_code"`
	WithheldTax             string         `orm:"withheld_tax"`
	OtherFee                string         `orm:"other_fee"`
	ShippingCharge          string         `orm:"shipping_charge"`
	VoucherPlatform         string         `orm:"voucher_platform"`
	OrderProduct            []OrderProduct `gorm:"foreignKey:OrderCode;references:OrderCode"`
}
type OrderProduct struct {
	ID               int    `orm:"id"`
	OrderCode        string `orm:"order_code"`
	ProductID        string `orm:"product_id"`
	ProductUnitPrice string `orm:"product_unit_price"`
	Subtotal         string `orm:"subtotal"`
	ProductQuantity  int    `orm:"product_quantity"`
	ProductDesc      string `orm:"product_desc"`
	ProductName      string `orm:"product_name"`
	StoreItemNumber  string `orm:"store_item_number"`
	StoreItemURL     string `orm:"store_item_url"`
	StoreOrderItemID string `orm:"store_order_item_id"`
	CreatedAt        string `orm:"created_at"`
	UpdatedAt        string `orm:"updated_at"`
	StoreItemID      string `orm:"store_item_id"`
	IsDelete         int    `orm:"is_delete"`
	CnTitle          string `orm:"cn_title"`
	EnTitle          string `orm:"en_title"`
	DeclarePrice     string `orm:"declare_price"`
	CustomsCode      string `orm:"customs_code"`
}
type LazadaAccount struct {
	ID          int    `orm:"id"`
	AccountName string `orm:"account_name"`
	AccessToken string `orm:"access_token"`
	SiteCode    string `orm:"site_code"`
	ShopName    string `orm:"shop_name"`
}

type OutOrderProductInfo struct {
	ID               int    `json:"id"`
	OrderCode        string `json:"order_code"`
	ProductID        string `json:"product_id"`
	StoreOrderItemID string `json:"store_order_item_id"`
	StoreOpStatus    string `json:"store_op_status"`
}

type OutOrderInfo struct {
	ID             int                   `json:"id"`
	OrderCode      string                `json:"order_code"`
	StoreOrderCode string                `json:"store_order_code"`
	OrderProduct   []OutOrderProductInfo `json:"order_product"`
}

type RequestJonInfo struct {
	OrderCode []interface{} `json:"order_code"`
}

func (g OutOrderInfo) String() string {
	return lib.ObjectToString(g)
}

var lazadaInfo chan *OutOrderInfo

func main() {
	r := gin.Default()
	initProDatabase()
	initProLazadaDatabase()
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	r.POST("/api/order/status_info", func(ctx *gin.Context) {
		//获取所有token信息
		//orderIds, _ := ctx.GetPostFormArray("order_code")

		var json RequestJonInfo
		ctx.BindJSON(&json)
		fmt.Println(json)
		//	orderIds := json.OrderCode
		var orderIds []string
		for _, orderCode := range json.OrderCode {
			orderIds = append(orderIds, fmt.Sprintf("%v", orderCode))
		}
		//查询所有
		var orderInfo []Order_Main
		var accountInfo []LazadaAccount
		if result := ProLazadaDB.Find(&accountInfo); result.RowsAffected == 0 {
			zap.S().Infof("lazada账号信息[%s]不存在", orderIds)
			util.FailJson(500, "lazada账号信息[%s]不存在", gin.H{}, gin.H{})(ctx)
			return
		}

		if result := ProDB.Where(map[string]interface{}{"order_code": orderIds}).Preload("OrderProduct").Find(&orderInfo); result.RowsAffected == 0 {
			zap.S().Infof("订单信息[%s]不存在", strings.Join(orderIds, ","))
			util.FailJson(500, fmt.Sprintf("订单信息[%s]不存在", strings.Join(orderIds, ",")), gin.H{}, gin.H{})(ctx)
			return
		}
		var outInfo []*OutOrderInfo

		//done := make(chan bool) //通道

		//执行的 这里要注意  需要指针类型传入  否则会异常
		wg := &sync.WaitGroup{}
		//并发控制 10
		limiter := make(chan bool, 20)
		defer close(limiter)

		response := make(chan *OutOrderInfo, 20)
		wgResponse := &sync.WaitGroup{}
		//var result []string
		//处理结果 接收结果
		go func() {
			wgResponse.Add(1)
			for rc := range response {
				outInfo = append(outInfo, rc)
			}
			wgResponse.Done()
		}()
		for _, orderRow := range orderInfo {
			var outRow OutOrderInfo
			var token string
			var country string
			for _, accountRow := range accountInfo {
				if orderRow.ShopName == accountRow.ShopName {
					token = accountRow.AccessToken
					country = strings.ToLower(accountRow.SiteCode)
				}
			}
			outRow.ID = orderRow.ID
			outRow.OrderCode = orderRow.OrderCode
			outRow.StoreOrderCode = orderRow.StoreOrderCode
			for _, opItem := range orderRow.OrderProduct {
				var outOpRow OutOrderProductInfo
				outOpRow.ID = opItem.ID
				outOpRow.OrderCode = opItem.OrderCode
				outOpRow.ProductID = opItem.ProductID
				outOpRow.StoreOrderItemID = opItem.StoreOrderItemID
				outOpRow.StoreOpStatus = ""
				outRow.OrderProduct = append(outRow.OrderProduct, outOpRow)
			}
			orderId, _ := strconv.ParseInt(orderRow.StoreOrderCode, 10, 64)
			//计数器
			wg.Add(1)
			//	cmd := &LazadaInfo{AccessToken: token, OrderId: orderId, Country: country, OutInfo: outRow, Wg: &wg, Ch: ch}
			//并发控制 20
			limiter <- true
			//发送请求
			go pushLazadaGetOrderItems(token, orderId, country, outRow, wg, response, limiter)
			//go cmd.getLazadaGetOrderItems()
		}
		//}
		//发送任务
		wg.Wait()
		//fmt.Println("发送完毕")
		close(response) //关闭 并不影响接收遍历
		//处理接收结果
		wgResponse.Wait()
		//fmt.Println("请求结束")
		//fmt.Println(result)
		//outInfo = append(outInfo, <-lazadaInfo)
		util.SuccessJson("请求成功", outInfo)(ctx)
		return
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
func consumer(in <-chan string) {
	for num := range in {
		fmt.Println(num)
	}
}

type LazadaInfo struct {
	AccessToken string
	OrderId     int64
	Country     string
	OutInfo     OutOrderInfo
	Wg          *sync.WaitGroup
	Ch          chan string
}

func (_this *LazadaInfo) getLazadaGetOrderItems() {
	color.Danger.Println(<-_this.Ch)
}
func pushLazadaGetOrderItems(AccessToken string, OrderId int64, Country string, OutInfo OutOrderInfo, Wg *sync.WaitGroup, response chan *OutOrderInfo, limiter chan bool) {
	//计数器-1
	defer Wg.Done()
	//AccessToken := _this.AccessToken
	//OrderId := _this.OrderId
	//Country := _this.Country
	//OutInfo := _this.OutInfo
	out := &OutOrderInfo{}
	if AccessToken != "" {
		api := lazadago.NewApi(&lazadaConfig.Config{
			AppKey:      "xxxx",
			AccessToken: AccessToken, //刚开始可以为空字符串
			AppSecret:   "xxxx",
			Country:     Country,
		})
		order := api.GetOrderItems(OrderId)
		if order.Code == "0" && order.Data != nil {
			out.ID = OutInfo.ID
			out.OrderCode = OutInfo.OrderCode
			out.StoreOrderCode = OutInfo.StoreOrderCode
			for _, opItem := range OutInfo.OrderProduct {
				orderId, _ := strconv.ParseInt(opItem.StoreOrderItemID, 10, 64)
				for _, orderItem := range order.Data {
					if orderItem.OrderItemId == orderId {
						opItem.StoreOpStatus = orderItem.Status
						break
					}
				}
				out.OrderProduct = append(out.OrderProduct, opItem)
			}
			//zap.S().Infof("%s", out)
		} else {
			//zap.S().Infof("%s", out)
		}
		//结果数据传入管道
		//response <- fmt.Sprintf("%s", out)
		response <- out
	} else {
		response <- &OutInfo
	}
	//释放一个并发
	<-limiter
}

func initProLazadaDatabase() {
	c := MysqlConfig{
		Host:     "xxxx",
		Port:     3306,
		Name:     "xxxx",
		User:     "xxxx",
		Password: "xxxx",
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.Name)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // 慢 SQL 阈值
			LogLevel:      logger.Error, // Log level
			Colorful:      false,        // 禁用彩色打印
		},
	)

	// 全局模式
	var err error
	ProLazadaDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}

func initProDatabase() {
	c := MysqlConfig{
		Host:     "xxxx",
		Port:     3306,
		Name:     "xxxx",
		User:     "xxxx",
		Password: "xxxx",
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.Name)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // 慢 SQL 阈值
			LogLevel:      logger.Error, // Log level
			Colorful:      false,        // 禁用彩色打印
		},
	)

	// 全局模式
	var err error
	ProDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}
