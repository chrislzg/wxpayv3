package dto

import "time"

// 请求补差请求
type SubsidyReq struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	TransactionId string `json:"transaction_id"` // 微信支付订单号
	OutSubsidyNo  string `json:"out_subsidy_no"` // 商户补差单号，商户系统内唯一
	Amount        int64  `json:"amount"`         // 补差金额，单位为分
	Description   string `json:"description"`   // 补差原因描述
	NotifyUrl     string `json:"notify_url,omitempty"` // 异步接收补差结果通知的回调地址
}

// 请求补差返回结果
type SubsidyResp struct {
	SubMchid       string    `json:"sub_mchid"`       // 二级商户号
	TransactionId string    `json:"transaction_id"`  // 微信支付订单号
	OutSubsidyNo   string    `json:"out_subsidy_no"`  // 商户补差单号
	SubsidyId      string    `json:"subsidy_id"`      // 微信补差单号
	Amount         int64     `json:"amount"`          // 补差金额
	Status         string    `json:"status"`         // 补差状态，枚举值：PROCESSING：处理中 SUCCESS：补差成功 FAIL：补差失败
	SuccessTime    time.Time `json:"success_time,omitempty"` // 补差成功时间
}

// 查询补差请求
type QuerySubsidyReq struct {
	OutSubsidyNo string // 商户补差单号
	SubMchid     string // 二级商户号
}

// 补差回退请求
type ReturnSubsidyReq struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	OutSubsidyNo  string `json:"out_subsidy_no"` // 商户补差单号
	OutReturnNo   string `json:"out_return_no"`  // 商户补差回退单号，商户系统内唯一
	Amount        int64  `json:"amount"`         // 补差回退金额，单位为分
	Description   string `json:"description"`    // 补差回退原因描述
	NotifyUrl     string `json:"notify_url,omitempty"` // 异步接收补差回退结果通知的回调地址
}

// 补差回退返回结果
type ReturnSubsidyResp struct {
	SubMchid     string    `json:"sub_mchid"`     // 二级商户号
	OutSubsidyNo string    `json:"out_subsidy_no"` // 商户补差单号
	OutReturnNo  string    `json:"out_return_no"` // 商户补差回退单号
	ReturnId     string    `json:"return_id"`     // 微信补差回退单号
	Amount       int64     `json:"amount"`        // 补差回退金额
	Status       string    `json:"status"`        // 补差回退状态，枚举值：PROCESSING：处理中 SUCCESS：补差回退成功 FAIL：补差回退失败
	SuccessTime  time.Time `json:"success_time,omitempty"` // 补差回退成功时间
}
