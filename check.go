package wlc

import (
	"net/http"
)

const (
	kCheckURL     = "https://api.wlc.nppa.gov.cn/idcard/authentication/check"
	kCheckTestURL = "https://wlc.nppa.gov.cn/test/authentication/check/"
)

func (this *client) Check(param CheckParam) (*CheckResult, error) {
	return this.check(kCheckURL, param)
}

func (this *client) CheckTest(code string, param CheckParam) (*CheckResult, error) {
	return this.check(kCheckTestURL+code, param)
}

func (this *client) check(api string, param CheckParam) (*CheckResult, error) {
	var aux = struct {
		*Error
		Data struct {
			Result *CheckResult `json:"result"`
		} `json:"data"`
	}{}

	if err := this.request(http.MethodPost, api, nil, param, &aux); err != nil {
		return nil, err
	}

	if aux.Error != nil && aux.Error.ErrCode != 0 {
		return nil, aux.Error
	}

	return aux.Data.Result, nil
}
