package ecommerce

import (
	"crypto/rsa"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/golang/glog"

	"github.com/chrislzg/wxpayv3/core"
	"github.com/chrislzg/wxpayv3/dto"
)

type payClient struct {
	appId                map[dto.TradeType]string //appId
	mchId                string                   // 商户号
	apiV3Key             string                   // apiV3密钥
	apiSerialNo          string                   // API证书序列号
	apiPrivateKey        *rsa.PrivateKey          // API证书私钥
	apiPublicKey         *rsa.PublicKey           // API公钥
	platformSerialNo     string                   // 平台证书序列号
	platformPublicKey    *rsa.PublicKey           // 平台证书公钥
	wechatPayPublicKeyId string                   // 微信支付公钥ID
	wechatPayPublicKey   *rsa.PublicKey           // 微信支付公钥
	httpClient           *http.Client
}

func NewClient(conf *core.ClientConf) (core.Client, error) {
	apiPrivateKey, err := core.ParsePrivateKey(conf.ApiCert.ApiPrivateKeyStr)
	if err != nil {
		glog.Errorf("Parse ApiPrivateKey failed: %v", err)
		return nil, err
	}
	apiCert, err := core.ParseCertification(conf.ApiCert.ApiCertKey)
	if err != nil {
		glog.Errorf("Parse ApiPublicKey failed: %v", err)
		return nil, err
	}
	httpClient := conf.HttpClient
	if httpClient == nil {
		httpClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   3 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   3 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
			Timeout: 5 * time.Second,
		}
	}

	c := &payClient{
		appId:         conf.AppId,
		mchId:         conf.MchId,
		apiV3Key:      conf.ApiV3Key,
		apiSerialNo:   conf.ApiCert.ApiSerialNo,
		apiPrivateKey: apiPrivateKey,
		apiPublicKey:  apiCert.PublicKey.(*rsa.PublicKey),
		httpClient:    httpClient,
	}
	if conf.PlatCert != nil {
		err = c.UpdatePlatformCert(conf.PlatCert)
		if err != nil {
			return nil, err
		}
	}
	if conf.WechatPayPublicKey != nil {
		wpk, err := core.ParsePublicKey(conf.WechatPayPublicKey.PublicKey)
		if err != nil {
			return nil, err
		}
		c.wechatPayPublicKey = wpk
		c.wechatPayPublicKeyId = conf.WechatPayPublicKey.PublicKeyId
	}
	return c, nil
}

func (c *payClient) doRequest(requestData interface{}, url string, httpMethod string) ([]byte, error) {
	var data []byte
	if requestData != nil {
		var err error
		data, err = json.Marshal(requestData)
		if err != nil {
			return nil, err
		}
	}
	// 重试3次，避免因网络原因导致失败；每次重试重新生成 authorization 防止签名时间戳过期
	retryTimes := 3
	var resp *http.Response
	var err error
	for i := 0; i < retryTimes; i++ {
		var authorization string
		authorization, err = c.Authorization(httpMethod, url, data)
		if err != nil {
			return nil, err
		}
		resp, err = core.SimpleRequest(c.httpClient, url, httpMethod, authorization, data)
		if err != nil {
			continue
		}
		break
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = c.VerifyResponse(resp.StatusCode, &resp.Header, body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
