/*
网易云信IM服务端SDK

Example:

package main

import (
	"fmt"
	"github.com/NSObjects/netease"
)

func main() {
	n := netease.NewNetEaseIM("appkey", "secret")
	token, err := n.CreateAccount(netease.Account{
		Accid: "xx",
		Name:  "xx",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(token)
}


Learn more at https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%8E%A5%E5%8F%A3%E6%A6%82%E8%BF%B0
*/

package netease

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/NSObjects/netease/encrypt"
	"github.com/NSObjects/netease/path"
)

type NetEaseIM struct {
	appKey   string
	secret   string
	basePath string
}

func NewNetEaseIM(appKey, secret string) NetEaseIM {
	return NetEaseIM{appKey: appKey, secret: secret, basePath: "https://api.netease.im/nimserver/"}
}

func (n NetEaseIM) buildHeader() http.Header {
	h := http.Header{}
	nonce := encrypt.RandomString(20)
	t := fmt.Sprintf("%d", time.Now().Unix())
	h.Add("AppKey", n.appKey)
	h.Add("Nonce", nonce)
	h.Add("CurTime", t)
	h.Add("CheckSum", encrypt.GenerateCheckSum(nonce, n.secret, t))
	h.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	return h
}

func (n NetEaseIM) request(path path.Path, params url.Values) ([]byte, error) {
	body := bytes.NewBufferString(params.Encode())

	client := &http.Client{}
	req, err := http.NewRequest("POST", n.basePath+string(path), body)
	req.Header = n.buildHeader()

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Printf(string(respBody))
	return respBody, nil
}

func structToMap(i interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(i)
	typ := iVal.Type()
	for index := 0; index < iVal.NumField(); index++ {
		f := iVal.Field(index)

		var v string

		switch f.Type().Kind() {
		case reflect.Int8, reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			v = strconv.FormatInt(f.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case reflect.Float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case reflect.Float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case reflect.String:
			v = f.String()
		case reflect.Slice:
			switch f.Interface().(type) {
			case []byte:
				v = string(f.Bytes())
			case []string:
				b, err := json.Marshal(f.Interface())
				if err != nil {
					panic(err)
				}
				v = string(b)
			}

		}
		tag := typ.Field(index).Tag.Get("json")
		values.Set(tag, v)
	}
	return
}

var stateCode = map[int]string{
	200: "操作成功", 201: "客户端版本不对，需升级sdk", 301: "被封禁",
	302: "用户名或密码错误", 315: "IP限制", 403: "非法操作或没有权限",
	404: "对象不存在", 405: "参数长度过长", 406: "对象只读",
	408: "客户端请求超时", 413: "验证失败(短信服务)", 414: "参数错误",
	415: "客户端网络问题", 416: "频率控制", 417: "重复操作",
	418: "通道不可用(短信服务)", 419: "数量超过上限", 422: "账号被禁用", 423: "帐号被禁言",
	431: "HTTP重复请求", 500: "服务器内部错误", 503: "服务器繁忙", 508: "消息撤回时间超限",
	509: "无效协议", 514: "服务不可用", 998: "解包错误", 999: "打包错误", 801: "群人数达到上限", 802: "没有权限",
	803: "群不存在", 804: "用户不在群", 805: "群类型不匹配", 806: "创建群数量达到限制", 807: "群成员状态错误",
	808: "申请成功", 809: "已经在群内", 810: "邀请成功", 811: "@账号数量超过限制", 812: "群禁言，普通成员不能发送消息",
	813: "群拉人部分成功", 814: "禁止使用群组已读服务", 815: "群管理员人数超过上限", 9102: "通道失效",
	9103: "已经在他端对这个呼叫响应过了", 11001: "通话不可达，对方离线状态", 13001: "IM主连接状态异常", 13002: "聊天室状态异常",
	13004: "账号在黑名单中,不允许进入聊天室", 13005: "在禁言列表中,不允许发言", 10431: "输入email不是邮箱",
	10432: "输入mobile不是手机号码", 10433: "注册输入的两次密码不相同", 10434: "企业不存在", 10435: "登陆密码或帐号不对",
	10436: "app不存在", 10437: "email已注册", 10438: "手机号已注册", 10441: "app名字已经存在",
	10404: "房间不存在", 10405: "房间已存在", 10406: "不在房间内", 10407: "已经在房间内",
	10408: "邀请不存在或已过期", 10409: "邀请已经拒绝", 10410: "邀请已经接受了", 10201: "对方云信不在线",
	10202: "对方云信不在线，且推送也不可达", 10419: "房间人数超限", 10420: "已经在房间内（自己的其他端）", 10417: "uid冲突",
}
