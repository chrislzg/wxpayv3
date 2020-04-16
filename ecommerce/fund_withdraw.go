package ecommerce

import (
	"encoding/json"
	"net/http"
	"net/url"

	"ptapp.cn/util/wechat.v3/core"
	"ptapp.cn/util/wechat.v3/dto"
)

func (c *payClient) WithdrawFund(req *dto.WithdrawFundReq) (*dto.WithdrawFundResp, error) {
	body, err := c.doRequest(req, core.BuildUrl(nil, nil, core.ApiWithdrawFund), http.MethodPost)
	if err != nil {
		return nil, err
	}
	var resp dto.WithdrawFundResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *payClient) QueryWithdrawalStatus(req *dto.QueryWithdrawalStatusReq) (*dto.QueryWithdrawalStatusResp, error) {
	params := map[string]string{"withdraw_id": req.WithdrawId}
	query := url.Values{}
	query.Set("sub_mchid", req.SubMchid)
	body, err := c.doRequest(nil, core.BuildUrl(params, query, core.ApiWithdrawFundStatus), http.MethodGet)
	if err != nil {
		return nil, err
	}
	var resp dto.QueryWithdrawalStatusResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *payClient) QueryFundBalance(req *dto.QueryFundBalanceReq) (*dto.QueryFundBalanceResp, error) {
	params := map[string]string{"sub_mchid": req.SubMchid}
	body, err := c.doRequest(nil, core.BuildUrl(params, nil, core.ApiFundBalance), http.MethodGet)
	if err != nil {
		return nil, err
	}
	var resp dto.QueryFundBalanceResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
