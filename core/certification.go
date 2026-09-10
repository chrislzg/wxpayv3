package core

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func ParsePrivateKey(privateKeyStr string) (*rsa.PrivateKey, error) {
	apiPrivateKeyBlock, _ := pem.Decode([]byte(privateKeyStr))
	if apiPrivateKeyBlock == nil {
		return nil, ErrPemDecodeFailed
	}
	apiPrivateKey, err := x509.ParsePKCS8PrivateKey(apiPrivateKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return apiPrivateKey.(*rsa.PrivateKey), nil
}

func ParseCertification(certKey string) (*x509.Certificate, error) {
	apiCertBlock, _ := pem.Decode([]byte(certKey))
	if apiCertBlock == nil {
		return nil, ErrPemDecodeFailed
	}
	apiCert, err := x509.ParseCertificate(apiCertBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return apiCert, nil
}

// ParsePublicKey 解析微信支付公钥（PEM 格式），支持 PKIX 和 PKCS1 两种格式
func ParsePublicKey(publicKeyStr string) (*rsa.PublicKey, error) {
	publicKeyBlock, _ := pem.Decode([]byte(publicKeyStr))
	if publicKeyBlock == nil {
		return nil, ErrPemDecodeFailed
	}
	pub, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		// 尝试 PKCS1 格式
		if rsaPub, err2 := x509.ParsePKCS1PublicKey(publicKeyBlock.Bytes); err2 == nil {
			return rsaPub, nil
		}
		return nil, err
	}
	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}
	return rsaPub, nil
}
