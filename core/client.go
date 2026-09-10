package core

import (
	"net/http"

	"github.com/chrislzg/wxpayv3/dto"
)

type ApiCert struct {
	ApiSerialNo      string
	ApiPrivateKeyStr string
	ApiCertKey       string
}
type PlatformCert struct {
	PlatformSerialNo string
	PlatformCertKey  string
}

type ClientConf struct {
	AppId      map[dto.TradeType]string // App平台的AppId
	MchId      string
	ApiV3Key   string
	ApiCert    *ApiCert
	PlatCert   *PlatformCert // 可选项，在开始没有平台证书的时候可以不设置，当获取到以后重新设置
	HttpClient *http.Client
}
type Client interface {
	/// 获取签名Authorization，由认证类型和签名信息组成
	Authorization(httpMethod string, urlString string, body []byte) (string, error)
	// 获取签名信息
	Token(httpMethod string, rawUrl string, body []byte) (string, error)
	// 更新平台证书,获取到证书后调用该方法更新平台证书
	UpdatePlatformCert(cert *PlatformCert) error
	// 合单支付
	CombineTransactions(req *dto.CombineTransactionsReq) (*dto.CombineTransactionsResp, error)
	// 提现
	WithdrawFund(req *dto.WithdrawFundReq) (*dto.WithdrawFundResp, error)
	// 提现状态查询
	QueryWithdrawalStatus(req *dto.QueryWithdrawalStatusReq) (*dto.QueryWithdrawalStatusResp, error)
	// 查询账户余额
	QueryFundBalance(req *dto.QueryFundBalanceReq) (*dto.QueryFundBalanceResp, error)
	// 分账
	ProfitSharing(req *dto.ProfitSharingReq) (*dto.ProfitSharingResp, error)
	// 分账退回
	ReturnProfitSharing(req *dto.ReturnProfitSharingReq) (*dto.ReturnProfitSharingResp, error)
	// 退款
	Refund(req *dto.RefundReq) (*dto.RefundResp, error)
	// 验证Response
	VerifyResponse(httpStatus int, header *http.Header, body []byte) error
	// 验证回调
	VerifyNotify(header *http.Header, body []byte) error
	// 证书和报文解密
	Decrypt(algorithm string, cipherText string, associatedData string, nonce string) ([]byte, error)
	// 下载平台证书
	Certificate() (*dto.CertificateResp, error)
	// 利用api证书私钥对签名串进行签名得到签名值
	Sign(message []byte) (string, error)
	// 对支付成功通知进行验签和解密
	HandlePaymentNotify(headerSerial string, headerSignature string, headerTimestamp string, headerNonce string, body []byte) (*dto.PaymentNotifyResp, error)
	// 对分账对账通知验签和解密
	HandleProfitSharingNotify(headerSerial string, headerSignature string, headerTimestamp string, headerNonce string, body []byte) (*dto.ProfitSharingNotifyResp, error)
	// 对退款通知验签和解密
	HandleRefundNotify(headerSerial string, headerSignature string, headerTimestamp string, headerNonce string, body []byte) (*dto.RefundNotifyResp, error)
	// 获取客户端调起支付时请求参数
	BuildTurnUpPaymentArgumentBody(tradeType dto.TradeType, prepayId string) (string, error)
	// 查询分账状态
	QueryProfitSharingStatus(req *dto.QueryProfitSharingStatusReq) (*dto.QueryProfitSharingStatusResp, error)
	// 普通支付下单
	PartnerTransactions(req *dto.PartnerTransactionsReq) (*dto.PartnerTransactionsResp, error)
	// 按微信订单号查询订单
	QueryOrderById(req *dto.QueryOrderByIdReq) (*dto.QueryOrderResp, error)
	// 按商户订单号查询订单
	QueryOrderByOutTradeNo(req *dto.QueryOrderByOutTradeNoReq) (*dto.QueryOrderResp, error)
	// 关闭订单
	CloseOrder(req *dto.CloseOrderReq) error
	// 按微信退款单号查询退款
	QueryRefundById(req *dto.QueryRefundByIdReq) (*dto.QueryRefundResp, error)
	// 按商户退款单号查询退款
	QueryRefundByOutRefundNo(req *dto.QueryRefundByOutRefundNoReq) (*dto.QueryRefundResp, error)
	// 查询分账回退
	QueryReturnProfitSharing(req *dto.QueryReturnProfitSharingReq) (*dto.ReturnProfitSharingResp, error)
	// 添加分账接收方
	AddProfitSharingReceiver(req *dto.AddProfitSharingReceiverReq) (*dto.ProfitSharingReceiverResp, error)
	// 删除分账接收方
	DeleteProfitSharingReceiver(req *dto.DeleteProfitSharingReceiverReq) (*dto.ProfitSharingReceiverResp, error)
	// 请求补差
	Subsidy(req *dto.SubsidyReq) (*dto.SubsidyResp, error)
	// 查询补差
	QuerySubsidy(req *dto.QuerySubsidyReq) (*dto.SubsidyResp, error)
	// 补差回退
	ReturnSubsidy(req *dto.ReturnSubsidyReq) (*dto.ReturnSubsidyResp, error)
}
