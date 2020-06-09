/*
 *
 * account.go
 * NIMSDK
 *
 * Created by lintao on 2020/6/9 1:50 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package NIMSDK

import (
	"github.com/NSObjects/netease-im/path"
	"encoding/json"
	"errors"
	"net/url"
)

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
		s, ok := stateCode[resp.Code]
		if ok {
			return "", errors.New(s)
		}
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
		s, ok := stateCode[resp.Code]
		if ok {
			return errors.New(s)
		}
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
		s, ok := stateCode[resp.Code]
		if ok {
			return "", errors.New(s)
		}
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
		s, ok := stateCode[resp.Code]
		if ok {
			return "", errors.New(s)
		}
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
		s, ok := stateCode[resp.Code]
		if ok {
			return errors.New(s)
		}
	}

	return nil
}
