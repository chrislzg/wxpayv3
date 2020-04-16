package core

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrInvalidWechatPaySerial    = errors.New("invalid Wechatpay-Serial")
	ErrInvalidWechatPayTimestamp = errors.New("invalid Wechatpay-Timestamp")
	ErrInvalidWechatPayNonce     = errors.New("invalid Wechatpay-Nonce")
	ErrInvalidWechatPaySignature = errors.New("invalid Wechatpay-Signature")
	ErrPemDecodeFailed           = errors.New("pem decode failed")
	ErrInvalidResponse           = errors.New("verify response failed")
	ErrEventTypeNotSuccess       = errors.New("notify event_type is not TRANSACTION.SUCCESS")
	ErrEmptyNotifyResource       = errors.New("empty notify resource")
	ErrDecryptFailed             = errors.New("decrypt failed")
	ErrIncorrectResourceType     = errors.New("incorrect resource-type")
)

type ErrResponseBody struct {
	HttpStatus int             `json:"http_status"`
	Code       string          `json:"code"`
	Message    string          `json:"message"`
	ReqId      string          `json:"req_id"`
	Detail     json.RawMessage `json:"detail"`
}

//type ErrResponseDetail struct {
//	Field    string `json:"field"`
//	Value    string `json:"value"`
//	Issue    string `json:"issue"`
//	Location string `json:"location"`
//}

func (r *ErrResponseBody) Error() string {
	if r.Detail == nil {
		return fmt.Sprintf("HttpStatus:%v Code:%s Message:%s RequestId:%s", r.HttpStatus, r.Code, r.Message, r.ReqId)
	}
	return fmt.Sprintf("HttpStatus:%v Code:%s Message:%s RequestId:%s Detail:%s", r.HttpStatus, r.Code, r.Message, r.ReqId, r.Detail)
}
