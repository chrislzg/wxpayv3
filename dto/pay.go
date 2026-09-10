package dto

import "time"

// 结算信息
type SettleInfo struct {
	ProfitSharing bool  `json:"profit_sharing,omitempty"` // 是否指定分账
	SubsidyAmount int64 `json:"subsidy_amount,omitempty"` // 补差金额，SettleInfo.profit_sharing为true时才生效
}

// 普通支付订单金额
type PartnerAmount struct {
	Total    int64  `json:"total"`            // 订单总金额，单位为分
	Currency string `json:"currency,omitempty"` // 货币类型，CNY：人民币
}

// 查询订单返回金额
type QueryOrderAmount struct {
	Total         int64  `json:"total"`            // 订单总金额，单位为分
	PayerTotal    int64  `json:"payer_total"`      // 用户支付金额，单位为分
	Currency      string `json:"currency,omitempty"` // 货币类型，CNY：人民币
	PayerCurrency string `json:"payer_currency,omitempty"` // 用户币种
}

// 普通支付者信息
type PartnerPayerInfo struct {
	SpOpenid  string `json:"sp_openid,omitempty"`  // 用户在服务商AppID下的唯一标识
	SubOpenid string `json:"sub_openid,omitempty"` // 用户在子商户AppID下的唯一标识
}

// 商品详情
type GoodsDetail struct {
	MerchantGoodsId  string `json:"merchant_goods_id"`           // 商户侧商品编码
	WechatpayGoodsId string `json:"wechatpay_goods_id,omitempty"` // 微信支付商品编码
	GoodsName        string `json:"goods_name,omitempty"`        // 商品名称
	Quantity         int64  `json:"quantity"`                    // 商品数量
	UnitPrice        int64  `json:"unit_price"`                  // 商品单价，单位为分
}

// 商品详情信息
type DetailInfo struct {
	CostPrice   int64          `json:"cost_price,omitempty"`    // 订单原价
	InvoiceId   string         `json:"invoice_id,omitempty"`     // 商品小票ID
	GoodsDetail []*GoodsDetail `json:"goods_detail,omitempty"`   // 单品列表
}

// 门店信息
type StoreInfo struct {
	Id       string `json:"id"`                 // 门店编号
	Name     string `json:"name,omitempty"`     // 门店名称
	AreaCode string `json:"area_code,omitempty"` // 地区编码
	Address  string `json:"address,omitempty"`  // 详细地址
}

// 普通支付场景信息
type PartnerSceneInfo struct {
	PayerClientIp string     `json:"payer_client_ip"`        // 用户终端IP
	DeviceId      string     `json:"device_id,omitempty"`    // 商户端设备号
	StoreInfo     *StoreInfo `json:"store_info,omitempty"`    // 商户门店信息
}

// 普通支付下单请求
type PartnerTransactionsReq struct {
	TradeType      TradeType         `json:"-"`
	SpAppid        string            `json:"sp_appid"`                  // 服务商APPID
	SpMchid        string            `json:"sp_mchid"`                  // 服务商户号
	SubAppid       string            `json:"sub_appid,omitempty"`       // 子商户APPID
	SubMchid       string            `json:"sub_mchid"`                 // 子商户号
	Description    string            `json:"description"`              // 商品描述
	OutTradeNo     string            `json:"out_trade_no"`             // 商户订单号
	TimeExpire     string            `json:"time_expire,omitempty"`     // 支付结束时间
	Attach         string            `json:"attach,omitempty"`          // 附加数据
	NotifyUrl      string            `json:"notify_url"`               // 通知地址
	GoodsTag       string            `json:"goods_tag,omitempty"`       // 订单优惠标记
	SettleInfo     *SettleInfo       `json:"settle_info,omitempty"`     // 结算信息
	SupportFapiao  bool              `json:"support_fapiao,omitempty"`  // 电子发票入口开放标识
	Amount         *PartnerAmount    `json:"amount"`                    // 订单金额
	Payer          *PartnerPayerInfo `json:"payer,omitempty"`          // 支付者信息，JSAPI/小程序必填
	Detail         *DetailInfo       `json:"detail,omitempty"`          // 商品详情
	SceneInfo      *PartnerSceneInfo `json:"scene_info,omitempty"`      // 场景信息
}

// 普通支付下单返回结果
// JSAPI/APP/小程序返回prepay_id，H5返回h5_url，Native返回code_url
type PartnerTransactionsResp struct {
	PrepayId string `json:"prepay_id,omitempty"` // 预支付交易会话标识，JSAPI/APP/小程序
	H5Url    string `json:"h5_url,omitempty"`    // H5支付跳转链接
	CodeUrl  string `json:"code_url,omitempty"` // Native支付二维码链接
}

// 查询订单请求（按微信订单号）
type QueryOrderByIdReq struct {
	TransactionId string // 微信支付订单号
	SpMchid        string // 服务商户号
	SubMchid       string // 子商户号
}

// 查询订单请求（按商户订单号）
type QueryOrderByOutTradeNoReq struct {
	OutTradeNo string // 商户订单号
	SpMchid    string // 服务商户号
	SubMchid   string // 子商户号
}

// 查询订单返回结果
type QueryOrderResp struct {
	SpAppid        string            `json:"sp_appid"`                  // 服务商APPID
	SpMchid        string            `json:"sp_mchid"`                  // 服务商户号
	SubAppid       string            `json:"sub_appid,omitempty"`       // 子商户APPID
	SubMchid       string            `json:"sub_mchid"`                 // 子商户号
	OutTradeNo     string            `json:"out_trade_no"`              // 商户订单号
	TransactionId  string            `json:"transaction_id"`           // 微信支付订单号
	TradeType      string            `json:"trade_type"`               // 交易类型
	TradeState     string            `json:"trade_state"`              // 交易状态
	TradeStateDesc string            `json:"trade_state_desc"`         // 交易状态描述
	BankType       string            `json:"bank_type,omitempty"`       // 银行类型
	Attach         string            `json:"attach,omitempty"`          // 附加数据
	SuccessTime    time.Time         `json:"success_time,omitempty"`    // 支付完成时间
	Amount         *QueryOrderAmount `json:"amount,omitempty"`          // 订单金额
	Payer          *PartnerPayerInfo `json:"payer,omitempty"`          // 支付者信息
	SceneInfo      *PartnerSceneInfo `json:"scene_info,omitempty"`      // 场景信息
}

// 关闭订单请求
type CloseOrderReq struct {
	OutTradeNo string `json:"-"` // 商户订单号（路径参数）
	SpMchid    string `json:"sp_mchid"`  // 服务商户号
	SubMchid   string `json:"sub_mchid"` // 子商户号
}
