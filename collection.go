package wlc

import (
	"encoding/json"
	"net/http"
)

const (
	kLoginTraceURL     = "http://api2.wlc.nppa.gov.cn/behavior/collection/loginout"
	kLoginTraceTestURL = "https://wlc.nppa.gov.cn/test/collection/loginout/"
)

func (this *client) LoginTrace(param LoginTraceParam) (result *LoginTraceRsp, err error) {
	data, err := this.request(http.MethodPost, kLoginTraceURL, nil, param)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (this *client) LoginTraceTest(code string, param LoginTraceParam) (result *LoginTraceRsp, err error) {
	data, err := this.request(http.MethodPost, kLoginTraceTestURL+code, nil, param)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
