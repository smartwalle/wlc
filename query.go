package wlc

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	kQueryURL     = "http://api2.wlc.nppa.gov.cn/idcard/authentication/query"
	kQueryTestURL = "https://wlc.nppa.gov.cn/test/authentication/query/"
)

func (this *client) Query(ai string) (result *QueryRsp, err error) {
	var values = url.Values{}
	values.Set("ai", ai)

	data, err := this.request(http.MethodGet, kQueryURL, values, nil)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (this *client) QueryTest(code, ai string) (result *QueryRsp, err error) {
	var values = url.Values{}
	values.Set("ai", ai)

	data, err := this.request(http.MethodGet, kQueryTestURL+code, values, nil)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
