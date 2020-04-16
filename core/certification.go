package core

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
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
