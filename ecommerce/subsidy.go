package ecommerce

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/chrislzg/wxpayv3/core"
	"github.com/chrislzg/wxpayv3/dto"
)

// 请求补差
func (c *payClient) Subsidy(req *dto.SubsidyReq) (*dto.SubsidyResp, error) {
	body, err := c.doRequest(req, core.BuildUrl(nil, nil, core.ApiSubsidy), http.MethodPost)
	if err != nil {
		return nil, err
	}
	var resp dto.SubsidyResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 查询补差
func (c *payClient) QuerySubsidy(req *dto.QuerySubsidyReq) (*dto.SubsidyResp, error) {
	params := map[string]string{"out_subsidy_no": req.OutSubsidyNo}
	query := url.Values{}
	query.Set("sub_mchid", req.SubMchid)
	body, err := c.doRequest(nil, core.BuildUrl(params, query, core.ApiSubsidyQuery), http.MethodGet)
	if err != nil {
		return nil, err
	}
	var resp dto.SubsidyResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 补差回退
func (c *payClient) ReturnSubsidy(req *dto.ReturnSubsidyReq) (*dto.ReturnSubsidyResp, error) {
	body, err := c.doRequest(req, core.BuildUrl(nil, nil, core.ApiSubsidyReturn), http.MethodPost)
	if err != nil {
		return nil, err
	}
	var resp dto.ReturnSubsidyResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
