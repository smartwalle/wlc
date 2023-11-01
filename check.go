package wlc

import (
	"net/http"
)

const (
	kCheckURL     = "https://api.wlc.nppa.gov.cn/idcard/authentication/check"
	kCheckTestURL = "https://wlc.nppa.gov.cn/test/authentication/check/"
)

func (c *client) Check(param CheckParam) (*CheckResult, error) {
	return c.check(kCheckURL, param)
}

func (c *client) CheckTest(code string, param CheckParam) (*CheckResult, error) {
	return c.check(kCheckTestURL+code, param)
}

func (c *client) check(api string, param CheckParam) (*CheckResult, error) {
	var aux = struct {
		*Error
		Data struct {
			Result *CheckResult `json:"result"`
		} `json:"data"`
	}{}

	if err := c.request(http.MethodPost, api, nil, param, &aux); err != nil {
		return nil, err
	}

	if aux.Error != nil && aux.Error.ErrCode != 0 {
		return nil, aux.Error
	}

	return aux.Data.Result, nil
}
