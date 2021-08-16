package Order

import (
	"fmt"
	"order/app/Http/Models/Kit"
)

type Order_main struct {
	Id                 int     `orm:"id" json:"id"`
	OrderCode          string  `orm:"order_code" json:"order_code"`                     // 订单号，系统生成
	OrderStatus        int     `orm:"order_status" json:"order_status"`                 // 订单状态0 等待sku解析 1待匹配物流 2待派单  3待拣货   4待核单 5待打包 6待发货 7已发货 21待取跟踪号
	ExceptionType      int     `orm:"exception_type" json:"exception_type"`             // 异常类型1物流 2sku 3地址 5利润 6缺货 201派单
	StoreOrderCode     string  `orm:"store_order_code" json:"store_order_code"`         // 订单号，店铺后台生成
	Platform           string  `orm:"platform" json:"platform"`                         // 订单店铺的平台
	ShopName           string  `orm:"shop_name" json:"shop_name"`                       // 店铺名称
	ShipFirstName      string  `orm:"ship_first_name" json:"ship_first_name"`           // 收件人first name
	ShipLastName       string  `orm:"ship_last_name" json:"ship_last_name"`             // 收件人last name
	ShipCompany        string  `orm:"ship_company" json:"ship_company"`                 // 收件人公司名称
	ShipStreet1        string  `orm:"ship_street1" json:"ship_street1"`                 // 收件人地址1
	ShipStreet2        string  `orm:"ship_street2" json:"ship_street2"`                 // 收件人地址2
	ShipCity           string  `orm:"ship_city" json:"ship_city"`                       // 收件人城市
	ShipState          string  `orm:"ship_state" json:"ship_state"`                     // 收件人省/州
	ShipZip            string  `orm:"ship_zip" json:"ship_zip"`                         // 收件人邮编
	ShipCountry        string  `orm:"ship_country" json:"ship_country"`                 // 收件人国家(这里国家全称)
	ShipCountryCode    string  `orm:"ship_country_code" json:"ship_country_code"`       // 收件人国家简称(店铺管理有国家简称)
	ShipPhone          string  `orm:"ship_phone" json:"ship_phone"`                     // 收件人电话
	ShipFax            string  `orm:"ship_fax" json:"ship_fax"`                         // 传真
	ShipRemark         string  `orm:"ship_remark" json:"ship_remark"`                   // 收件人备注
	OrderUserEmail     string  `orm:"order_user_email" json:"order_user_email"`         // 订单用户email
	OrdersUserId       string  `orm:"orders_user_id" json:"orders_user_id"`             // 客户名称或者唯一ID
	Currency           string  `orm:"currency" json:"currency"`                         // 订单货币
	CurrenciesId       int     `orm:"currencies_id" json:"currencies_id"`               // 订单货币id
	CurrencyRate       float64 `orm:"currency_rate" json:"currency_rate"`               // 订单货币当前汇率
	ShippingMethod     string  `orm:"shipping_method" json:"shipping_method"`           // 系统运输方式
	ShippingMethodOrig string  `orm:"shipping_method_orig" json:"shipping_method_orig"` // 平台的原始运输方式
	ShippingMethodName string  `orm:"shipping_method_name" json:"shipping_method_name"` // 运输方式中文名称(主要普源数据id问题)
	GrandTotal         float64 `orm:"grand_total" json:"grand_total"`                   // 订单总金额
	TrackNumber        string  `orm:"track_number" json:"track_number"`                 // 跟踪号(空单号）
	FollowNumber       string  `orm:"follow_number" json:"follow_number"`               // 真正的跟踪号
	DatePayment        string  `orm:"date_payment" json:"date_payment"`                 // 支付时间
	TransactionNumber  string  `orm:"transaction_number" json:"transaction_number"`     // 交易号
	TransactionFee     float64 `orm:"transaction_fee" json:"transaction_fee"`           // 交易费用
	UpdateTrackTime    string  `orm:"update_track_time" json:"update_track_time"`       // 跟踪号(空单号）上传时间
	UpdateFollowTime   string  `orm:"update_follow_time" json:"update_follow_time"`     // 真正的跟踪号上传时间
	ShippingCost       float64 `orm:"shipping_cost" json:"shipping_cost"`               // 订单运输成本
	TaxesNumber        string  `orm:"taxes_number" json:"taxes_number"`                 // 税号
	Tax                float64 `orm:"tax" json:"tax"`                                   // 订单税费
	OrderDiscount      float64 `orm:"order_discount" json:"order_discount"`             // 平台优惠
	Insurance          float64 `orm:"insurance" json:"insurance"`                       // 保险费用
	OrderWeight        float64 `orm:"order_weight" json:"order_weight"`                 // 订单重量
	OrderCretateType   string  `orm:"order_cretate_type" json:"order_cretate_type"`     // 订单创建类型:create手动创建，download系统下载
	OrderType          int     `orm:"order_type" json:"order_type"`                     // 订单的类型 0普通单1 合并单 2拆分订单 3重寄订单
	StoreCreatedAt     string  `orm:"store_created_at" json:"store_created_at"`         // 订单店铺创建时间
	SubOrdersCode      string  `orm:"sub_orders_code" json:"sub_orders_code"`           // 关联(子)订单号
	UpdatedAt          string  `orm:"updated_at" json:"updated_at"`                     // 更新时间
	CreatedAt          string  `orm:"created_at" json:"created_at"`                     // 创建时间
	WarehouseId        string  `orm:"warehouse_id" json:"warehouse_id"`                 // 仓库id
	OrderListType      int     `orm:"order_list_type" json:"order_list_type"`           // 订单列表类型 默认0正常 1可合并 2作废 3合并被取消
	OrderBatchNo       string  `orm:"order_batch_no" json:"order_batch_no"`             // 批次号
	ShopId             int     `orm:"shop_id" json:"shop_id"`                           // 店铺管理的店铺ID
	Subtotal           float64 `orm:"subtotal" json:"subtotal"`                         // 订单金额(不含运费)
	ShippingFree       float64 `orm:"shipping_free" json:"shipping_free"`               // 订单运费(平台给的)
	OrderCost          float64 `orm:"order_cost" json:"order_cost"`                     // 订单成本
	OrderProfit        float64 `orm:"order_profit" json:"order_profit"`                 // 订单利润
	ShipDate           string  `orm:"ship_date" json:"ship_date"`                       // 发货时间
	Commission         float64 `orm:"commission" json:"commission"`                     // 佣金
	InterceptState     int     `orm:"intercept_state" json:"intercept_state"`           // 截单状态 0:正式, 1作废成功, 2等待截单 21直发仓待作废 22直发仓待恢复
	SendOrderTime      string  `orm:"send_order_time" json:"send_order_time"`           // 派单时间
	SendOrderUser      string  `orm:"send_order_user" json:"send_order_user"`           // 派单人,user表的staff_id
	ProviderType       int     `orm:"provider_type" json:"provider_type"`               // 1直发仓 2海外仓 3FBA订单
	OrderPackCost      float64 `orm:"order_pack_cost" json:"order_pack_cost"`           // 包装成本
	PaymentFixCost     float64 `orm:"payment_fix_cost" json:"payment_fix_cost"`         // 收款平台交易费
	IsFbaBehalf        int     `orm:"is_fba_behalf" json:"is_fba_behalf"`               // 是否是fba代发单 0 不是 | 1是
	FulfillType        int     `orm:"fulfill_type" json:"fulfill_type"`                 // 标记类型 0 未标记 1系统标记 2手工标记
	FulfillException   int     `orm:"fulfill_exception" json:"fulfill_exception"`       // 标记异常 0 无异常 1失败异常
	IossNumber         string  `orm:"ioss_number" json:"ioss_number"`                   // 欧盟IOSS编号
	WarehouseShopId    int     `orm:"warehouse_shop_id" json:"warehouse_shop_id"`       // 仓库店铺id
}

func GetOrder(orderCode string) (orderMian *Order_main, err error) {
	//查询数据库

	//WhereMap := map[string]interface{}{}
	//WhereMap["order_code"] = orderCode

	var result Order_main
	fmt.Println(Kit.DB)
	if Kit.DB != nil {

	}

	Kit.DB.Raw("select * from order_main where order_code = ?", orderCode).Scan(&result)

	return &result, nil
}

func GetOrderList(orderCode string) (orderMian *Order_main, err error) {
	//查询数据库

	//WhereMap := map[string]interface{}{}
	//WhereMap["order_code"] = orderCode

	var result Order_main

	Kit.DB.Raw("select * from order_main where order_code = ?", orderCode).Scan(&result)

	return &result, nil
}
