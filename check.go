package wlc

import (
	"encoding/json"
	"net/http"
)

const (
	kCheckURL     = "https://api.wlc.nppa.gov.cn/idcard/authentication/check"
	kCheckTestURL = "https://wlc.nppa.gov.cn/test/authentication/check/"
)

func (this *client) Check(param CheckParam) (result *CheckRsp, err error) {
	data, err := this.request(http.MethodPost, kCheckURL, nil, param)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (this *client) CheckTest(code string, param CheckParam) (result *CheckRsp, err error) {
	data, err := this.request(http.MethodPost, kCheckTestURL+code, nil, param)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
