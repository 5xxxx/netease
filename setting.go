/*
 *
 * setting.go
 * netease-im
 *
 * Created by lintao on 2020/6/9 2:33 下午
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

//设置桌面端在线时，移动端是否需要推送
//accid	String	是	用户帐号
//donnopOpen	String	是	桌面端在线时，移动端是否不推送：
//true:移动端不需要推送，false:移动端需要推送
func (n NetEaseIM) SetDonnop(accid, donnopOpen string) error {
	params := url.Values{}
	params.Set("accid", accid)
	params.Set("donnopOpen", donnopOpen)
	b, err := n.request(path.SetDonnop, params)
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

//设置或取消账号的全局禁言状态；
//账号被设置为全局禁言后，不能发送“点对点”、“群”、“聊天室”消息
//accid	String	是	用户帐号
//mute	Boolean	是	是否全局禁言：
//true：全局禁言，false:取消全局禁言
func (n NetEaseIM) Mute(accid, mute string) error {
	params := url.Values{}
	params.Set("accid", accid)
	params.Set("mute", mute)
	b, err := n.request(path.Mute, params)
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

//账号全局禁用音视频
//设置或取消账号是否可以发起音视频功能；
//账号被设置为禁用音视频后，不能发起点对点音视频、创建多人音视频、发起点对点白板、创建多人白板
func (n NetEaseIM) MuteAV(accid, mute string) error {
	params := url.Values{}
	params.Set("accid", accid)
	params.Set("mute", mute)
	b, err := n.request(path.MuteAv, params)
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
