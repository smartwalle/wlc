package main

import (
	"fmt"
	"github.com/smartwalle/wlc"
	"time"
)

func main() {
	var client = wlc.NewTest("a0298befcc6340dcac218d5016669df3", "bbd3322487be9f3b2af1e7d1d3734f41", "1101999999")

	// 注意：运行前需要设置每个函数的第二个参数，该参数可以从 网络游戏防沉迷实名认证系统->数据共享->接口测试 中获取

	// testcase01-实名认证接口 - 认证成功
	check(client, "", "100000000000000001", "某一一", "110000190101010001")

	// testcase02-实名认证接口 - 认证中
	check(client, "", "200000000000000002", "某二二", "110000190201020004")

	// testcase03-实名认证接口 - 认证失败
	check(client, "", "3200000000000000002", "某二二", "110000190201020004")

	// testcase04-实名认证结果查询接口 - 认证成功
	query(client, "", "100000000000000001")

	// testcase05-实名认证结果查询接口 - 认证中
	query(client, "", "200000000000000001")

	// testcase06-实名认证结果查询接口 - 认证失败
	query(client, "", "300000000000000001")

	// testcase07-游戏用户行为数据上报接口 - 游客
	loginTraceGuest(client, "")

	// testcase08-游戏用户行为数据上报接口 - 认证用户
	loginTraceUser(client, "")
}

func check(client wlc.TestClient, code string, ai, name, idNum string) {
	var p = wlc.CheckParam{}
	p.AI = ai
	p.Name = name
	p.IdNum = idNum
	var result, err = client.CheckTest(code, p)
	if err != nil {
		fmt.Println("实名认证发生错误:", err)
		return
	}

	if result != nil {
		fmt.Println("实名认证结果:", result.PI, result.Status)
	}
}

func query(client wlc.TestClient, code string, ai string) {
	var result, err = client.QueryTest(code, ai)
	if err != nil {
		fmt.Println("实名认证查询发生错误:", err)
		return
	}

	if result != nil {
		fmt.Println("实名认证查询结果:", result.PI, result.Status)
	}
}

func loginTraceGuest(client wlc.TestClient, code string) {
	var p = wlc.LoginTraceParam{}
	p.AddGuestLogin("12345678901234567890123456789012", time.Now().Unix(), "12345678901234567890123456789012")

	var result, err = client.LoginTraceTest(code, p)
	if err != nil {
		fmt.Println("上报数据发生错误:", err)
		return
	}

	for _, result := range result {
		fmt.Println("上报数据发生错误:", result.No, result.ErrCode, result.ErrMsg)
	}
}

func loginTraceUser(client wlc.TestClient, code string) {
	var p = wlc.LoginTraceParam{}
	p.AddUserLogin("12345678901234567890123456789012", time.Now().Unix(), "1fffbjzos82bs9cnyj1dna7d6d29zg4esnh99u")

	var result, err = client.LoginTraceTest(code, p)
	if err != nil {
		fmt.Println("上报数据发生错误:", err)
		return
	}

	for _, result := range result {
		fmt.Println("上报数据发生错误:", result.No, result.ErrCode, result.ErrMsg)
	}
}
