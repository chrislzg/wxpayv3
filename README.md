# WxPayV3
微信支付的v3版本，GO语言实现

The v3 version of WeChat Pay, implemented in GO language

## Install
`go get -v github.com/chrislzg/wxpay.v3`

## Initialize
```go
package main

import (
"wxpay.v3/core"
"wxpay.v3/dto"
"wxpay.v3/ecommerce"
)

func Test() (core.Client,error) {
    client, err := ecommerce.NewClient(&core.ClientConf{
    AppId:    map[dto.TradeType]string{dto.TradeTypeApp: "${AppId}", dto.TradeTypeJsApi: "${AppId}", dto.TradeTypeXCX: "${AppId}", dto.TradeTypeH5: "${AppId}"},
    MchId:    "${MchId}",
    ApiV3Key: "${APIKEY}",
    ApiCert: &core.ApiCert{
        ApiSerialNo:      "${APISerialNo}",
        ApiPrivateKeyStr: "${apiPrivateKey}",
        ApiCertKey:       "${apiPublickCert}",
    },
    PlatCert: &core.PlatformCert{
        PlatformSerialNo: "${PlatformSerialNo}",
        PlatformCertKey:  "${platformPublicCert}",
    },
    HttpClient: nil,
    })
    if err != nil {
        return nil, err
    }
    return client, nil
}
```

## Usage
```
/// 获取签名Authorization，由认证类型和签名信息组成
Authorization(httpMethod string, urlString string, body []byte) (string, error)
// 获取签名信息
Token(httpMethod string, rawUrl string, body []byte) (string, error)
// 合单支付
CombineTransactions(req *dto.CombineTransactionsReq) (*dto.CombineTransactionsResp, error)
// 提现
WithdrawFund(req *dto.WithdrawFundReq) (*dto.WithdrawFundResp, error)
// 提现状态查询
QueryWithdrawalStatus(req *dto.QueryWithdrawalStatusReq) (*dto.QueryWithdrawalStatusResp, error)
// 查询账户余额
QueryFundBalance(req *dto.QueryFundBalanceReq) (*dto.QueryFundBalanceResp, error)
// 分账
ProfitSharing(req *dto.ProfitSharingReq) (*dto.ProfitSharingResp, error)
// 分账退回
ReturnProfitSharing(req *dto.ReturnProfitSharingReq) (*dto.ReturnProfitSharingResp, error)
// 退款
Refund(req *dto.RefundReq) (*dto.RefundResp, error)
// 验证Response
VerifyResponse(httpStatus int, header *http.Header, body []byte) error
// 验证回调
VerifyNotify(header *http.Header, body []byte) error
// 证书和报文解密
Decrypt(algorithm string, cipherText string, associatedData string, nonce string) ([]byte, error)
// 下载平台证书
Certificate() (*dto.CertificateResp, error)
// 利用api证书私钥对签名串进行签名得到签名值
Sign(message []byte) (string, error)
// 对支付成功通知进行验签和解密
HandlePaymentNotify(headerSerial string, headerSignature string, headerTimestamp string, headerNonce string, body []byte) (*dto.PaymentNotifyResp, error)
// 对分账对账通知验签和解密
HandleProfitSharingNotify(headerSerial string, headerSignature string, headerTimestamp string, headerNonce string, body []byte) (*dto.ProfitSharingNotifyResp, error)
// 对退款通知验签和解密
HandleRefundNotify(headerSerial string, headerSignature string, headerTimestamp string, headerNonce string, body []byte) (*dto.RefundNotifyResp, error)
// 获取客户端调起支付时请求参数
BuildTurnUpPaymentArgumentBody(tradeType dto.TradeType, prepayId string) (string, error)
// 查询分账状态
QueryProfitSharingStatus(req *dto.QueryProfitSharingStatusReq) (*dto.QueryProfitSharingStatusResp, error)
```

## Contact
`email: lzg635935643@qq.com`

## Licence
```
MIT License

Copyright (c) 2020 chrislzg

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```