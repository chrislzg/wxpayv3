package ecommerce

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/chrislzg/wxpayv3/core"
	"github.com/chrislzg/wxpayv3/dto"
)

// 查询分账回退
func (c *payClient) QueryReturnProfitSharing(req *dto.QueryReturnProfitSharingReq) (*dto.ReturnProfitSharingResp, error) {
	query := url.Values{}
	query.Set("sub_mchid", req.SubMchid)
	if req.OrderId != "" {
		query.Set("order_id", req.OrderId)
	}
	if req.OutOrderNo != "" {
		query.Set("out_order_no", req.OutOrderNo)
	}
	if req.OutReturnNo != "" {
		query.Set("out_return_no", req.OutReturnNo)
	}
	if req.ReturnNo != "" {
		query.Set("return_no", req.ReturnNo)
	}
	body, err := c.doRequest(nil, core.BuildUrl(nil, query, core.ApiProfitSharingReturnQuery), http.MethodGet)
	if err != nil {
		return nil, err
	}
	var resp dto.ReturnProfitSharingResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 添加分账接收方
func (c *payClient) AddProfitSharingReceiver(req *dto.AddProfitSharingReceiverReq) (*dto.ProfitSharingReceiverResp, error) {
	body, err := c.doRequest(req, core.BuildUrl(nil, nil, core.ApiProfitSharingReceiverAdd), http.MethodPost)
	if err != nil {
		return nil, err
	}
	var resp dto.ProfitSharingReceiverResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 删除分账接收方
func (c *payClient) DeleteProfitSharingReceiver(req *dto.DeleteProfitSharingReceiverReq) (*dto.ProfitSharingReceiverResp, error) {
	body, err := c.doRequest(req, core.BuildUrl(nil, nil, core.ApiProfitSharingReceiverDelete), http.MethodPost)
	if err != nil {
		return nil, err
	}
	var resp dto.ProfitSharingReceiverResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
