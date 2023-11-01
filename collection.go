package wlc

import (
	"net/http"
)

const (
	kLoginTraceURL     = "http://api2.wlc.nppa.gov.cn/behavior/collection/loginout"
	kLoginTraceTestURL = "https://wlc.nppa.gov.cn/test/collection/loginout/"
)

func (c *client) LoginTrace(param LoginTraceParam) ([]*LoginTraceResult, error) {
	return c.loginTrace(kLoginTraceURL, param)
}

func (c *client) LoginTraceTest(code string, param LoginTraceParam) ([]*LoginTraceResult, error) {
	return c.loginTrace(kLoginTraceTestURL+code, param)
}

func (c *client) loginTrace(api string, param LoginTraceParam) ([]*LoginTraceResult, error) {
	var aux = struct {
		*Error
		Data struct {
			Results []*LoginTraceResult `json:"results"`
		} `json:"data"`
	}{}

	if err := c.request(http.MethodPost, api, nil, param, &aux); err != nil {
		return nil, err
	}

	if aux.Error != nil && aux.Error.ErrCode != 0 {
		return aux.Data.Results, aux.Error
	}

	return aux.Data.Results, nil
}
