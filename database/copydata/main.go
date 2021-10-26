package main

import (
	"fmt"
	"github.com/namsral/flag"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"hash/crc32"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

var orderCode = flag.String("order_code", "", "please input order code")
var (
	ProLogDB *gorm.DB
	DevLogDB *gorm.DB
	ProDB    *gorm.DB
	DevDB    *gorm.DB
)

type Order_Main struct {
	ID                      int    `orm:"id"`
	OrderCode               string `orm:"order_code"`
	OrderStatus             int    `orm:"order_status"`
	ExceptionType           int    `orm:"exception_type"`
	StoreOrderCode          string `orm:"store_order_code"`
	Platform                string `orm:"platform"`
	ShopName                string `orm:"shop_name"`
	ShipFirstName           string `orm:"ship_first_name"`
	ShipLastName            string `orm:"ship_last_name"`
	ShipCompany             string `orm:"ship_company"`
	ShipStreet1             string `orm:"ship_street1"`
	ShipStreet2             string `orm:"ship_street2"`
	ShipCity                string `orm:"ship_city"`
	ShipState               string `orm:"ship_state"`
	ShipZip                 string `orm:"ship_zip"`
	ShipCountry             string `orm:"ship_country"`
	ShipCountryCode         string `orm:"ship_country_code"`
	ShipPhone               string `orm:"ship_phone"`
	ShipFax                 string `orm:"ship_fax"`
	ShipRemark              string `orm:"ship_remark"`
	OrderUserEmail          string `orm:"order_user_email"`
	OrdersUserID            string `orm:"orders_user_id"`
	Currency                string `orm:"currency"`
	CurrenciesID            int    `orm:"currencies_id"`
	CurrencyRate            string `orm:"currency_rate"`
	ShippingMethod          string `orm:"shipping_method"`
	ShippingMethodOrig      string `orm:"shipping_method_orig"`
	ShippingMethodName      string `orm:"shipping_method_name"`
	GrandTotal              string `orm:"grand_total"`
	TrackNumber             string `orm:"track_number"`
	FollowNumber            string `orm:"follow_number"`
	DatePayment             string `orm:"date_payment"`
	TransactionNumber       string `orm:"transaction_number"`
	TransactionFee          string `orm:"transaction_fee"`
	UpdateTrackTime         string `orm:"update_track_time"`
	UpdateFollowTime        string `orm:"update_follow_time"`
	ShippingCost            string `orm:"shipping_cost"`
	TaxesNumber             string `orm:"taxes_number"`
	Tax                     string `orm:"tax"`
	OrderDiscount           string `orm:"order_discount"`
	Insurance               string `orm:"insurance"`
	OrderWeight             string `orm:"order_weight"`
	OrderCretateType        string `orm:"order_cretate_type"`
	OrderType               string `orm:"order_type"`
	StoreCreatedAt          string `orm:"store_created_at"`
	SubOrdersCode           string `orm:"sub_orders_code"`
	UpdatedAt               string `orm:"updated_at"`
	CreatedAt               string `orm:"created_at"`
	WarehouseID             string `orm:"warehouse_id"`
	OrderListType           int    `orm:"order_list_type"`
	OrderBatchNo            string `orm:"order_batch_no"`
	ShopID                  int    `orm:"shop_id"`
	Subtotal                string `orm:"subtotal"`
	ShippingFree            string `orm:"shipping_free"`
	OrderCost               string `orm:"order_cost"`
	OrderProfit             string `orm:"order_profit"`
	ShipDate                string `orm:"ship_date"`
	Commission              string `orm:"commission"`
	InterceptState          int    `orm:"intercept_state"`
	SendOrderTime           string `orm:"send_order_time"`
	SendOrderUser           string `orm:"send_order_user"`
	ProviderType            int    `orm:"provider_type"`
	OrderPackCost           string `orm:"order_pack_cost"`
	PaymentFixCost          string `orm:"payment_fix_cost"`
	IsFbaBehalf             int    `orm:"is_fba_behalf"`
	FulfillType             int    `orm:"fulfill_type"`
	FulfillException        int    `orm:"fulfill_exception"`
	IossNumber              string `orm:"ioss_number"`
	WarehouseShopID         int    `orm:"warehouse_shop_id"`
	SalesRecordNumber       int    `orm:"sales_record_number"`
	TicketCode              string `orm:"ticket_code"`
	OverseasWarehouseStatus string `orm:"overseas_warehouse_status"`
	PayPalID                string `orm:"pay_pal_id"`
	IsUpdate                int    `orm:"is_update"`
	ProductCost             string `orm:"product_cost"`
	HeadCost                string `orm:"head_cost"`
	CarrierCode             string `orm:"carrier_code"`
	WithheldTax             string `orm:"withheld_tax"`
	OtherFee                string `orm:"other_fee"`
	ShippingCharge          string `orm:"shipping_charge"`
	VoucherPlatform         string `orm:"voucher_platform"`
}

func (user *Order_Main) BeforeSave(scope *gorm.DB) (err error) {
	//if pw, err := bcrypt.GenerateFromPassword(user.Password, 0); err == nil {
	user.DatePayment = timeToData(GetTimestamp(user.DatePayment))
	user.CreatedAt = timeToData(GetTimestamp(user.CreatedAt))
	user.UpdatedAt = timeToData(GetTimestamp(user.UpdatedAt))
	user.SendOrderTime = timeToData(GetTimestamp(user.SendOrderTime))
	user.StoreCreatedAt = timeToData(GetTimestamp(user.StoreCreatedAt))
	user.ShipDate = timeToData(GetTimestamp(user.ShipDate))
	user.UpdateTrackTime = timeToData(GetTimestamp(user.UpdateTrackTime))
	user.UpdateFollowTime = timeToData(GetTimestamp(user.UpdateFollowTime))
	//fmt.Println(user)
	return nil
}

func timeToData(timestamp int64) string {
	timeFormat := "2006-01-02 15:04:05"
	// 时间戳转日期
	t3 := time.Unix(timestamp, 0)
	return t3.Format(timeFormat)
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

func (user *OrderProduct) BeforeSave(scope *gorm.DB) (err error) {
	//if pw, err := bcrypt.GenerateFromPassword(user.Password, 0); err == nil {
	user.CreatedAt = timeToData(GetTimestamp(user.CreatedAt))
	user.UpdatedAt = timeToData(GetTimestamp(user.UpdatedAt))
	//fmt.Println(user)
	return nil
}

type OrderStatusHistory struct {
	OshId       int    `orm:"osh_id"`
	OrderCode   string `orm:"order_code"`
	OshType     string `orm:"osh_type"`
	OshStatus   string `orm:"osh_status"`
	LogType     int    `orm:"log_type"`
	OshComments string `orm:"osh_comments"`
	OshRemark   string `orm:"osh_remark"`
	OshIp       string `orm:"osh_ip"`
	CreateUser  string `orm:"create_user"`
	CreatedAt   string `orm:"created_at"`
}

func (user *OrderStatusHistory) BeforeSave(scope *gorm.DB) (err error) {
	//if pw, err := bcrypt.GenerateFromPassword(user.Password, 0); err == nil {
	user.CreatedAt = timeToData(GetTimestamp(user.CreatedAt))
	//fmt.Println(user)
	return nil
}

func getTableName(OrderCode string) string {
	codeint := crc32.ChecksumIEEE([]byte(OrderCode))
	prefix := math.Mod(float64(codeint), 300)
	//	fmt.Println(prefix)
	return "order_status_history_" + fmt.Sprintf("%03d", int(prefix))
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" orm:"host"`
	Port     int    `mapstructure:"port" orm:"port"`
	Name     string `mapstructure:"db" orm:"db"`
	User     string `mapstructure:"user" orm:"user"`
	Password string `mapstructure:"password" orm:"password"`
}

//_, err := sqlx.Open("mysql", "nt_order:8Iwi+GEimp3cmwEphIVe@tcp(101.132.43.121:3306)/nt_order")

func main() {
	flag.Parse()
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	zap.S().Infof("开始复制订单号[%s]", *orderCode)
	if *orderCode == "" {
		zap.S().Infof("必填数据[%s]为空,结束复制", "OrderCode")
		os.Exit(200)
	}
	zap.S().Infof("初始化数据库[%s]", "ProDB")
	initProDatabase()
	initProLogDatabase()
	zap.S().Infof("初始化数据库[%s]", "DevDB")
	initTestDatabase()
	initTestLogDatabase()
	zap.S().Infof("开始数据迁移[%s]", time.Now().Format("2006-01-02 15:04:05"))

	orderCodeArr := strings.Split(*orderCode, ",")
	for _, v := range orderCodeArr {
		var orderInfo Order_Main
		var orderProductInfo []OrderProduct
		var orderLogInfo []OrderStatusHistory
		if result := ProLogDB.Table(getTableName(v)).Where(map[string]interface{}{"order_code": v}).Find(&orderLogInfo); result.RowsAffected == 0 {
			zap.S().Infof("订单日志信息[%s]不存在", v)
			os.Exit(200)
		}
		if result := ProDB.Where(map[string]interface{}{"order_code": v}).First(&orderInfo); result.RowsAffected == 0 {
			zap.S().Infof("订单信息[%s]不存在", v)
			os.Exit(200)
		}

		if result := ProDB.Where(map[string]interface{}{"order_code": v}).Find(&orderProductInfo); result.RowsAffected == 0 {
			zap.S().Infof("订单产品信息[%s]不存在", v)
			os.Exit(200)
		}

		if result := DevDB.Save(&orderInfo); result.RowsAffected != 0 {
			zap.S().Infof("复制订单信息[%s]成功", v)
		} else {
			zap.S().Infof("复制订单信息[%s]成功,已经存在", v)
		}

		if result := DevDB.Save(&orderProductInfo); result.RowsAffected != 0 {
			zap.S().Infof("复制订单行信息[%s]成功", v)
		} else {
			zap.S().Infof("复制订单行信息[%s]成功,已经存在", v)
		}

		if result := DevLogDB.Table(getTableName(v)).Save(&orderLogInfo); result.RowsAffected != 0 {
			zap.S().Infof("复制订单日志信息[%s]成功", v)
		} else {
			zap.S().Infof("复制订单日志信息[%s]成功,已经存在", v)
		}
	}

	zap.S().Infof("结束数据迁移[%s]", time.Now().Format("2006-01-02 15:04:05"))

}

func initProDatabase() {
	c := MysqlConfig{
		Host:     "XXXXXXXXXXXX",
		Port:     3306,
		Name:     "nt_order",
		User:     "nt_order",
		Password: "XXXXXXXXXXXXX",
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

func initTestDatabase() {
	c := MysqlConfig{
		Host:     "XXXXXXXXXXXX",
		Port:     3311,
		Name:     "nt_order",
		User:     "test",
		Password: "XXXXXXXXXXXXX",
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
	DevDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	//DevDB.Callback().Create().Replace("orm:updated_at", updateTimeStampForUpdateCallback)
}
func GetTimestamp(change string) int64 {
	times, _ := time.Parse("2006-01-02 15:04:05", change)
	timeUnix := times.Unix()
	return timeUnix
}

func initProLogDatabase() {
	c := MysqlConfig{
		Host:     "XXXXXXXXXXXX",
		Port:     3306,
		Name:     "nt_order_log",
		User:     "nt_order",
		Password: "XXXXXXXXXXXXX",
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
	ProLogDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}

func initTestLogDatabase() {
	c := MysqlConfig{
		Host:     "XXXXXXXXXXXX",
		Port:     3311,
		Name:     "nt_order_log",
		User:     "test",
		Password: "XXXXXXXXXXXXX",
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
	DevLogDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	//DevDB.Callback().Create().Replace("orm:updated_at", updateTimeStampForUpdateCallback)
}
