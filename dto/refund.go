package dto

import "time"

// 退款通知event_type字段常量值
const (
	NotifyEventTypeRefundSuccess  = "REFUND.SUCCESS"  // 退款成功
	NotifyEventTypeRefundAbnormal = "REFUND.ABNORMAL" // 退款异常
	NotifyEventTypeRefundClosed   = "REFUND.CLOSED"   // 退款关闭
)

type RefundAmount struct {
	Refund   int64  `json:"refund"`             // 退款金额，币种的最小单位，只能为整数，不能超过原订单支付金额
	Total    int64  `json:"total"`              // 原支付交易的订单总金额，币种的最小单位，只能为整数
	Currency string `json:"currency,omitempty"` // 符合ISO 4217标准的三位字母代码，目前只支持人民币：CNY
}

type RefundReq struct {
	TradeType     TradeType     `json:"-"`
	Amount        *RefundAmount `json:"amount"`                   // 订单金额信息
	SubMchid      string        `json:"sub_mchid"`                // 微信支付分配二级商户的商户号
	SpAppid       string        `json:"sp_appid"`                 // 电商平台在微信公众平台申请服务号对应的APPID，申请商户功能的时候微信支付会配置绑定关系
	SubAppid      string        `json:"sub_appid,omitempty"`      // 二级商户在微信申请公众号成功后分配的帐号ID，需要电商平台侧配置绑定关系才能传参
	TransactionID string        `json:"transaction_id,omitempty"` // 原支付交易对应的微信订单号
	OutTradeNo    string        `json:"out_trade_no,omitempty"`   // 原支付交易对应的商户订单号
	OutRefundNo   string        `json:"out_refund_no"`            // 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@，同一退款单号多次请求只退一笔。
	Reason        string        `json:"reason,omitempty"`         // 若商户传入，会在下发给用户的退款消息中体现退款原因
	NotifyURL     string        `json:"notify_url,omitempty"`     // 异步接收微信支付退款结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。 如果参数中传了notify_url，则商户平台上配置的回调地址将不会生效，优先回调当前传的地址。
}

type RefundResp struct {
	Amount          *RefundAmountInfo  `json:"amount"`           // 订单金额信息
	RefundID        string             `json:"refund_id"`        // 微信支付退款订单号
	OutRefundNo     string             `json:"out_refund_no"`    // 商户系统内部的退款单号，商户系统内部唯一，同一退款单号多次请求只退一笔。
	Currency        string             `json:"currency"`         // 符合ISO 4217标准的三位字母代码，目前只支持人民币：CNY
	CreateTime      time.Time          `json:"create_time"`      // 退款受理时间
	PromotionDetail []*PromotionDetail `json:"promotion_detail"` // 优惠退款功能信息
}

type RefundAmountInfo struct {
	Refund         int64  `json:"refund"`          // 退款金额，币种的最小单位，只能为整数，不能超过原订单支付金额。
	PayerRefund    int64  `json:"payer_refund"`    // 退款给用户的金额，不包含所有优惠券金额
	DiscountRefund int64  `json:"discount_refund"` // 优惠券的退款金额，原支付单的优惠按比例退款
	Currency       string `json:"currency"`        // 符合ISO 4217标准的三位字母代码，目前只支持人民币：CNY
}

type PromotionDetail struct {
	Amount       int64  `json:"amount"`        // 优惠券面额,用户享受优惠的金额（优惠券面额=微信出资金额+商家出资金额+其他出资方金额 ）
	RefundAmount int64  `json:"refund_amount"` // 优惠退款金额,代金券退款金额<=退款金额，退款金额-代金券或立减优惠退款金额为现金，说明详见
	PromotionID  string `json:"promotion_id"`  // 券或者立减优惠id
	Scope        string `json:"scope"`         // 优惠范围,GLOBAL- 全场代金券 SINGLE- 单品优惠
	Type         string `json:"type"`          // 优惠类型,COUPON：充值型代金券，商户需要预先充值营销经费 DISCOUNT：免充值型优惠券，商户不需要预先充值营销经费
}

type RefundNotify struct {
	Resource     *ProfitSharingNotifyResource `json:"resource"`
	Id           string                       `json:"id"`
	EventType    string                       `json:"event_type"`
	Summary      string                       `json:"summary"`
	ResourceType string                       `json:"resource_type"`
	CreateTime   time.Time                    `json:"create_time"`
}

type RefundNotifyResource struct {
	Algorithm      string `json:"algorithm"`
	Ciphertext     string `json:"ciphertext"`
	Nonce          string `json:"nonce"`
	AssociatedData string `json:"associated_data"`
	OriginalType   string `json:"original_type"`
}

type RefundNotifyDecryption struct {
	Amount              *RefundNotifyAmount `json:"amount"`
	SpMchid             string              `json:"sp_mchid"`
	SubMchid            string              `json:"sub_mchid"`
	TransactionID       string              `json:"transaction_id"`
	OutTradeNo          string              `json:"out_trade_no"`
	RefundID            string              `json:"refund_id"`
	OutRefundNo         string              `json:"out_refund_no"`
	RefundStatus        string              `json:"refund_status"`
	UserReceivedAccount string              `json:"user_received_account"`
	SuccessTime         time.Time           `json:"success_time"`
}
type RefundNotifyAmount struct {
	Total       int64 `json:"total"`
	Refund      int64 `json:"refund"`
	PayerTotal  int64 `json:"payer_total"`
	PayerRefund int64 `json:"payer_refund"`
}

type RefundNotifyResp struct {
	*RefundNotify
	Decryption *RefundNotifyDecryption `json:"decryption"`
}
