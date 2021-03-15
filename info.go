/*
 *
 * info.go
 * netease-im
 *
 * Created by lintao on 2020/6/9 2:18 下午
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

//更新用户名片。用户名片中包含的用户信息，在群组、聊天室等场景下，会暴露给群组、聊天室内的其他用户。
//这些字段里mobile，email，birth，gender等字段属于非必填、可能涉及隐私的信息，如果您的业务下，这些信息为敏感信息
//，建议在通过扩展字段ex填写相关资料并事先加密
func (n NetEaseIM) UpdateUinfo(account Account) error {
	b, err := n.request(path.UpdateUinfo, structToMap(account))
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

//获取用户名片，可批量
//accids	String	是	用户帐号（例如：JSONArray对应的accid串，如：["zhangsan"]，如果解析出错，会报414）（一次查询最多为200）
func (n NetEaseIM) GetUinfo(accids string) ([]Account, error) {
	params := url.Values{}
	params.Set("accids", "[\""+accids+"\"]")

	b, err := n.request(path.GetUinfos, params)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Code   int       `json:"code"`
		Uinfos []Account `json:"uinfos" `
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

	return resp.Uinfos, nil
}
