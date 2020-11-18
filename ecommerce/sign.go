package ecommerce

import (
	"fmt"
	"time"

	"github.com/golang/glog"

	"github.com/chrislzg/wxpayv3/core"
)

const CertificationType = "WECHATPAY2-SHA256-RSA2048"

// 获取WechatPayV3的header信息Authorization
// 格式如下：
// Authorization: 认证类型 签名信息
func (c *payClient) Authorization(httpMethod string, urlString string, body []byte) (string, error) {
	token, err := c.Token(httpMethod, urlString, body)
	if err != nil {
		glog.Error("get token failed", err)
		return "", err
	}
	return CertificationType + " " + token, nil
}

/*
获取签名信息
请求方法为GET时，报文主体为空。
当请求方法为POST或PUT时，请使用真实发送的JSON报文。
图片上传API，请使用meta对应的JSON报文。
*/
func (c *payClient) Token(httpMethod string, rawUrl string, body []byte) (string, error) {
	nonce := core.NonceStr()
	timestamp := time.Now().Unix()
	message, err := core.BuildMessage(httpMethod, rawUrl, body, nonce, timestamp)
	if err != nil {
		glog.Errorf("buildMessage failed, httpMethod:%v, url:%v, body:%v, nonce:%v, %v", httpMethod, rawUrl, body, nonce, err)
		return "", err
	}
	signature, err := c.Sign(message)
	if err != nil {
		glog.Errorf("sign failed, message:%v, %v", message, err)
		return "", err
	}
	return fmt.Sprintf(`mchid="%s",nonce_str="%s",signature="%s",timestamp="%v",serial_no="%s"`, c.mchId, nonce, signature, timestamp, c.apiSerialNo), nil
}

func (c *payClient) Sign(message []byte) (string, error) {
	return core.Sign(message, c.apiPrivateKey)
}
