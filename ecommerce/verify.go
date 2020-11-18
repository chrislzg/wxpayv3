package ecommerce

import (
	"bufio"
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/golang/glog"

	"github.com/chrislzg/wxpayv3/core"
)

func (c *payClient) VerifyNotify(header *http.Header, body []byte) error {
	headerSerial := c.getWechatPaySerial(header)
	headerSignature := c.getWechatPaySignature(header)
	headerTimestamp := c.getWechatPayTimestamp(header)
	headerNonce := c.getWechatPayNonce(header)
	return c.verify(headerSerial, headerSignature, headerTimestamp, headerNonce, body)
}

func (c *payClient) VerifyResponse(httpStatus int, header *http.Header, body []byte) error {
	if httpStatus != http.StatusOK && httpStatus != http.StatusNoContent {
		if body == nil {
			return core.ErrInvalidResponse
		}
		var response core.ErrResponseBody
		err := json.Unmarshal(body, &response)
		if err != nil {
			return err
		}
		// 先Unmarshal再赋值，防止被覆盖为空值
		response.HttpStatus = httpStatus
		response.ReqId = header.Get("Request-Id")
		return &response
	}
	headerSerial := c.getWechatPaySerial(header)
	headerSignature := c.getWechatPaySignature(header)
	headerTimestamp := c.getWechatPayTimestamp(header)
	headerNonce := c.getWechatPayNonce(header)
	return c.verify(headerSerial, headerSignature, headerTimestamp, headerNonce, body)
}

func (c *payClient) getWechatPaySerial(header *http.Header) string {
	return header.Get("Wechatpay-Serial")
}

func (c *payClient) getWechatPaySignature(header *http.Header) string {
	return header.Get("Wechatpay-Signature")
}

func (c *payClient) getWechatPayTimestamp(header *http.Header) string {
	return header.Get("Wechatpay-Timestamp")
}

func (c *payClient) getWechatPayNonce(header *http.Header) string {
	return header.Get("Wechatpay-Nonce")
}

/*
获取验签名串
格式为：
	应答时间戳\n
	应答随机串\n
	应答报文主体\n
*/
func (c *payClient) buildVerificationString(timestamp string, nonce string, body []byte) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	bufw := bufio.NewWriter(buffer)
	_, _ = bufw.WriteString(timestamp)
	_ = bufw.WriteByte('\n')
	_, _ = bufw.WriteString(nonce)
	_ = bufw.WriteByte('\n')
	if len(body) != 0 {
		_, _ = bufw.Write(body)
	}
	_ = bufw.WriteByte('\n')
	err := bufw.Flush()
	if err != nil {
		glog.Error("bufw.Flush failed", err)
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (c *payClient) verifySignature(signature string, verificationStr []byte) error {
	h := sha256.New()
	h.Write(verificationStr)
	return rsa.VerifyPKCS1v15(c.platformPublicKey, crypto.SHA256, h.Sum(nil), []byte(signature))
}

func (c *payClient) verify(headerSerial string, headerSignature string, headerTimestamp string, headerNonce string, body []byte) error {
	if headerSerial != c.platformSerialNo {
		glog.Error("wechatPaySerial:", headerSerial)
		return core.ErrInvalidWechatPaySerial
	}
	switch {
	case headerSignature == "":
		return core.ErrInvalidWechatPaySignature
	case headerTimestamp == "":
		return core.ErrInvalidWechatPayTimestamp
	case headerNonce == "":
		return core.ErrInvalidWechatPayNonce
	}
	verificationStr, err := c.buildVerificationString(headerTimestamp, headerNonce, body)
	if err != nil {
		return err
	}
	// 应答header中的signature是base64加密的，所以要先解密
	decodedSignature, err := base64.StdEncoding.DecodeString(headerSignature)
	if err != nil {
		return err
	}
	err = c.verifySignature(string(decodedSignature), verificationStr)
	if err != nil {
		return err
	}
	return nil
}
