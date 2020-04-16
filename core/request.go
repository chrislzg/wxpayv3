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
	ApiProfitSharing            = "/v3/ecommerce/profitsharing/orders"        // 分账请求
	ApiProfitSharingReturn      = "/v3/ecommerce/profitsharing/returnorders"  // 分账回退请求
	ApiProfitSharingStatus      = "/v3/ecommerce/profitsharing/orders"        // 分账状态查询
	ApiWithdrawFund             = "/v3/ecommerce/fund/withdraw"               // 提现请求
	ApiWithdrawFundStatus       = "/v3/ecommerce/fund/withdraw/{withdraw_id}" // 提现状态查询
	ApiRefund                   = "/v3/ecommerce/refunds/apply"               // 退款
	ApiFundBalance              = "/v3/ecommerce/fund/balance/{sub_mchid}"    // 余额查询
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
