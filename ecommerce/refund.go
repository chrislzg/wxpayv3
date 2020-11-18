package ecommerce

import (
	"encoding/json"
	"net/http"

	"github.com/chrislzg/wxpayv3/core"
	"github.com/chrislzg/wxpayv3/dto"
)

func (c *payClient) Refund(req *dto.RefundReq) (*dto.RefundResp, error) {
	req = c.enrichRefundReq(req)
	body, err := c.doRequest(req, core.BuildUrl(nil, nil, core.ApiRefund), http.MethodPost)
	if err != nil {
		return nil, err
	}
	var resp dto.RefundResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *payClient) enrichRefundReq(req *dto.RefundReq) *dto.RefundReq {
	if req.SpAppid == "" {
		req.SpAppid = c.appId[req.TradeType]
	}
	return req
}

func (c *payClient) HandleRefundNotify(headerSerial string, headerSignature string, headerTimestamp string, headerNonce string, body []byte) (*dto.RefundNotifyResp, error) {
	if err := c.verify(headerSerial, headerSignature, headerTimestamp, headerNonce, body); err != nil {
		return nil, err
	}
	var notify dto.RefundNotify
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
	var resourceDecryption dto.RefundNotifyDecryption
	err = json.Unmarshal(resourceDecryptionJson, &resourceDecryption)
	if err != nil {
		return nil, err
	}
	return &dto.RefundNotifyResp{
		RefundNotify: &notify,
		Decryption:   &resourceDecryption,
	}, nil
}
