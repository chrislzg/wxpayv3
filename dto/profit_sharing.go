package dto

import "time"

// 分账动帐通知常量
const (
	NotifyEventTypeProfitsharing       = "PROFITSHARING"
	NotifyEventTypeProfitsharingReturn = "PROFITSHARING_RETURN"
)

// 分账接收方
type ProfitSharingReceiver struct {
	Amount        int64  `json:"amount"`         // 分账金额，单位为分，只能为整数，不能超过原订单支付金额及最大分账比例金额
	ReceiverMchid string `json:"receiver_mchid"` // 只支持电商平台商户和电商平台二级商户，填写微信支付分配的商户号
	Description   string `json:"description"`    // 分账的原因描述，分账账单中需要体现
}

type ProfitSharingReq struct {
	Finish        bool                     `json:"finish"`         // 是否完成分账 1、如果为true，则分账接收商户只支持电商平台商户，且该笔订单剩余未分账的金额会解冻回电商平台二级商户; 2、如果为false，则分账接收商户可以为电商平台商户或者电商平台二级商户，且该笔订单剩余未分账的金额不会解冻回电商平台二级商户，可以对该笔订单再次进行分账。
	SubMchid      string                   `json:"sub_mchid"`      // 分账出资的电商平台二级商户，填写微信支付分配的商户号
	TransactionId string                   `json:"transaction_id"` // 微信支付订单号
	OutOrderNo    string                   `json:"out_order_no"`   // 商户系统内部的分账单号，在商户系统内部唯一（单次分账、多次分账、完结分账应使用不同的商户分账单号），同一分账单号多次请求等同一次
	Receivers     []*ProfitSharingReceiver `json:"receivers"`      // 分账接收方列表，可以设置出资商户作为分账接受方，电商平台模式下，最多可有2个分账接收方
}

type ProfitSharingResp struct {
	SubMchid      string `json:"sub_mchid"`      // 分账出资的电商平台二级商户，填写微信支付分配的商户号
	TransactionID string `json:"transaction_id"` // 微信支付订单号
	OutOrderNo    string `json:"out_order_no"`   // 商户系统内部的分账单号，在商户系统内部唯一（单次分账、多次分账、完结分账应使用不同的商户分账单号），同一分账单号多次请求等同一次
	OrderID       string `json:"order_id"`       // 微信分账单号，微信系统返回的唯一标识
}

type ReturnProfitSharingReq struct {
	Amount      int64  `json:"amount"`
	SubMchid    string `json:"sub_mchid"`              // 分账出资的电商平台二级商户，填写微信支付分配的商户号
	OrderID     string `json:"order_id,omitempty"`     // 微信分账单号，微信系统返回的唯一标识。微信分账单号和商户分账单号二选一填写，与OutOrderNo二选一
	OutOrderNo  string `json:"out_order_no,omitempty"` // 商户系统内部的分账单号，在商户系统内部唯一（单次分账、多次分账、完结分账应使用不同的商户分账单号），同一分账单号多次请求等同一次
	OutReturnNo string `json:"out_return_no"`          // 此回退单号是商户在自己后台生成的一个新的回退单号，在商户后台唯一
	ReturnMchid string `json:"return_mchid"`           // 只能对原分账请求中成功分给商户接收方进行回退
	Description string `json:"description"`            // 需要从分账接收方回退的金额，单位为分，只能为整数，不能超过原始分账单分出给该接收方的金额
}

type ReturnProfitSharingResp struct {
	Amount      int64     `json:"amount"`        // 需要从分账接收方回退的金额，单位为分，只能为整数，不能超过原始分账单分出给该接收方的金额
	SubMchid    string    `json:"sub_mchid"`     // 分账出资的电商平台二级商户，填写微信支付分配的商户号
	OrderID     string    `json:"order_id"`      // 原发起分账请求时，微信返回的微信分账单号，与商户分账单号一一对应。 微信分账单号与商户分账单号二选一填写
	OutOrderNo  string    `json:"out_order_no"`  // 商户系统内部的分账单号，在商户系统内部唯一（单次分账、多次分账、完结分账应使用不同的商户分账单号），同一分账单号多次请求等同一次
	OutReturnNo string    `json:"out_return_no"` // 此回退单号是商户在自己后台生成的一个新的回退单号，在商户后台唯一 只能是数字、大小写字母_-*@ ，同一回退单号多次请求等同一次
	ReturnMchid string    `json:"return_mchid"`  // 只能对原分账请求中成功分给商户接收方进行回退
	ReturnNo    string    `json:"return_no"`     // 微信分账回退单号，微信系统返回的唯一标识
	Result      string    `json:"result"`        // 如果请求返回为处理中，则商户可以通过调用回退结果查询接口获取请求的最终处理结果，枚举值：PROCESSING：处理中 SUCCESS：已成功 FAIL：已失败 注意：如果返回为处理中，请勿变更商户回退单号，使用相同的参数再次发起分账回退，否则会出现资金风险 在处理中状态的回退单如果5天没有成功，会因为超时被设置为已失败
	FailReason  string    `json:"fail_reason"`   // 回退失败的原因，此字段仅回退结果为FAIL时存在，枚举值： ACCOUNT_ABNORMAL：分账接收方账户异常 TIME_OUT_CLOSED:：超时关单
	FinishTime  time.Time `json:"finish_time"`   // 分账回退完成时间，遵循RFC3339标准格式 格式为YYYY-MM-DDTHH:mm:ss:sss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss:sss表示时分秒毫秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35.120+08:00表示，北京时间2015年5月20日 13点29分35秒。
}

