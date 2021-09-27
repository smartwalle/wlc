package wlc

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type client struct {
	appId        string
	secretKey    string
	secretKeyHex []byte
	bizId        string
	client       *http.Client
}

// Client 生产客户端
type Client interface {
	// Check 实名认证接口,
	// 网络游戏用户实名认证服务接口，面向已经接入网络游戏防
	// 沉迷实名认证系统的游戏运营单位提供服务，游戏运营单位调用
	// 该接口进行用户实名认证工作，本版本仅支持大陆地区的姓名和
	// 二代身份证号核实认证。
	Check(param CheckParam) (result *CheckRsp, err error)

	// Query 实名认证结果查询接口,
	// 网络游戏用户实名认证结果查询服务接口，面向已经提交用
	// 户实名认证且没有返回结果的游戏运营单位提供服务，游戏运营
	// 单位可以调用该接口，查询已经提交但未返回结果用户的实名认
	// 证结果。
	Query(ai string) (result *QueryRsp, err error)

	// LoginTrace 游戏用户行为数据上报接口
	// 游戏用户行为数据上报接口，面向已经接入网络游戏防沉迷
	// 实名认证系统的游戏运营单位提供服务，游戏运营单位调用该接
	// 口上报游戏用户上下线行为数据。
	LoginTrace(param LoginTraceParam) (result *LoginTraceRsp, err error)
}

// TestClient 接口测试辅助客户端
type TestClient interface {
	CheckTest(code string, param CheckParam) (result *CheckRsp, err error)

	QueryTest(code, ai string) (result *QueryRsp, err error)

	LoginTraceTest(code string, param LoginTraceParam) (result *LoginTraceRsp, err error)
}

func New(appId, secretKey, bizId string) Client {
	var c = &client{}
	c.appId = appId
	c.secretKey = secretKey
	c.secretKeyHex, _ = hex.DecodeString(secretKey)
	c.bizId = bizId
	c.client = http.DefaultClient
	return c
}

func NewTest(appId, secretKey, bizId string) TestClient {
	var c = &client{}
	c.appId = appId
	c.secretKey = secretKey
	c.secretKeyHex, _ = hex.DecodeString(secretKey)
	c.bizId = bizId
	c.client = http.DefaultClient
	return c
}

func (this *client) request(method, api string, values url.Values, param interface{}) ([]byte, error) {
	if values == nil {
		values = url.Values{}
	}

	var body string
	var bodyReader io.Reader
	if param != nil {
		data, err := json.Marshal(param)
		if err != nil {
			return nil, err
		}

		// 加密请求参数
		data, err = this.encrypt(this.secretKeyHex, data)
		if err != nil {
			return nil, err
		}

		// 构造新的请求参数
		var p = &Param{}
		p.Data = base64.StdEncoding.EncodeToString(data)
		data, err = json.Marshal(p)
		if err != nil {
			return nil, err
		}

		body = string(data)
		bodyReader = bytes.NewReader(data)
	}

	var nURL = api
	if len(values) > 0 {
		nURL = api + "?" + values.Encode()
	}

	req, err := http.NewRequest(method, nURL, bodyReader)
	if err != nil {
		return nil, err
	}

	var now = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)

	values.Add("appId", this.appId)
	values.Add("bizId", this.bizId)
	values.Add("timestamps", now)

	var sign = this.sign(this.secretKey, values, body)

	req.Header.Set("appId", this.appId)
	req.Header.Set("bizId", this.bizId)
	req.Header.Set("timestamps", now)
	req.Header.Set("sign", sign)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	rsp, err := this.client.Do(req)
	if rsp != nil && rsp.Body != nil {
		defer rsp.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (this *client) encrypt(secretKeyHex []byte, data []byte) ([]byte, error) {
	var block, err = aes.NewCipher(secretKeyHex)
	if err != nil {
		return nil, err
	}

	mode, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	var nonce = make([]byte, mode.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return mode.Seal(nonce, nonce, data, nil), nil
}

func (this *client) sign(secretKey string, values url.Values, body string) string {
	var pList = make([]string, 0, 3+len(values))

	for key := range values {
		pList = append(pList, key+values.Get(key))
	}
	sort.Strings(pList)
	var data = secretKey + strings.Join(pList, "") + body

	var h = sha256.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
