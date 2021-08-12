/*
 *
 * account.go
 * NIMSDK
 *
 * Created by lintao on 2020/6/9 1:50 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package netease

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/5xxxx/netease/path"
)

//accid	String	是	网易云通信ID，最大长度32字符，必须保证一个
//APP内唯一（只允许字母、数字、半角下划线_、 @、半角点以及半角-组成，不区分大小写，
//会统一小写处理，请注意以此接口返回结果中的accid为准）。
//name	String	否	网易云通信ID昵称，最大长度64字符。
//props	String	否	json属性，开发者可选填，最大长度1024字符。该参数已不建议使用。
//icon	String	否	网易云通信ID头像URL，开发者可选填，最大长度1024
//token	String	否	网易云通信ID可以指定登录token值，最大长度128字符，
//并更新，如果未指定，会自动生成token，并在创建成功后返回
//sign	String	否	用户签名，最大长度256字符
//email	String	否	用户email，最大长度64字符
//birth	String	否	用户生日，最大长度16字符
//mobile	String	否	用户mobile，最大长度32字符，非中国大陆手机号码需要填写国家代码(如美国：+1-xxxxxxxxxx)或地区代码(如香港：+852-xxxxxxxx)
//gender	int	否	用户性别，0表示未知，1表示男，2女表示女，其它会报参数错误
//ex	String	否	用户名片扩展字段，最大长度1024字符，用户可自行扩展，建议封装成JSON字符串
type Account struct {
	Accid  string `json:"accid" `
	Name   string `json:"name" `
	Props  string `json:"props" `
	Icon   string `json:"icon" `
	Token  string `json:"token" `
	Sign   string `json:"sign" `
	Email  string `json:"email" `
	Birth  string `json:"birth" `
	Mobile string `json:"mobile" `
	Gender int    `json:"gender" `
	Ex     string `json:"ex" `
}

type AccountResponse struct {
	Code int `json:"code"`
	Info struct {
		Token string `json:"token"`
		Accid string `json:"accid"`
		Name  string `json:"name"`
	} `json:"info"`
}

//第三方帐号导入到网易云通信平台。注册成功后务必在自身的应用服务器上维护好accid与token。
//注意accid，name长度以及考虑管理token。
//云信应用内的accid若涉及字母，请一律为小写，并确保服务端与所有客户端均保持小写
func (n NetEaseIM) CreateAccount(account Account) (string, error) {
	b, err := n.request(path.CreateAccount, structToMap(account))
	if err != nil {
		return "", err
	}
	var resp AccountResponse
	if err = json.Unmarshal(b, &resp); err != nil {
		return "", err
	}

	if resp.Code != 200 {
		//s, ok := stateCode[resp.Code]
		//if ok {
		//	return "", errors.New(s)
		//}
		return "", errors.New(string(b))
	}

	return resp.Info.Token, nil
}

//更新网易云通信token。通过该接口，可以对accid更新到指定的token，更新后请开发者务必做好本地的维护。更新后，需要确保客户端SDK再次登录时携带的token保持最新。
//accid	String	是	网易云通信ID，最大长度32字符，必须保证一个 APP内唯一
//props	String	否	该参数已不建议使用。
//token	String	否	网易云通信ID可以指定登录token值，最大长度128字符
func (n NetEaseIM) UpdateAccount(accid, token string) error {
	params := url.Values{}
	params.Set("accid", accid)
	params.Set("token", token)
	b, err := n.request(path.UpdateAccount, params)
	if err != nil {
		return err
	}

	var resp struct {
		Code int `json:"code"`
	}

	if err = json.Unmarshal(b, &resp); err != nil {
		return err
	}

	if resp.Code != 200 {
		//s, ok := stateCode[resp.Code]
		//if ok {
		//	return errors.New(s)
		//}

		return errors.New(string(b))
	}

	return nil
}

//由云信webserver随机重置网易云通信ID的token，同时将新的token返回，更新后请开发者务必做好本地的维护。
//此接口与网易云通信token更新接口最大的区别在于：前者的token是由云信服务器指定，后者的token是由开发者自己指定。
// accid	String	是	网易云通信ID，最大长度32字符，必须保证一个APP内唯一
func (n NetEaseIM) RefreshToken(accid string) (string, error) {
	params := url.Values{}
	params.Set("accid", accid)
	b, err := n.request(path.RefreshToken, params)
	if err != nil {
		return "", err
	}
	var resp AccountResponse

	if err = json.Unmarshal(b, &resp); err != nil {
		return "", err
	}

	if resp.Code != 200 {
		//s, ok := stateCode[resp.Code]
		//if ok {
		//	return "", errors.New(s)
		//}
		return "", errors.New(string(b))
	}

	return resp.Info.Token, nil
}

//1.封禁网易云通信ID后，此ID将不能再次登录。若封禁时，该id处于登录状态，则当前登录不受影响，仍然可以收发消息。封禁效果会在下次登录时生效。因此建议，
//将needkick设置为true，让该账号同时被踢出登录。
//2.出于安全目的，账号创建后只能封禁，不能删除；封禁后账号仍计入应用内账号总数。
//参数	类型	必须	说明
//accid	String	是	网易云通信ID，最大长度32字符，必须保证一个 APP内唯一
//needkick	String	否	是否踢掉被禁用户，true或false，默认false
//kickNotifyExt	String	否	踢人时的扩展字段，SDK版本需要大于等于v7.7.0//
func (n NetEaseIM) BlockAccount(accid string, needkick string, kickNotifyExt string) (string, error) {
	params := url.Values{}
	params.Set("accid", accid)
	params.Set("needkick", needkick)
	params.Set("kickNotifyExt", kickNotifyExt)
	b, err := n.request(path.BlockAccount, params)
	if err != nil {
		return "", err
	}
	var resp AccountResponse

	if err = json.Unmarshal(b, &resp); err != nil {
		return "", err
	}

	if resp.Code != 200 {
		//s, ok := stateCode[resp.Code]
		//if ok {
		//	return "", errors.New(s)
		//}
		return "", errors.New(string(b))
	}

	return resp.Info.Token, nil
}

// 解禁被封禁的网易云通信ID
// accid	String	是	网易云通信ID，最大长度32字符，必须保证一个 APP内唯一
func (n NetEaseIM) UnBlockAccount(accid string) error {
	params := url.Values{}
	params.Set("accid", accid)

	b, err := n.request(path.UnblockAccount, params)
	if err != nil {
		return err
	}
	var resp AccountResponse

	if err = json.Unmarshal(b, &resp); err != nil {
		return err
	}

	if resp.Code != 200 {
		//s, ok := stateCode[resp.Code]
		//if ok {
		//	return errors.New(s)
		//}
		return errors.New(string(b))
	}

	return nil
}
