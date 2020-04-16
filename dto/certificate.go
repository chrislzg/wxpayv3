package dto

import "time"

type CertificateResp struct {
	Data []*CertificateData `json:"data"`
}
type EncryptCertificate struct {
	Algorithm      string `json:"algorithm"`
	Nonce          string `json:"nonce"`
	AssociatedData string `json:"associated_data"`
	Ciphertext     string `json:"ciphertext"`
}
type CertificateData struct {
	EncryptCertificate *EncryptCertificate `json:"encrypt_certificate"`
	DecryptCertificate string              `json:"decrypt_certificate"`
	SerialNo           string              `json:"serial_no"`
	EffectiveTime      time.Time           `json:"effective_time "`
	ExpireTime         time.Time           `json:"expire_time "`
}
