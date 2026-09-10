package ecommerce

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/chrislzg/wxpayv3/core"
	"github.com/chrislzg/wxpayv3/dto"
)

// 按微信退款单号查询退款
func (c *payClient) QueryRefundById(req *dto.QueryRefundByIdReq) (*dto.QueryRefundResp, error) {
	params := map[string]string{"refund_id": req.RefundId}
	query := url.Values{}
	query.Set("sub_mchid", req.SubMchid)
	body, err := c.doRequest(nil, core.BuildUrl(params, query, core.ApiRefundQueryByRefundId), http.MethodGet)
	if err != nil {
		return nil, err
	}
	var resp dto.QueryRefundResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 按商户退款单号查询退款
func (c *payClient) QueryRefundByOutRefundNo(req *dto.QueryRefundByOutRefundNoReq) (*dto.QueryRefundResp, error) {
	params := map[string]string{"out_refund_no": req.OutRefundNo}
	query := url.Values{}
	query.Set("sub_mchid", req.SubMchid)
	body, err := c.doRequest(nil, core.BuildUrl(params, query, core.ApiRefundQueryByOutRefundNo), http.MethodGet)
	if err != nil {
		return nil, err
	}
	var resp dto.QueryRefundResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