type ProfitSharingNotifyResource struct {
	Algorithm      string `json:"algorithm"`       // 对开启结果数据进行加密的加密算法，目前只支持AEAD_AES_256_GCM
	Ciphertext     string `json:"ciphertext"`      // Base64编码后的开启/停用结果数据密文
	Nonce          string `json:"nonce"`           // 附加数据
	AssociatedData string `json:"associated_data"` // 加密使用的随机串
	OriginalType   string `json:"original_type"`   // 加密前的对象类型，分账动账通知的类型为profitsharing
}

type ProfitSharingNotify struct {
	Resource     *ProfitSharingNotifyResource `json:"resource"`      // 通知资源数据
	Id           string                       `json:"id"`            // 通知的唯一ID
	EventType    string                       `json:"event_type"`    // 通知的类型：PROFITSHARING：分账 PROFITSHARING_RETURN：分账回退
	Summary      string                       `json:"summary"`       // 通知简要说明
	ResourceType string                       `json:"resource_type"` // 通知的资源数据类型，支付成功通知为encrypt-resource
	CreateTime   time.Time                    `json:"create_time"`   // 通知创建的时间，Rfc3339标准 格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。
}

type ProfitSharingNotifyResp struct {
	*ProfitSharingNotify
	Decryption *ProfitSharingNotifyDecryption `json:"decryption"` // 解密数据
}

type ProfitSharingNotifyDecryption struct {
	Receiver      *NotifyProfitSharingReceiver `json:"receiver"`       // 分账接收方对象
	Mchid         string                       `json:"mchid"`          // 直连模式分账发起和出资商户
	SpMchid       string                       `json:"sp_mchid"`       // 服务商模式分账发起商户
	SubMchid      string                       `json:"sub_mchid"`      // 服务商模式分账出资商户
	TransactionID string                       `json:"transaction_id"` // 微信支付订单号
	OrderID       string                       `json:"order_id"`       // 微信分账/回退单号
	OutOrderNo    string                       `json:"out_order_no"`   // 分账方系统内部的分账/回退单号
	SuccessTime   time.Time                    `json:"success_time"`   // 成功时间，Rfc3339标准 格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。
}
type NotifyProfitSharingReceiver struct {
	Amount      int64  `json:"amount"`      // 分账动账金额，单位为分，只能为整数
	Type        string `json:"type"`        // MERCHANT_ID：商户ID PERSONAL_WECHATID：个人微信号 PERSONAL_OPENID：个人openid（由父商户APPID转换得到） PERSONAL_SUB_OPENID：个人sub_openid（由子商户APPID转换得到）
	Account     string `json:"account"`     // 1、类型是MERCHANT_ID时，是商户ID 2、类型是PERSONAL_WECHATID时，是个人微信号 3、类型是PERSONAL_OPENID时，是个人openid 4、类型是PERSONAL_SUB_OPENID时，是个人sub_openid
	Description string `json:"description"` // 分账/回退描述
}

type QueryProfitSharingStatusReq struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	TransactionId string `json:"transaction_id"` // 微信支付订单号
	OutOrderNo    string `json:"out_order_no"`   // 商户系统内部的分账单号
}

type QueryProfitSharingStatusResp struct {
	FinishAmount      int64                          `json:"finish_amount"`      // 分账完结的分账金额，单位为分， 仅当查询分账完结的执行结果时，存在本字段
	SubMchid          string                         `json:"sub_mchid"`          // 分账出资的电商平台二级商户，填写微信支付分配的商户号
	TransactionID     string                         `json:"transaction_id"`     // 微信支付订单号
	OutOrderNo        string                         `json:"out_order_no"`       // 商户系统内部的分账单号，在商户系统内部唯一（单次分账、多次分账、完结分账应使用不同的商户分账单号），同一分账单号多次请求等同一次
	OrderID           string                         `json:"order_id"`           // 微信分账单号，微信系统返回的唯一标识
	Status            string                         `json:"status"`             //  分账单状态，枚举值： ACCEPTED：受理成功 PROCESSING：处理中 FINISHED：分账成功 CLOSED：处理失败，已关单
	CloseReason       string                         `json:"close_reason"`       // 关单原因描述，枚举值： NO_AUTH：分账授权已解除
	FinishDescription string                         `json:"finish_description"` // 分账完结的原因描述，仅当查询分账完结的执行结果时，存在本字段
	Receivers         []*ProfitSharingStatusReceiver `json:"receivers"`          // 分账接收方列表，可以设置出资商户作为分账接受方，电商平台模式下，最多可有2个分账接收方。
}

type ProfitSharingStatusReceiver struct {
	Amount        int64     `json:"amount"`         // 分账金额，单位为分，只能为整数，不能超过原订单支付金额及最大分账比例金额
	ReceiverMchid string    `json:"receiver_mchid"` // 只支持电商平台商户和电商平台二级商户，填写微信支付分配的商户号
	Description   string    `json:"description"`    // 分账的原因描述，分账账单中需要体现
	Result        string    `json:"result"`         // 分账结果，枚举值： PENDING：待分账 SUCCESS：分账成功 ADJUST：分账失败待调账 RETURNED：已转回分账方 CLOSED：已关闭
	FailReason    string    `json:"fail_reason"`    // 分账失败原因，枚举值： ACCOUNT_ABNORMAL：分账接收账户异常 NO_RELATION：分账关系已解除
	FinishTime    time.Time `json:"finish_time"`    // 分账完成时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss:sss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss:sss表示时分秒毫秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35.120+08:00表示，北京时间2015年5月20日 13点29分35秒。
}
