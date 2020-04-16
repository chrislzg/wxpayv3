package ecommerce

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"ptapp.cn/util/wechat.v3/core"
	"ptapp.cn/util/wechat.v3/dto"
)

// 合单支付
//使用合单支付接口，用户只输入一次密码，即可完成多个订单的支付。目前最多一次可支持10笔订单进行合单支付
//此处官网文档写10笔，但实际电商方案为50笔
//注意：
//• 订单如果需要进行抽佣等，需要在合单中指定需要进行分账（profit_sharing为true）；指定后，交易资金进入二级商户账户，处于冻结状态，可在后续使用分账接口进行分账，利用分账完结进行资金解冻，实现抽佣和对二级商户的账期。
//• 合单中同一个二级商户只允许有一笔子订单。
func (c *payClient) CombineTransactions(req *dto.CombineTransactionsReq) (*dto.CombineTransactionsResp, error) {
	req = c.enrichCombineTransactionReq(req)
	var url string
	switch req.TradeType {
	case dto.TradeTypeApp:
		url = core.BuildUrl(nil, nil, core.ApiCombinedTransactionApp)
	case dto.TradeTypeJsApi, dto.TradeTypeXCX:
		url = core.BuildUrl(nil, nil, core.ApiCombinedTransactionJsApi)
	}
	body, err := c.doRequest(req, url, http.MethodPost)
	if err != nil {
		return nil, err
	}
	var ret dto.CombineTransactionsResp
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (c *payClient) enrichCombineTransactionReq(req *dto.CombineTransactionsReq) *dto.CombineTransactionsReq {
	// 添加combine_appid和combine_mchid信息
	if req.CombineAppid == "" {
		req.CombineAppid = c.appId[req.TradeType]
	}
	if req.CombineMchid == "" {
		req.CombineMchid = c.mchId
	}
	// 如果子单信息中商户号为空，设置为服务商商户号
	if len(req.SubOrders) > 0 {
		for _, subOrder := range req.SubOrders {
			if subOrder.Mchid == "" {
				subOrder.Mchid = c.mchId
			}
		}
	}
	return req
}

func (c *payClient) HandlePaymentNotify(headerSerial string, headerSignature string, headerTimestamp string, headerNonce string, body []byte) (*dto.PaymentNotifyResp, error) {
	err := c.verify(headerSerial, headerSignature, headerTimestamp, headerNonce, body)
	if err != nil {
		return nil, err
	}
	var notify dto.PaymentNotify
	err = json.Unmarshal(body, &notify)
	if err != nil {
		return nil, err
	}
	// 值的检验交给业务方
	//if notify.EventType != "TRANSACTION.SUCCESS" {
	//	return nil, core.ErrEventTypeNotSuccess
	//}
	//if notify.ResourceType != "encrypt-resource" {
	//	return nil, core.ErrIncorrectResourceType
	//}
	if notify.Resource == nil {
		return nil, core.ErrEmptyNotifyResource
	}
	resource := notify.Resource
	decryptedNotifyResourceStr, err := c.Decrypt(resource.Algorithm, resource.Ciphertext, resource.AssociatedData, resource.Nonce)
	if err != nil {
		return nil, err
	}
	var decryptResource dto.PaymentNotifyDecryption
	err = json.Unmarshal(decryptedNotifyResourceStr, &decryptResource)
	if err != nil {
		return nil, err
	}
	return &dto.PaymentNotifyResp{
		PaymentNotify: &notify,
		Decryption:    &decryptResource,
	}, nil
}

// 目前只支持APP、JSAPI、小程序
func (c *payClient) BuildTurnUpPaymentArgumentBody(tradeType dto.TradeType, prepayId string) (string, error) {
	var argument core.TurnUpPaymentArgument

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := core.NonceStr()

	switch tradeType {
	case dto.TradeTypeApp:
		argument = &core.TurnUpPaymentArgumentApp{
			Appid:        c.appId[tradeType],
			Partnerid:    c.mchId,
			Prepayid:     prepayId,
			Noncestr:     nonce,
			Timestamp:    timestamp,
			PackageValue: "Sign=WXPay",
		}
	case dto.TradeTypeJsApi:
		packageKey := "prepay_id=" + prepayId
		argument = &core.TurnUpPaymentArgumentJsApi{
			AppId:     c.appId[tradeType],
			TimeStamp: timestamp,
			Package:   packageKey,
			SignType:  "RSA",
			NonceStr:  nonce,
		}
	case dto.TradeTypeXCX:
		packageKey := "prepay_id=" + prepayId
		argument = &core.TurnUpPaymentArgumentXcx{
			AppId:     c.appId[tradeType],
			TimeStamp: timestamp,
			Package:   packageKey,
			SignType:  "RSA",
			NonceStr:  nonce,
		}
	default:
		return "", errors.New("invalid trade type")
	}
	err := argument.SetPaySign(c.apiPrivateKey)
	if err != nil {
		return "", nil
	}
	argumentBody, err := json.Marshal(argument)
	if err != nil {
		return "", err
	}
	return string(argumentBody), nil
}
