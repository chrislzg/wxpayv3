package core

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNonceStr(t *testing.T) {
	ast := assert.New(t)

	nonce := NonceStr()
	t.Logf("nonce:%v", nonce)
	ast.Equal(hex.EncodedLen(16), len(nonce))
}

func TestBuildMessage(t *testing.T) {
	ast := assert.New(t)

	type input struct {
		httpMethod string
		url        string
		body       []byte
		nonce      string
		timestamp  int64
	}
	timestamp := time.Now().Unix()
	nonce := NonceStr()
	testCases := []struct {
		Input  input
		Except []byte
	}{
		{
			Input: input{
				httpMethod: http.MethodGet,
				url:        BuildUrl(nil, nil, ApiCertification),
				body:       nil,
				nonce:      nonce,
				timestamp:  timestamp,
			},
			Except: []byte(fmt.Sprintf(`GET
/v3/certificates
%v
%v

`, timestamp, nonce)),
		},
		{
			Input: input{
				httpMethod: http.MethodGet,
				url:        BuildUrl(nil, nil, ApiCertification),
				body:       []byte(`{"req"":"body test"}`),
				nonce:      nonce,
				timestamp:  timestamp,
			},
			Except: []byte(fmt.Sprintf(`GET
/v3/certificates
%v
%v

`, timestamp, nonce)),
		},
		{
			Input: input{
				httpMethod: http.MethodPost,
				url:        BuildUrl(nil, nil, ApiCertification),
				body:       []byte(`{"req"":"body test"}`),
				nonce:      nonce,
				timestamp:  timestamp,
			},
			Except: []byte(fmt.Sprintf(`POST
/v3/certificates
%v
%v
{"req"":"body test"}
`, timestamp, nonce)),
		},
		{
			Input: input{
				httpMethod: http.MethodPost,
				url:        BuildUrl(nil, nil, ApiCertification),
				body:       nil,
				nonce:      nonce,
				timestamp:  timestamp,
			},
			Except: []byte(fmt.Sprintf(`POST
/v3/certificates
%v
%v

`, timestamp, nonce)),
		},
		{
			Input: input{
				httpMethod: http.MethodPut,
				url:        BuildUrl(nil, nil, ApiCertification),
				body:       nil,
				nonce:      nonce,
				timestamp:  timestamp,
			},
			Except: []byte(fmt.Sprintf(`PUT
/v3/certificates
%v
%v

`, timestamp, nonce)),
		},
		{
			Input: input{
				httpMethod: http.MethodPut,
				url:        BuildUrl(nil, nil, ApiCertification),
				body:       []byte(`{"req"":"body test"}`),
				nonce:      nonce,
				timestamp:  timestamp,
			},
			Except: []byte(fmt.Sprintf(`PUT
/v3/certificates
%v
%v
{"req"":"body test"}
`, timestamp, nonce)),
		},
	}

	for _, tCase := range testCases {
		input := tCase.Input
		resp, err := BuildMessage(input.httpMethod, input.url, input.body, input.nonce, input.timestamp)
		ast.NoError(err)
		t.Logf("resp:%s", resp)
		t.Logf("except:%s", tCase.Except)
		ast.Equal(tCase.Except, resp)
	}
}
