package wlc

import (
	"context"
	"net/http"
	"net/url"
)

const (
	kQueryURL     = "http://api2.wlc.nppa.gov.cn/idcard/authentication/query"
	kQueryTestURL = "https://wlc.nppa.gov.cn/test/authentication/query/"
)

func (c *client) Query(ctx context.Context, ai string) (*QueryResult, error) {
	return c.query(ctx, kQueryURL, ai)
}

func (c *client) QueryTest(ctx context.Context, code, ai string) (*QueryResult, error) {
	return c.query(ctx, kQueryTestURL+code, ai)
}

func (c *client) query(ctx context.Context, api, ai string) (*QueryResult, error) {
	var aux = struct {
		*Error
		Data struct {
			Result *QueryResult `json:"result"`
		} `json:"data"`
	}{}

	var values = url.Values{}
	values.Set("ai", ai)

	if err := c.request(ctx, http.MethodGet, api, values, nil, &aux); err != nil {
		return nil, err
	}

	if aux.Error != nil && aux.Error.ErrCode != 0 {
		return nil, aux.Error
	}

	return aux.Data.Result, nil
}
