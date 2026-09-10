package core

import (
	"bytes"
	"io"
	"net/http"
	url2 "net/url"
	"strings"
)

const (
	ApiDomain                   = "https://api.mch.weixin.qq.com/"
	ApiCertification            = "/v3/certificates"                          // 平台证书下载
	ApiCombinedTransactionApp   = "/v3/combine-transactions/app"              // App合单支付
	ApiCombinedTransactionJsApi = "/v3/combine-transactions/jsapi"            // JsApi合单支付
	ApiCombinedTransactionH5    = "/v3/combine-transactions/h5"               // H5合单支付
	ApiCombinedTransactionNative = "/v3/combine-transactions/native"          // Native合单支付
	ApiProfitSharing            = "/v3/ecommerce/profitsharing/orders"        // 分账请求
	ApiProfitSharingReturn      = "/v3/ecommerce/profitsharing/returnorders"  // 分账回退请求
	ApiProfitSharingReturnQuery = "/v3/ecommerce/profitsharing/returnorders"  // 分账回退查询
	ApiProfitSharingStatus      = "/v3/ecommerce/profitsharing/orders"        // 分账状态查询
	ApiProfitSharingReceiverAdd    = "/v3/ecommerce/profitsharing/receivers/add"    // 添加分账接收方
	ApiProfitSharingReceiverDelete = "/v3/ecommerce/profitsharing/receivers/delete" // 删除分账接收方
	ApiWithdrawFund             = "/v3/ecommerce/fund/withdraw"               // 提现请求
	ApiWithdrawFundStatus       = "/v3/ecommerce/fund/withdraw/{withdraw_id}" // 提现状态查询
	ApiRefund                   = "/v3/ecommerce/refunds/apply"               // 退款
	ApiRefundQueryByRefundId    = "/v3/ecommerce/refunds/id/{refund_id}"      // 按微信退款单号查询退款
	ApiRefundQueryByOutRefundNo = "/v3/ecommerce/refunds/out-refund-no/{out_refund_no}" // 按商户退款单号查询退款
	ApiFundBalance              = "/v3/ecommerce/fund/balance/{sub_mchid}"    // 余额查询
	ApiSubsidy                  = "/v3/ecommerce/subsidies"                    // 请求补差
	ApiSubsidyQuery             = "/v3/ecommerce/subsidies/out-subsidy-no/{out_subsidy_no}" // 查询补差
	ApiSubsidyReturn            = "/v3/ecommerce/subsidies/return"             // 补差回退
	ApiPartnerTransactionJsApi  = "/v3/pay/partner/transactions/jsapi"         // JSAPI下单
	ApiPartnerTransactionApp    = "/v3/pay/partner/transactions/app"           // APP下单
	ApiPartnerTransactionH5     = "/v3/pay/partner/transactions/h5"            // H5下单
	ApiPartnerTransactionNative = "/v3/pay/partner/transactions/native"        // Native下单
	ApiPartnerTransactionQueryById          = "/v3/pay/partner/transactions/id/{transaction_id}"           // 按微信订单号查询订单
	ApiPartnerTransactionQueryByOutTradeNo = "/v3/pay/partner/transactions/out-trade-no/{out_trade_no}"    // 按商户订单号查询订单
	ApiPartnerTransactionClose  = "/v3/pay/partner/transactions/out-trade-no/{out_trade_no}/close" // 关闭订单
	// 直连商户
	ApiTransactionJsApi                 = "/v3/pay/transactions/jsapi"                                 // 直连JSAPI下单
	ApiTransactionApp                   = "/v3/pay/transactions/app"                                   // 直连APP下单
	ApiTransactionH5                    = "/v3/pay/transactions/h5"                                    // 直连H5下单
	ApiTransactionNative                = "/v3/pay/transactions/native"                                 // 直连Native下单
	ApiTransactionQueryById            = "/v3/pay/transactions/id/{transaction_id}"                   // 直连按微信订单号查询订单
	ApiTransactionQueryByOutTradeNo    = "/v3/pay/transactions/out-trade-no/{out_trade_no}"            // 直连按商户订单号查询订单
	ApiTransactionClose                = "/v3/pay/transactions/out-trade-no/{out_trade_no}/close"       // 直连关闭订单
	ApiRefundDomestic                  = "/v3/refund/domestic/refunds"                                 // 直连退款
	ApiRefundDomesticQueryById         = "/v3/refund/domestic/refunds/{refund_id}"                     // 直连按微信退款单号查询退款
	ApiRefundDomesticQueryByOutRefundNo = "/v3/refund/domestic/refunds/out-refund-no/{out_refund_no}" // 直连按商户退款单号查询退款
)

// @param params url中的path，例:/v3/ecommerce/fund/withdraw/{withdraw_id}中的{withdraw_id}
// @param query url中的query信息
// @subRoutes domain后的路由路径，例：https://api.mch.weixin.qq.com/v3/certificates中的/v3/certificates
func BuildUrl(params map[string]string, query url2.Values, subRoutes ...string) string {
	url := ApiDomain
	for _, route := range subRoutes {
		url += strings.TrimLeft(route, "/")
	}
	for key, param := range params {
		url = strings.ReplaceAll(url, "{"+key+"}", param)
	}
	if query != nil {
		url += "?"
		url += query.Encode()
	}
	return url
}

func NewRequest(authorization string, url string, method string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	return WithBaseHeader(req, authorization), nil
}

func WithBaseHeader(req *http.Request, authorization string) *http.Request {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", authorization)
	return req
}

func SimpleRequest(client *http.Client, url string, method string, authorization string, body []byte) (*http.Response, error) {
	if client == nil {
		client = http.DefaultClient
	}
	var requestBody io.Reader
	if body != nil {
		requestBody = bytes.NewReader(body)
	}
	request, err := NewRequest(authorization, url, method, requestBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
