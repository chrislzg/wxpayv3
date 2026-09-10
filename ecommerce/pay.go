package ecommerce

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/chrislzg/wxpayv3/core"
	"github.com/chrislzg/wxpayv3/dto"
)

// 普通支付下单
// 支持JSAPI、APP、H5、Native、小程序支付
// JSAPI/APP/小程序返回prepay_id，H5返回h5_url，Native返回code_url
func (c *payClient) PartnerTransactions(req *dto.PartnerTransactionsReq) (*dto.PartnerTransactionsResp, error) {
	req = c.enrichPartnerTransactionsReq(req)
	var url string
	switch req.TradeType {
	case dto.TradeTypeApp:
		url = core.BuildUrl(nil, nil, core.ApiPartnerTransactionApp)
	case dto.TradeTypeJsApi, dto.TradeTypeXCX:
		url = core.BuildUrl(nil, nil, core.ApiPartnerTransactionJsApi)
	case dto.TradeTypeH5:
		url = core.BuildUrl(nil, nil, core.ApiPartnerTransactionH5)
	case dto.TradeTypeNative:
		url = core.BuildUrl(nil, nil, core.ApiPartnerTransactionNative)
	}
	body, err := c.doRequest(req, url, http.MethodPost)
	if err != nil {
		return nil, err
	}
	var ret dto.PartnerTransactionsResp
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (c *payClient) enrichPartnerTransactionsReq(req *dto.PartnerTransactionsReq) *dto.PartnerTransactionsReq {
	if req.SpAppid == "" {
		req.SpAppid = c.appId[req.TradeType]
	}
	if req.SpMchid == "" {
		req.SpMchid = c.mchId
	}
	return req
}

// 按微信订单号查询订单
func (c *payClient) QueryOrderById(req *dto.QueryOrderByIdReq) (*dto.QueryOrderResp, error) {
	params := map[string]string{"transaction_id": req.TransactionId}
	query := url.Values{}
	query.Set("sp_mchid", req.SpMchid)
	query.Set("sub_mchid", req.SubMchid)
	body, err := c.doRequest(nil, core.BuildUrl(params, query, core.ApiPartnerTransactionQueryById), http.MethodGet)
	if err != nil {
		return nil, err
	}
	var resp dto.QueryOrderResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 按商户订单号查询订单
func (c *payClient) QueryOrderByOutTradeNo(req *dto.QueryOrderByOutTradeNoReq) (*dto.QueryOrderResp, error) {
	params := map[string]string{"out_trade_no": req.OutTradeNo}
	query := url.Values{}
	query.Set("sp_mchid", req.SpMchid)
	query.Set("sub_mchid", req.SubMchid)
	body, err := c.doRequest(nil, core.BuildUrl(params, query, core.ApiPartnerTransactionQueryByOutTradeNo), http.MethodGet)
	if err != nil {
		return nil, err
	}
	var resp dto.QueryOrderResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 关闭订单
func (c *payClient) CloseOrder(req *dto.CloseOrderReq) error {
	params := map[string]string{"out_trade_no": req.OutTradeNo}
	_, err := c.doRequest(req, core.BuildUrl(params, nil, core.ApiPartnerTransactionClose), http.MethodPost)
	return err
}
