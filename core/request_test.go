package core

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildUrl(t *testing.T) {
	ast := assert.New(t)

	type input struct {
		params map[string]string
		routes []string
		query  url.Values
	}
	testCases := []struct {
		Input  input
		Except string
	}{
		{
			Input: input{
				params: map[string]string{"withdraw_id": "112236"},
				routes: []string{ApiWithdrawFundStatus},
				query:  url.Values{"tq": {"d"}},
			},
			Except: "https://api.mch.weixin.qq.com/v3/ecommerce/fund/withdraw/112236?tq=d",
		},
		{
			Input: input{
				params: map[string]string{"withdraw_id": "112236"},
				routes: []string{ApiWithdrawFundStatus},
			},
			Except: "https://api.mch.weixin.qq.com/v3/ecommerce/fund/withdraw/112236",
		},
		{
			Input: input{
				routes: []string{ApiCertification},
			},
			Except: "https://api.mch.weixin.qq.com/v3/certificates",
		},
		{
			Input: input{
				routes: []string{ApiCertification},
				query:  url.Values{"tq": {"d"}},
			},
			Except: "https://api.mch.weixin.qq.com/v3/certificates?tq=d",
		},
		{
			Input: input{
				params: map[string]string{"withdraw_id": "112236"},
				routes: []string{ApiCertification},
				query:  url.Values{"tq": {"d"}},
			},
			Except: "https://api.mch.weixin.qq.com/v3/certificates?tq=d",
		},
	}
	for _, testCase := range testCases {
		buildUrl := BuildUrl(testCase.Input.params, testCase.Input.query, testCase.Input.routes...)
		ast.Equal(testCase.Except, buildUrl)
	}

}
