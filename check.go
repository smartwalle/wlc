package wlc

import (
	"context"
	"net/http"
)

const (
	kCheckURL     = "https://api.wlc.nppa.gov.cn/idcard/authentication/check"
	kCheckTestURL = "https://wlc.nppa.gov.cn/test/authentication/check/"
)

func (c *client) Check(ctx context.Context, param CheckParam) (*CheckResult, error) {
	return c.check(ctx, kCheckURL, param)
}

func (c *client) CheckTest(ctx context.Context, code string, param CheckParam) (*CheckResult, error) {
	return c.check(ctx, kCheckTestURL+code, param)
}

func (c *client) check(ctx context.Context, api string, param CheckParam) (*CheckResult, error) {
	var aux = struct {
		*Error
		Data struct {
			Result *CheckResult `json:"result"`
		} `json:"data"`
	}{}

	if err := c.request(ctx, http.MethodPost, api, nil, param, &aux); err != nil {
		return nil, err
	}

	if aux.Error != nil && aux.Error.ErrCode != 0 {
		return nil, aux.Error
	}

	return aux.Data.Result, nil
}
