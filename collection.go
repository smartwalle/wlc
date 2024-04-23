package wlc

import (
	"context"
	"net/http"
)

const (
	kLoginTraceURL     = "http://api2.wlc.nppa.gov.cn/behavior/collection/loginout"
	kLoginTraceTestURL = "https://wlc.nppa.gov.cn/test/collection/loginout/"
)

func (c *client) LoginTrace(ctx context.Context, param LoginTraceParam) ([]*LoginTraceResult, error) {
	return c.loginTrace(ctx, kLoginTraceURL, param)
}

func (c *client) LoginTraceTest(ctx context.Context, code string, param LoginTraceParam) ([]*LoginTraceResult, error) {
	return c.loginTrace(ctx, kLoginTraceTestURL+code, param)
}

func (c *client) loginTrace(ctx context.Context, api string, param LoginTraceParam) ([]*LoginTraceResult, error) {
	var aux = struct {
		*Error
		Data struct {
			Results []*LoginTraceResult `json:"results"`
		} `json:"data"`
	}{}

	if err := c.request(ctx, http.MethodPost, api, nil, param, &aux); err != nil {
		return nil, err
	}

	if aux.Error != nil && aux.Error.ErrCode != 0 {
		return aux.Data.Results, aux.Error
	}

	return aux.Data.Results, nil
}
