package core

import (
	"bufio"
	"bytes"
	"crypto/rsa"
)

type TurnUpPaymentArgument interface {
	SetPaySign(privateKey *rsa.PrivateKey) error
	GetPaySign() string
}

// APP调起支付的参数
type TurnUpPaymentArgumentApp struct {
	Appid        string `json:"appid"`
	Partnerid    string `json:"partnerid"`
	Prepayid     string `json:"prepayid"`
	Noncestr     string `json:"noncestr"`
	Timestamp    string `json:"timestamp"`
	PackageValue string `json:"_package"`
	Sign         string `json:"sign"`
}

// JSAPI调起支付的参数
type TurnUpPaymentArgumentJsApi struct {
	AppId     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
	NonceStr  string `json:"nonceStr"`
}

// 小程序调起支付的参数
type TurnUpPaymentArgumentXcx struct {
	AppId     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
	NonceStr  string `json:"nonceStr"`
}

func (a *TurnUpPaymentArgumentApp) SetPaySign(privateKey *rsa.PrivateKey) error {
	bf := bytes.NewBuffer([]byte{})
	bw := bufio.NewWriter(bf)
	if err := BufWriteStringWithLn(bw, a.Appid); err != nil {
		return err
	}
	if err := BufWriteStringWithLn(bw, a.Timestamp); err != nil {
		return err
	}
	if err := BufWriteStringWithLn(bw, a.Noncestr); err != nil {
		return err
	}
	if err := BufWriteStringWithLn(bw, a.Prepayid); err != nil {
		return err
	}
	if err := bw.Flush(); err != nil {
		return err
	}
	err := bw.Flush()
	if err != nil {
		return err
	}
	paySign, err := Sign(bf.Bytes(), privateKey)
	if err != nil {
		return err
	}
	a.Sign = paySign
	return nil
}

func (a *TurnUpPaymentArgumentApp) GetPaySign() string {
	return a.Sign
}

func (a *TurnUpPaymentArgumentJsApi) SetPaySign(privateKey *rsa.PrivateKey) error {
	bf := bytes.NewBuffer([]byte{})
	bw := bufio.NewWriter(bf)
	if err := BufWriteStringWithLn(bw, a.AppId); err != nil {
		return err
	}
	if err := BufWriteStringWithLn(bw, a.TimeStamp); err != nil {
		return err
	}
	if err := BufWriteStringWithLn(bw, a.NonceStr); err != nil {
		return err
	}
	if err := BufWriteStringWithLn(bw, a.Package); err != nil {
		return err
	}
	if err := bw.Flush(); err != nil {
		return err
	}
	paySign, err := Sign(bf.Bytes(), privateKey)
	if err != nil {
		return err
	}
	a.PaySign = paySign
	return nil
}

func (a *TurnUpPaymentArgumentJsApi) GetPaySign() string {
	return a.PaySign
}

func (a *TurnUpPaymentArgumentXcx) SetPaySign(privateKey *rsa.PrivateKey) error {
	bf := bytes.NewBuffer([]byte{})
	bw := bufio.NewWriter(bf)
	if err := BufWriteStringWithLn(bw, a.AppId); err != nil {
		return err
	}
	if err := BufWriteStringWithLn(bw, a.TimeStamp); err != nil {
		return err
	}
	if err := BufWriteStringWithLn(bw, a.NonceStr); err != nil {
		return err
	}
	if err := BufWriteStringWithLn(bw, a.Package); err != nil {
		return err
	}
	if err := bw.Flush(); err != nil {
		return err
	}
	//h := md5.New()
	//h.Write(bf.Bytes())
	//paySign := h.Sum(nil)
	//a.PaySign = strings.ToUpper(string(paySign))
	paySign, err := Sign(bf.Bytes(), privateKey)
	if err != nil {
		return err
	}
	a.PaySign = paySign
	return nil
}

func (a *TurnUpPaymentArgumentXcx) GetPaySign() string {
	return a.PaySign
}
