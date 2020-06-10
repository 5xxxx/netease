/*
 *
 * friend.go
 * netease-im
 *
 * Created by lintao on 2020/6/9 2:39 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package netease

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/NSObjects/netease/path"
)

//accid	    String	是	加好友发起者accid
//faccid	String	是	加好友接收者accid
//type	    int	是	1直接加好友，2请求加好友，3同意加好友，4拒绝加好友
//msg	    String	否	加好友对应的请求消息，第三方组装，最长256字符
//serverex	String	否	服务器端扩展字段，限制长度256 此字段client端只读，server端读写
//ex		String	否	修改ex字段，限制长度256，可设置为空字符串
//alias		String	否	给好友增加备注名，限制长度128，可设置为空字符串
type Friend struct {
	Accid    string `json:"accid"`
	Faccid   string `json:"faccid"`
	Type     int    `json:"type"`
	Msg      string `json:"msg"`
	Serverex string `json:"serverex"`
	Alias    string `json:"alias"`
	Ex       string `json:"ex"`
}

//加好友
func (n NetEaseIM) AddFriend(friend Friend) error {
	//params := url.Values{}
	//params.Set("accid", accid)
	//params.Set("token", token)
	b, err := n.request(path.AddFriend, structToMap(friend))
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

//删除好友关系
//accid	String	是	发起者accid
//faccid	String	是	要删除朋友的accid
//isDeleteAlias	Boolean	否	是否需要删除备注信息 默认false:不需要，true:需要
func (n NetEaseIM) DeleteFriend(accid string, faccid string, isDeleteAlias string) error {
	params := url.Values{}
	params.Set("accid", accid)
	params.Set("faccid", faccid)
	params.Set("isDeleteAlias", isDeleteAlias)
	b, err := n.request(path.DeleteFriend, params)
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

type Bidirection struct {
	Createtime  int64  `json:"createtime"`
	Bidirection bool   `json:"bidirection"`
	Faccid      string `json:"faccid"`
	Alias       string `json:"alias"`
}

//获取好友关系 查询某时间点起到现在有更新的双向好友
//accid			String	是	发起者accid
//updatetime	Long	是	更新时间戳，接口返回该时间戳之后有更新的好友列表
func (n NetEaseIM) GetFriend(accid string, updatetime string) ([]Bidirection, error) {
	params := url.Values{}
	params.Set("accid", accid)
	params.Set("updatetime", updatetime)
	b, err := n.request(path.GetFriend, params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Code    int           `json:"code"`
		Friends []Bidirection `json:"friends" `
	}

	if err = json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		s, ok := stateCode[resp.Code]
		if ok {
			return nil, errors.New(s)
		}
	}

	return resp.Friends, nil
}

//更新好友相关信息，如加备注名，必须是好友才可以
func (n NetEaseIM) UpdateFriend(special Friend) error {
	b, err := n.request(path.UpdateFriend, structToMap(special))
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

//accid			String	是	用户帐号，最大长度32字符，必须保证一个 APP内唯一
//targetAcc	    String	是	被加黑或加静音的帐号
//relationType	int	    是	本次操作的关系类型,1:黑名单操作，2:静音列表操作
//value	        int	    是	操作值，0:取消黑名单或静音，1:加入黑名单或静音
type SpecialRelation struct {
	Accid        string `json:"accid"`
	TargetAcc    string `json:"target_acc"`
	RelationType int    `json:"relation_type"`
	Value        int    `json:"value"`
}

//拉黑/取消拉黑；设置静音/取消静音
func (n NetEaseIM) SetSpecialRelation(special SpecialRelation) error {
	b, err := n.request(path.SetSpecialRelation, structToMap(special))
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

//查看指定用户的黑名单和静音列表
// accid	String	是	用户帐号，最大长度32字符，必须保证一个 APP内唯一
func (n NetEaseIM) ListBlackAndMuteList(accid string) ([]string, []string, error) {
	params := url.Values{}
	params.Set("accid", accid)
	b, err := n.request(path.ListBlackAndMuteList, params)
	if err != nil {
		return nil, nil, err
	}

	var resp struct {
		Mutelist  []string `json:"mutelist"`
		Blacklist []string `json:"blacklist"`
		Code      int      `json:"code"`
	}

	if err = json.Unmarshal(b, &resp); err != nil {
		return nil, nil, err
	}

	if resp.Code != 200 {
		s, ok := stateCode[resp.Code]
		if ok {
			return nil, nil, errors.New(s)
		}
	}

	return resp.Mutelist, resp.Blacklist, nil
}
