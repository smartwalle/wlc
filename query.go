package wlc

import (
	"net/http"
	"net/url"
)

const (
	kQueryURL     = "http://api2.wlc.nppa.gov.cn/idcard/authentication/query"
	kQueryTestURL = "https://wlc.nppa.gov.cn/test/authentication/query/"
)

func (this *client) Query(ai string) (*QueryResult, error) {
	return this.query(kQueryURL, ai)
}

func (this *client) QueryTest(code, ai string) (*QueryResult, error) {
	return this.query(kQueryTestURL+code, ai)
}

func (this *client) query(api, ai string) (*QueryResult, error) {
	var aux = struct {
		*Error
		Data struct {
			Result *QueryResult `json:"result"`
		} `json:"data"`
	}{}

	var values = url.Values{}
	values.Set("ai", ai)

	if err := this.request(http.MethodGet, api, values, nil, &aux); err != nil {
		return nil, err
	}

	if aux.Error != nil && aux.Error.ErrCode != 0 {
		return nil, aux.Error
	}

	return aux.Data.Result, nil
}
