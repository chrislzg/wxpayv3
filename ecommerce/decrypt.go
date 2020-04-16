package ecommerce

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/golang/glog"

	"ptapp.cn/util/wechat.v3/core"
)

func (c *payClient) Decrypt(algorithm string, cipherText string, associatedData string, nonce string) ([]byte, error) {
	// 默认使用AEAD_AES_256_GCM
	switch algorithm {
	default:
		fallthrough
	case core.AlgorithmAEADAES256GCM:
		decodedCipherText, _ := base64.StdEncoding.DecodeString(cipherText)

		block, err := aes.NewCipher([]byte(c.apiV3Key))
		if err != nil {
			glog.Errorf("invalid key:%v, %v", c.apiV3Key, err)
			return nil, err
		}

		aesGcm, err := cipher.NewGCM(block)
		if err != nil {
			return nil, err
		}

		plaintext, err := aesGcm.Open(nil, []byte(nonce), decodedCipherText, []byte(associatedData))
		if err != nil {
			return nil, err
		}
		return plaintext, nil
	}
}
