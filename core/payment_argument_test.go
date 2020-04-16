package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTurnUpPaymentArgumentApp_SetPaySign(t *testing.T) {
	ast := assert.New(t)

	apiPrivateKeyStr := `
-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCiOmDEO+fRPKYb
AOdkbtzqNkdRi2fIwOYDeu/Sa4Q5TasCfyxFamf3ZgeOZ+PrGWN9/HhRMLH5VJn5
qlEoOosqOeuMbPAtl9V2Azyg89PlCnE5NdUs6K89lllp80bcA65RFRi9q6SrewrE
tL0q6WMIX+1/6ItTCscJTuvaU5wwlf53TWVj6TWYCTKUSEZpZrf2nZABzRJPG/pE
3zbanjEjEdY5dlQOCyKMVCOW6PSbNG+/M6Pc23XqyOrnmZBs/shESVMaI3LBsE04
HAZvUtvumAw2Dd4dtfTou9Na1dv7wgJslhimcQsMgas6pqiw3F5+Ln55KSbpNKN5
vcmVX8zPAgMBAAECggEAJqFsJY52V1bv3wJKF5pmtMcUGJDvt9hnbXC65wp/Q/4A
uOl7q1P5Sepb5kdK+NLk5D1YkUzdNJg2ykMKbF/0f6unMZtHXjQHusBukY0+ag4g
hBUTnEwUXsAMrrQlrYKa4ef6KkBTSBpWqBy55xwIXxgowAqqEq5AUzEd7LF4QgE0
nIcO7gxy+GqfOyygNI9INyNdjfAecZrOzdTW3FAhroTNcbVtdE0TgqTXxJJkbLXE
HTt4Z+JfbI70B1UYFpN4j44o8yKWFXjPtrbLY3As8/b7cq8n+wfUDfmKXK6AKT1O
xeExAwyKctR1GAZqyn/tdEH3PBLAu88yCQcYhzKDoQKBgQDVYhbAATHsALOLtprV
uW5IMt9fnXAB50imeqXmn86lA0/ntmloKozsmmLLudAOzidT6e3ODGFhv/KUyIk4
5q8SgWVo30GQxrsXPfOrJlOElEff3ALVKOMBBjmAZxec14fHoohOa3jMhgsyUKQf
w5hgcmOu72eHuHdlW/2z0cjWNwKBgQDCoNEEHdzFKg7kFruXLzU3hmYncmkhkZHF
OqmSFs3fzms4lXD/zreHejAp4GiBjgYsOH2dVGreBN9Q7tnCUdvYwB9jK+WR7Zch
lBV9iGOxtBOemCN3SUZR+yeKAAgui4Py40A733nPlEr4Zc1XAs8AC1epzvLJonPa
mzBtMTdyKQKBgQC3jWwSeDSwVZ4dBdRFGwCBvLknb6+lA4YcJw7Exx0kFyhKI0Ci
6U9WTCvGIa2WvsFXzrfQchfm1Q3f7G0V9GIPIh3Qy5OD7V+My67qv8pCFqeJKqGJ
KWW0QN1/1a6bLU/Qa8Ci7JH6JShGfNXhuQg/lsam+atuNUEHgM1JPKFtmwKBgQCn
fBaCUWRjcw7/bySdNG26S3jrJ0SbM5bav+GequsdVpfkSI3GRNCg0CBUWR31pw9e
zHokgrm4Nz8peXGBDEqBGsun3uWej3PH3JQlw9Hu4UUk7E1Q4IiYEeZzlhV0YHD6
+l6TZ3t+i2F8orZy0yLpKdmVclZx98905qlkvb62CQKBgDaaMHvYJsqxMKAfA8s+
EIvb6Bl00dkbxt2vvd4H/D15Z1szbNfXScSzDRt4eC2OhZAlax1yKoYeoCLGzIQq
bImljqKs6OH98YQc2GdGPjg3ir8iLg79ME3GHbSesGxT9qGwqek9GnanaqZD25PL
7WG+VNo1LqeswxVBgNT969IV
-----END PRIVATE KEY-----
`

	apiPrivateKey, err := ParsePrivateKey(apiPrivateKeyStr)
	ast.NoError(err)

	testCases := []struct {
		Input  TurnUpPaymentArgument
		Except string
	}{
		{
			Input: &TurnUpPaymentArgumentApp{
				Appid:        "wx2513074d5de96702",
				Partnerid:    "1579156067",
				Prepayid:     "001",
				Noncestr:     "9a45c5190bf23f0824ade96c69ab2504",
				Timestamp:    "1579156067",
				PackageValue: "Sign=WXPay",
			},
			Except: "EXiyssl5N+q36nxqsKxC2xTLT07O1XWbfOz0jyX1usIk26CVu/L9x1tozNoGF7hRXrb35Nvg2OoLSAYIWRQqSeiBrsW/JcvAeshVKrRXT/jA3RzuH1nZB31qDgjmUPnT4j3JmHhHFyMSRdRBvx6VwcF/aSWpAX3ieobGjcm98hrBpTHj5bElPwWxL6bRDN5bMA05OktKq8xs7njKNlkWOBqtVDPxrCsu3UMvB8A8fdiUduhbQkyfSf1f4TdoElL2MZ2/G4R/10Mt10J2KDaRhuUutIsN8xr/jwP9ty7VnNGLd4s+NAZZvfREf+luTuk7TT/WvKVyKlrS6Klnkr0jKg==",
		},
		{
			Input: &TurnUpPaymentArgumentJsApi{
				AppId:     "wxf940127ebc224e39",
				TimeStamp: "1579156067",
				Package:   "prepay_id=001",
				SignType:  "RSA",
				NonceStr:  "a97f72200bef478d77337eb9be0f0361",
			},
			Except: "E43snEjZEa9Oj4zImDbJ7pxU95l/U0yfjngCB91Vowk6r8XLRxyqEVha80tuycbZk+ySRTzJhe7ebvcb3P/FlTW9WZDaMUAQov500rw/uOvOoysTvs7WYAgoPG6kjQ8BDuTyfD5sNcaJav1z9J5mDYYZN26FIfwVziIhMed9dlJxUXadPRvo2jyH4uRaeS1Rx2O1U2c/Chq4nXjHWSS5hW1WfqCP98CSI5jyQZItPIJvvLnNAMriK/tL67HWc6VV2TWm743v3YJIqy+CFbSxaquA+9BJqJ4x81efSybmxJpVQTpszMGNW+A3SFnNBWTT2NICBPl0SvM7FaDHiHfsBA==",
		},
		{
			Input: &TurnUpPaymentArgumentXcx{
				AppId:     "wx99d20a019ced69ee",
				TimeStamp: "1579156067",
				Package:   "prepay_id=001",
				SignType:  "RSA",
				NonceStr:  "bbd9f6b2bcad10456a24fb2bd158e042",
			},
			Except: "DVfKWeM+sc4NQjRRkGPil5VjrGGIvHhRbHLcJTv0v4v2UDKaFfJTD68V+ueWn2d60uBtMO+eA5uXm0pMRJBmSkQ+oJ2+7efnsLyadS+xr3vJrZCjTAqLU38uIhh2jTQQK4ir8W9TXs7WPgpXd5VCzhLgsXzl5yrhvXHhVAvPYb3IHY9J3a3CCYrtUqXFmMLgaIbCKsFaWrd76GnuRrqjaR0iBK/bTRfj95xjIKSSCtLZpRd7E2k+Dg3nFI8GgL/dEuxz+9jGnW690mus8bhWuFMELtwbykAgkhn91bsJBpB0SsSOBI1r8SLZOC2ysVnWCNG7253PbK4mpibn4GukwQ==",
		},
	}

	for _, t := range testCases {
		err := t.Input.SetPaySign(apiPrivateKey)
		ast.NoError(err)
		ast.Equal(t.Except, t.Input.GetPaySign())
	}
}
