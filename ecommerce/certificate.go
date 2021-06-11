package ecommerce

import (
	"crypto/rsa"
	"encoding/json"
	"net/http"

	"github.com/golang/glog"

	"github.com/chrislzg/wxpayv3/core"
	"github.com/chrislzg/wxpayv3/dto"
)

func (c *payClient) Certificate() (*dto.CertificateResp, error) {
	body, err := c.doRequest(nil, core.BuildUrl(nil, nil, core.ApiCertification), http.MethodGet)
	if err != nil {
		return nil, err
	}
	var resp dto.CertificateResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	for _, data := range resp.Data {
		encryptCert := data.EncryptCertificate
		if encryptCert == nil {
			continue
		}
		decryptCert, err := c.Decrypt(encryptCert.Algorithm, encryptCert.Ciphertext, encryptCert.AssociatedData, encryptCert.Nonce)
		if err != nil {
			return nil, err
		}
		data.DecryptCertificate = string(decryptCert)
	}
	return &resp, nil
}
func (c *payClient) UpdatePlatformCert(cert *core.PlatformCert) error {
	if cert == nil || cert.PlatformCertKey == "" {
		return core.ErrEmptyCert
	}
	platformCert, err := core.ParseCertification(cert.PlatformCertKey)
	if err != nil {
		glog.Errorf("Parse PlatformPublicKey failed privateKeyStr:%v, %v", cert.PlatformCertKey, err)
		return err
	}
	c.platformPublicKey = platformCert.PublicKey.(*rsa.PublicKey)
	c.platformSerialNo = cert.PlatformSerialNo
	return nil
}
