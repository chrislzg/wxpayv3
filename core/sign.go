package core

import (
	"bufio"
	"bytes"
	"crypto"
	rand2 "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

/*
组织签名串信息
签名串格式：
	HTTP请求方法\n
	URL\n
	请求时间戳\n
	请求随机串\n
	请求报文主体\n
*/
func BuildMessage(httpMethod string, urlString string, body []byte, nonceStr string, timestamp int64) ([]byte, error) {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	urlPart := parsedUrl.Path
	if len(parsedUrl.RawQuery) != 0 {
		urlPart = urlPart + "?" + parsedUrl.RawQuery
	}

	buffer := bytes.NewBuffer([]byte{})
	bufw := bufio.NewWriter(buffer)

	bufw.WriteString(httpMethod)
	bufw.WriteByte('\n')
	bufw.WriteString(urlPart)
	bufw.WriteByte('\n')
	bufw.WriteString(strconv.FormatInt(timestamp, 10))
	bufw.WriteByte('\n')
	bufw.WriteString(nonceStr)
	bufw.WriteByte('\n')
	if httpMethod == http.MethodPost || httpMethod == http.MethodPut {
		_, _ = bufw.Write(body)
	}
	bufw.WriteByte('\n')
	_ = bufw.Flush()
	return buffer.Bytes(), nil
}

// 生成随机字符串
func NonceStr() string {
	rand.Seed(time.Now().UnixNano())
	byteLen := 16
	randBytes := make([]byte, byteLen)
	for i := 0; i < byteLen; i++ {
		randBytes[i] = byte(rand.Intn(256))
	}
	return hex.EncodeToString(randBytes)
}

// 利用api证书私钥对签名串进行签名，采用sha256-rsa,对结果进行base64加密
func Sign(message []byte, privateKey *rsa.PrivateKey) (string, error) {
	h := sha256.New()
	h.Write(message)
	signature, err := rsa.SignPKCS1v15(rand2.Reader, privateKey, crypto.SHA256, h.Sum(nil))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}
