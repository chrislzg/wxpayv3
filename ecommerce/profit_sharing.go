package ecommerce

import (
	"encoding/json"
	"net/http"
	"net/url"

	"ptapp.cn/util/wechat.v3/core"
	"ptapp.cn/util/wechat.v3/dto"
)

func (c *payClient) ProfitSharing(req *dto.ProfitSharingReq) (*dto.ProfitSharingResp, error) {
	body, err := c.doRequest(req, core.BuildUrl(nil, nil, core.ApiProfitSharing), http.MethodPost)
	if err != nil {
		return nil, err
	}
	var resp dto.ProfitSharingResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *payClient) ReturnProfitSharing(req *dto.ReturnProfitSharingReq) (*dto.ReturnProfitSharingResp, error) {
	body, err := c.doRequest(req, core.BuildUrl(nil, nil, core.ApiProfitSharingReturn), http.MethodPost)
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

func (c *payClient) HandleProfitSharingNotify(headerSerial string, headerSignature string, headerTimestamp string, headerNonce string, body []byte) (*dto.ProfitSharingNotifyResp, error) {
	if err := c.verify(headerSerial, headerSignature, headerTimestamp, headerNonce, body); err != nil {
		return nil, err
	}
	var notify dto.ProfitSharingNotify
	err := json.Unmarshal(body, &notify)
	if err != nil {
		return nil, err
	}
	if notify.Resource == nil {
		return nil, core.ErrEmptyNotifyResource
	}
	resource := notify.Resource
	resourceDecryptionJson, err := c.Decrypt(resource.Algorithm, resource.Ciphertext, resource.AssociatedData, resource.Nonce)
	if err != nil {
		return nil, err
	}
	var resourceDecryption dto.ProfitSharingNotifyDecryption
	err = json.Unmarshal(resourceDecryptionJson, &resourceDecryption)
	if err != nil {
		return nil, err
	}
	return &dto.ProfitSharingNotifyResp{
		ProfitSharingNotify: &notify,
		Decryption:          &resourceDecryption,
	}, nil
}
func (c *payClient) QueryProfitSharingStatus(req *dto.QueryProfitSharingStatusReq) (*dto.QueryProfitSharingStatusResp, error) {
	query := url.Values{}
	query.Set("sub_mchid", req.SubMchid)
	query.Set("transaction_id", req.TransactionId)
	query.Set("out_order_no", req.OutOrderNo)
	body, err := c.doRequest(req, core.BuildUrl(nil, query, core.ApiProfitSharingStatus), http.MethodGet)
	if err != nil {
		return nil, err
	}
	var resp dto.QueryProfitSharingStatusResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
