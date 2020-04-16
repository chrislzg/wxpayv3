package dto

import "time"

type TradeType int8

const (
	_ TradeType = iota
	TradeTypeApp
	TradeTypeJsApi
	TradeTypeXCX
)

// 支付通知event_type字段常量值,具体值
const (
	NotifyEventTypeTransactionType = "TRANSACTION.SUCCESS"
)

// 合单支付场景信息
type SceneInfo struct {
	DeviceId      string `json:"device_id,omitempty"` // 商户端设备号,必填：否
	PayerClientIp string `json:"payer_client_ip"`     // 用户终端IP,必填：是
}

// 订单金额
type SubOrderAmount struct {
	TotalAmount int64  `json:"total_amount"` // 子单金额，单位为分
	Currency    string `json:"currency"`     // 符合ISO 4217标准的三位字母代码，人民币：CNY
}

// 合单支付子单信息
type SubOrder struct {
	ProfitSharing bool            `json:"profit_sharing"`   // 是否指定分账
	Amount        *SubOrderAmount `json:"amount"`           // 订单金额，必填：是
	Mchid         string          `json:"mchid"`            // 子单商户号，必填：是
	Attach        string          `json:"attach"`           // 附加信息，必填：是
	OutTradeNo    string          `json:"out_trade_no"`     // 子单商户订单号
	SubMchid      string          `json:"sub_mchid"`        // 子单发起方商户号，必须与发起方appid有绑定关系
	Detail        string          `json:"detail,omitempty"` // 商品详细描述（商品列表）
	Description   string          `json:"description"`      //商品简单描述。需传入应用市场上的APP名字-实际商品名称，天天爱消除-游戏充值
}

// 支付者
type CombinePayerInfo struct {
	Openid string `json:"openid"` // 使用合单appid获取的对应用户openid。是用户在商户appid下的唯一标识。
}

// 合单支付参数
type CombineTransactionsReq struct {
	TradeType         TradeType         `json:"-"`
	SceneInfo         *SceneInfo        `json:"scene_info,omitempty"`         // 支付场景信息描述
	CombinePayerInfo  *CombinePayerInfo `json:"combine_payer_info,omitempty"` // 非必填，App支付没有OpenId所以该字段不传
	CombineAppid      string            `json:"combine_appid" `               // 该字段在URL中传参
	CombineMchid      string            `json:"combine_mchid"`                // 合单发起方商户号
	CombineOutTradeNo string            `json:"combine_out_trade_no"`         // 合单支付总订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一
	TimeStart         string            `json:"time_start,omitempty"`         // 订单生成时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。
	TimeExpire        string            `json:"time_expire,omitempty"`        // 订单失效时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。
	NotifyUrl         string            `json:"notify_url"`                   // 接收微信支付异步通知回调地址，通知url必须为直接可访问的URL，不能携带参数。 格式: URL
	LimitPay          []string          `json:"limit_pay,omitempty"`          // 指定支付方式
	SubOrders         []*SubOrder       `json:"sub_orders"`                   // 最多支持子单条数：50
}

// 合单支付返回结果
type CombineTransactionsResp struct {
	PrepayId string `json:"prepay_id"`
}

type PaymentNotify struct {
	Resource     *PaymentNotifyResource `json:"resource"`      // 通知资源数据 json格式
	ID           string                 `json:"id"`            // 通知的唯一ID
	CreateTime   string                 `json:"create_time"`   // 通知创建的时间，格式为yyyyMMddHHmmss
	ResourceType string                 `json:"resource_type"` // 通知的资源数据类型，支付成功通知为encrypt-resource
	EventType    string                 `json:"event_type"`    // 通知的类型，支付成功通知的类型为TRANSACTION.SUCCESS
}

type PaymentNotifyResource struct {
	Algorithm      string `json:"algorithm"`       // 对开启结果数据进行加密的加密算法，目前只支持AEAD_AES_256_GCM
	Ciphertext     string `json:"ciphertext"`      // Base64编码后的开启/停用结果数据密文
	Nonce          string `json:"nonce"`           // 附加数据
	AssociatedData string `json:"associated_data"` // 加密使用的随机串
}

type PaymentNotifyDecryption struct {
	SceneInfo         *SceneInfo               `json:"scene_info"` // 支付场景信息描述
	CombinePayerInfo  *CombinePayerInfo        `json:"combine_payer_info"`
	CombineAppid      string                   `json:"combine_appid"`        // 合单发起方的Appid
	CombineOutTradeNo string                   `json:"combine_out_trade_no"` // 合单支付总订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一
	CombineMchid      string                   `json:"combine_mchid"`        // 合单发起方商户号
	SubOrders         []*PaymentNotifySubOrder `json:"sub_orders"`           // 最多支持子单条数：50
}

type PaymentNotifyAmount struct {
	TotalAmount   int64  `json:"total_amount"`   // 子单金额，单位为分
	PayerAmount   int64  `json:"payer_amount"`   // 订单现金支付金额
	Currency      string `json:"currency"`       // 符合ISO 4217标准的三位字母代码，人民币：CNY
	PayerCurrency string `json:"payer_currency"` // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY
}

type PaymentNotifySubOrder struct {
	Amount        *PaymentNotifyAmount `json:"amount"`         // 订单金额信息
	Mchid         string               `json:"mchid"`          // 子单发起方商户号，必须与发起方Appid有绑定关系
	TradeType     string               `json:"trade_type"`     // 交易类型
	TradeState    string               `json:"trade_state"`    // 交易状态 枚举值：SUCCESS：支付成功 REFUND：转入退款 NOTPAY：未支付 CLOSED：已关闭 USERPAYING：用户支付中 PAYERROR：支付失败(其他原因，如银行返回失败)
	BankType      string               `json:"bank_type"`      // 银行类型，采用字符串类型的银行标识
	Attach        string               `json:"attach"`         // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用。
	TransactionID string               `json:"transaction_id"` // 微信支付订单号
	OutTradeNo    string               `json:"out_trade_no"`   // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。 字符字节限制: [6, 32]
	SubMchid      string               `json:"sub_mchid"`      // 二级商户商户号，由微信支付生成并下发。
	SuccessTime   time.Time            `json:"success_time"`   // 订单支付时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss:sss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss:sss表示时分秒毫秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35.120+08:00表示，北京时间2015年5月20日 13点29分35秒。
}

type PaymentNotifyResp struct {
	*PaymentNotify
	Decryption *PaymentNotifyDecryption `json:"decryption"` // 解码后的信息
}
