package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// OrderMain holds the schema definition for the OrderMain entity.
type OrderMain struct {
	ent.Schema
}

// Fields of the OrderMain.
func (OrderMain) Fields() []ent.Field {

	return []ent.Field{

		field.Int32("id").SchemaType(map[string]string{
			dialect.MySQL: "int(10)unsigned", // Override MySQL.
		}).Unique(),

		field.String("order_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default("").Comment("订单号，系统生成"),

		field.Int8("order_status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(0).Comment("订单状态0 等待sku解析 1待匹配物流 2待派单  3待拣货   4待核单 5待打包 6待发货 7已发货 21待取跟踪号 23待获取物流单"),

		field.Int8("exception_type").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(3)unsigned", // Override MySQL.
		}).Comment("异常类型1物流 2sku 3地址 5利润 6缺货 7海外仓异常 8仓库异常 9物流单异常 10 采购异常 11其他异常 201派单 202代发截停"),

		field.String("store_order_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(128)", // Override MySQL.
		}).Default("").Comment("订单号，店铺后台生成"),

		field.String("platform").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default("").Comment("订单店铺的平台"),

		field.String("shop_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default("").Comment("店铺名称"),

		field.String("ship_first_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)", // Override MySQL.
		}).Default("").Comment("收件人first name"),

		field.String("ship_last_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)", // Override MySQL.
		}).Default("").Comment("收件人last name"),

		field.String("ship_company").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("收件人公司名称"),

		field.String("ship_street1").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("收件人地址1"),

		field.String("ship_street2").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("收件人地址2"),

		field.String("ship_city").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}).Default("").Comment("收件人城市"),

		field.String("ship_state").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}).Default("").Comment("收件人省/州"),

		field.String("ship_zip").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)", // Override MySQL.
		}).Default("").Comment("收件人邮编"),

		field.String("ship_country").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)", // Override MySQL.
		}).Default("").Comment("收件人国家(这里国家全称)"),

		field.String("ship_country_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("收件人国家简称(店铺管理有国家简称)"),

		field.String("ship_phone").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)", // Override MySQL.
		}).Default("").Comment("收件人电话"),

		field.String("ship_fax").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("传真"),

		field.String("ship_remark").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)", // Override MySQL.
		}).Default("").Comment("收件人备注"),

		field.String("ship_house").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("收件人门牌号"),

		field.String("ship_certificate_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)", // Override MySQL.
		}).Default("").Comment("收件人证书号"),

		field.String("order_user_email").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}).Default("").Comment("订单用户email"),

		field.String("orders_user_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Comment("客户名称或者唯一ID"),

		field.String("currency").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default("").Comment("订单货币"),

		field.Int32("currencies_id").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Default(0).Comment("订单货币id"),

		field.Float("currency_rate").SchemaType(map[string]string{
			dialect.MySQL: "decimal(8,6)", // Override MySQL.
		}).Default(0.000000).Comment("订单货币当前汇率"),

		field.String("shipping_method").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}).Default("").Comment("系统运输方式"),

		field.String("shipping_method_orig").SchemaType(map[string]string{
			dialect.MySQL: "varchar(128)", // Override MySQL.
		}).Default("").Comment("平台的原始运输方式"),

		field.String("shipping_method_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("运输方式中文名称(主要普源数据id问题)"),

		field.Float("grand_total").SchemaType(map[string]string{
			dialect.MySQL: "decimal(12,2)", // Override MySQL.
		}).Default(0.00).Comment("订单总金额"),

		field.String("track_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(500)", // Override MySQL.
		}).Default("").Comment("跟踪号(空单号）"),

		field.String("follow_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(500)", // Override MySQL.
		}).Default("").Comment("真正的跟踪号"),

		field.Time("date_payment").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Comment("支付时间"),

		field.String("transaction_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(128)", // Override MySQL.
		}).Default("").Comment("交易号"),

		field.Float("transaction_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(8,2)", // Override MySQL.
		}).Default(0.00).Comment("交易费用"),

		field.Time("update_track_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Comment("跟踪号(空单号）上传时间"),

		field.Time("update_follow_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Comment("真正的跟踪号上传时间"),

		field.Float("shipping_cost").SchemaType(map[string]string{
			dialect.MySQL: "decimal(12,2)", // Override MySQL.
		}).Default(0.00).Comment("订单运输成本"),

		field.Float("shipping_cost_ext").SchemaType(map[string]string{
			dialect.MySQL: "decimal(12,2)", // Override MySQL.
		}).Default(0.00).Comment("订单附加运输成本"),

		field.String("taxes_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("税号"),

		field.Float("tax").SchemaType(map[string]string{
			dialect.MySQL: "decimal(12,2)", // Override MySQL.
		}).Default(0.00).Comment("订单税费"),

		field.Float("order_discount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(12,2)", // Override MySQL.
		}).Default(0.00).Comment("平台优惠(lazada卖家平台优惠)"),

		field.Float("insurance").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)", // Override MySQL.
		}).Default(0.00).Comment("保险费用"),

		field.Float("order_weight").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)unsigned", // Override MySQL.
		}).Default(0.00).Comment("订单重量"),

		field.Enum("order_cretate_type").SchemaType(map[string]string{
			dialect.MySQL: "enum('create','download')", // Override MySQL.
		}).Default("download").Values("create", "download").Comment("订单创建类型:create手动创建，download系统下载"),

		field.Int8("order_type").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(0).Comment("订单的类型 0普通单1 合并单 2拆分订单 3重寄订单 4内部订单 5补发单"),

		field.Time("store_created_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Comment("订单店铺创建时间"),

		field.String("sub_orders_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(1048)", // Override MySQL.
		}).Default("").Comment("关联(子)订单号"),

		field.Time("updated_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Comment("更新时间"),

		field.Time("created_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Comment("创建时间"),

		field.String("warehouse_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default("").Comment("仓库id"),

		field.Int8("order_list_type").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(0).Comment("订单列表类型 默认0正常 1可合并 2作废 3合并被取消"),

		field.String("order_batch_no").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default("").Comment("批次号"),

		field.Int32("shop_id").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Default(0).Comment("店铺管理的店铺ID"),

		field.Float("subtotal").SchemaType(map[string]string{
			dialect.MySQL: "decimal(12,2)", // Override MySQL.
		}).Comment("订单金额(不含运费)"),

		field.Float("shipping_free").SchemaType(map[string]string{
			dialect.MySQL: "decimal(12,2)", // Override MySQL.
		}).Comment("订单运费(平台给的)"),

		field.Float("order_cost").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)", // Override MySQL.
		}).Default(0.00).Comment("订单成本"),

		field.Float("order_profit").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)", // Override MySQL.
		}).Default(0.00).Comment("订单利润"),

		field.Time("ship_date").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Comment("发货时间"),

		field.Float("commission").SchemaType(map[string]string{
			dialect.MySQL: "decimal(8,2)", // Override MySQL.
		}).Default(0.00).Comment("佣金"),

		field.Int8("intercept_state").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}).Default(0).Comment("截单状态 0:正式, 1作废成功, 2等待截单 21直发仓待作废 22直发仓待恢复"),

		field.Time("send_order_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Comment("派单时间"),

		field.String("send_order_user").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Comment("派单人,user表的staff_id"),

		field.Int8("provider_type").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(1).Comment("1直发仓 2海外仓 3FBA订单"),

		field.Float("order_pack_cost").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,4)unsigned", // Override MySQL.
		}).Default(0.0000).Comment("包装成本"),

		field.Float("payment_fix_cost").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)", // Override MySQL.
		}).Default(0.00).Comment("收款平台交易费"),

		field.Int8("is_fba_behalf").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(3)unsigned", // Override MySQL.
		}).Default(0).Comment("是否是fba代发单 0 不是 | 1是"),

		field.Int8("fulfill_type").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(0).Comment("0=>未标记,3=>无需标记,1=>系统标记成功,101=>系统标记失败,2=>手工标记成功,201=>手工标记失败,5=>一次标记成功,4=>一次标记失败,7=>二次标记成功,8=>二次标记失败"),

		field.Int8("fulfill_exception").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(0).Comment("标记异常 0 无异常 1失败异常"),

		field.String("ioss_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default("").Comment("欧盟IOSS编号"),

		field.Int32("warehouse_shop_id").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Comment("仓库店铺id"),

		field.String("sales_record_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(20)", // Override MySQL.
		}).Default("0").Comment("ebay销售记录编号和shopify订单和shopee包裹号"),

		field.String("ticket_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("仓库单据"),

		field.String("overseas_warehouse_status").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("海外仓状态"),

		field.String("pay_pal_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("pp交易id"),

		field.Int8("is_update").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(0),

		field.Int8("order_sale_state").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(0).Comment("订单状态0 正常,1 停售订单"),

		field.Float("product_cost").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,4)", // Override MySQL.
		}).Default(0.0000).Comment("商品成本"),

		field.Float("head_cost").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,4)", // Override MySQL.
		}).Default(0.0000).Comment("头程成本"),

		field.String("carrier_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("代发承运商,多个逗号分割"),

		field.Float("withheld_tax").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)", // Override MySQL.
		}).Default(0.00).Comment("代缴代扣税金(amazon)"),

		field.Float("other_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)", // Override MySQL.
		}).Default(0.00).Comment("其它费用(amazon)"),

		field.Float("escrow_tax").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)unsigned", // Override MySQL.
		}).Default(0.00).Comment("关税"),

		field.Float("final_product_vat_tax").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)unsigned", // Override MySQL.
		}).Default(0.00).Comment("商品增值税"),

		field.Float("final_shipping_vat_tax").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)unsigned", // Override MySQL.
		}).Default(0.00).Comment("运费增值税"),

		field.Float("shipping_charge").SchemaType(map[string]string{
			dialect.MySQL: "decimal(12,2)", // Override MySQL.
		}).Default(0.00).Comment("FBA买家付的运费"),

		field.Float("voucher_platform").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)", // Override MySQL.
		}).Default(0.00).Comment("卖家平台优惠(lazada平台优惠)"),

		field.Time("over_time_left").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Optional().Comment("第一次标记的最后有效期，只针对速卖通,Shopee"),

		field.Float("shipping_rebate").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)unsigned", // Override MySQL.
		}).Default(0.00).Comment("运输补贴（shopee平台运输补贴）"),

		field.String("label_list").SchemaType(map[string]string{
			dialect.MySQL: "varchar(500)", // Override MySQL.
		}).Default("").Comment("物流面单地址"),

		field.Float("out_real_grand_total").SchemaType(map[string]string{
			dialect.MySQL: "decimal(12,2)", // Override MySQL.
		}).Default(0.00).Comment("销售总金额原币"),

		field.String("bill_list").SchemaType(map[string]string{
			dialect.MySQL: "varchar(500)", // Override MySQL.
		}).Default("").Comment("发票链接地址"),

		field.Float("out_grand_total").SchemaType(map[string]string{
			dialect.MySQL: "decimal(12,2)", // Override MySQL.
		}).Default(0.00).Comment("销售总金额"),

		field.String("deliver_shop").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("店铺id（逗号隔开，只有缺货匹配有这个字段）	"),

		field.Float("shipping_cost_final").SchemaType(map[string]string{
			dialect.MySQL: "decimal(12,2)", // Override MySQL.
		}).Default(0.00).Comment("订单最终运输成本"),

		field.Float("shipping_free_about").SchemaType(map[string]string{
			dialect.MySQL: "decimal(12,2)", // Override MySQL.
		}).Default(0.00).Comment("FBA订单预估尾程"),

		field.Int8("create_order_status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(0).Comment("订单创建状态 0 正常可以处理的订单1 待提交运营审核2 运营审核中3 财务审核中 4 驳回"),

		field.Time("finance_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Comment("财务结算(推送时间)"),

		field.Int8("is_qcc").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(2)", // Override MySQL.
		}).Default(2).Comment("是否QCC订单1 是 | 2 不是"),

		field.String("reject_reason").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("手工单订单驳回原因"),

		field.Time("order_local_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Comment("订单创建当地时间"),

		field.String("discount_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(52)", // Override MySQL.
		}).Default("").Comment("优惠码"),

		field.Int8("stockout_status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(0).Comment("缺货状态: 0默认 1待补货 2未补货"),

		field.Float("order_length").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)", // Override MySQL.
		}).Default(0.00).Comment("订单长"),

		field.Float("order_width").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)", // Override MySQL.
		}).Default(0.00).Comment("订单宽"),

		field.Float("order_height").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)", // Override MySQL.
		}).Default(0.00).Comment("订单高"),

		field.Float("first_side").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)", // Override MySQL.
		}).Default(0.00).Comment("包裹三边1"),

		field.Float("second_side").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)", // Override MySQL.
		}).Default(0.00).Comment("包裹三边2"),

		field.Float("third_side").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)", // Override MySQL.
		}).Default(0.00).Comment("包裹三边3"),

		field.Float("order_fee_weight").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,3)", // Override MySQL.
		}).Default(0.000).Comment("订单计费重"),

		field.String("shop_manager").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default("").Comment("店铺负责人"),

		field.Int8("push_status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)unsigned", // Override MySQL.
		}).Default(0).Comment("推送状态：0无面单 1待推送 2待同步 3已同步"),

		field.Float("buyer_paid_shipping_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,0)", // Override MySQL.
		}).Optional().Comment("买家支付运费(shopee)"),
	}

}

// Edges of the OrderMain.
func (OrderMain) Edges() []ent.Edge {
	return nil
}
