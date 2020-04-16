package dto

import "time"

// 提现状态
const (
	WithdrawalStatusCreateSuccess = "CREATE_SUCCESS" // 受理成功
	WithdrawalStatusSuccess       = "SUCCESS"        // 提现成功
	WithdrawalStatusFAIL          = "FAIL"           // 提现失败
	WithdrawalStatusREFUND        = "REFUND"         // 提现退票
	WithdrawalStatusCLOSE         = "CLOSE"          // 关单
)

type WithdrawFundReq struct {
	Amount       int64  `json:"amount"`              // 提现金额（单位：分）
	SubMchid     string `json:"sub_mchid"`           // 电商平台二级商户号，由微信支付生成并下发。
	OutRequestNo string `json:"out_request_no"`      // 商户提现单号,必须是字母数字
	Remark       string `json:"remark,omitempty"`    // 商户对提现单的备注
	BankMemo     string `json:"bank_memo,omitempty"` // 展示在收款银行系统中的附言，数字、字母最长32个汉字（能否成功展示依赖银行系统支持）
}

type WithdrawFundResp struct {
	SubMchid     string `json:"sub_mchid"`      // 电商平台二级商户号，由微信支付生成并下发。
	WithdrawID   string `json:"withdraw_id"`    // 电商平台提交二级商户提现申请后，由微信支付返回的申请单号，作为查询申请状态的唯一标识
	OutRequestNo string `json:"out_request_no"` // 商户提现单号
}

type QueryWithdrawalStatusReq struct {
	SubMchid   string `json:"sub_mchid"`   // 电商平台二级商户号，由微信支付生成并下发。
	WithdrawId string `json:"withdraw_id"` // 电商平台提交二级商户提现申请后，由微信支付返回的申请单号，作为查询申请状态的唯一标识
}

type QueryWithdrawalStatusResp struct {
	Amount       int64     `json:"amount"`              // 提现金额单位(分)
	SubMchid     string    `json:"sub_mchid,omitempty"` // 电商平台二级商户号，由微信支付生成并下发。
	SpMchid      string    `json:"sp_mchid"`            // 电商平台商户号
	Status       string    `json:"status"`              // 提现状态，枚举值： CREATE_SUCCESS：受理成功 SUCCESS：提现成功 FAIL：提现失败 REFUND：提现退票 CLOSE：关单
	WithdrawID   string    `json:"withdraw_id"`         // 电商平台提交二级商户提现申请后，由微信支付返回的申请单号，作为查询申请状态的唯一标识
	OutRequestNo string    `json:"out_request_no"`      // 商户提现单号
	Reason       string    `json:"reason"`              // 提现失败原因
	Remark       string    `json:"remark"`              // 提现备注
	BankMemo     string    `json:"bank_memo"`           // 银行备注
	CreateTime   time.Time `json:"create_time"`         // 发起提现时间，遵循RFC3339标准格式
	UpdateTime   time.Time `json:"update_time"`         // 提现状态更新时间，遵循RFC3339标准格式
}

type QueryFundBalanceReq struct {
	SubMchid string `json:"sub_mchid"` // 电商平台二级商户号，由微信支付生成并下发。
}

type QueryFundBalanceResp struct {
	AvailableAmount int64  `json:"available_amount"` // 可用余额
	PendingAmount   int64  `json:"pending_amount"`   // 不可用余额
	SubMchid        string `json:"sub_mchid"`        // 电商平台二级商户号，由微信支付生成并下发。
}
