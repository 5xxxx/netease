/*
 *
 * msg_test.go
 * netease
 *
 * Created by lintao on 2020/6/10 10:53 上午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package netease

import (
	"reflect"
	"testing"
)

func TestNetEaseIM_SendBatchAttachMsg(t *testing.T) {
	type fields struct {
		appKey   string
		secret   string
		basePath string
	}
	type args struct {
		msg AttachMsg
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "测试批量发送消息",
			args: args{AttachMsg{
				FromAccid: "jsm_myfavors_session",
				ToAccids:  []string{"5efc4b1a301cf400019d53f4"},
				Attach:    "{`heihei`}",
				Save:      2,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			n := NewNetEaseIM(AppKey, Secret)
			if err := n.SendBatchAttachMsg(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("SendBatchAttachMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}

func TestNetEaseIM_SendMsg(t *testing.T) {
	type fields struct {
		appKey   string
		secret   string
		basePath string
	}
	type args struct {
		msg Msg
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    MsgResult
		wantErr bool
	}{
		{
			name: "发送消息",
			args: args{msg: Msg{
				From:             "",
				Ope:              0,
				To:               "",
				Type:             0,
				Body:             "",
				Antispam:         "",
				AntispamCustom:   "",
				Option:           "",
				Pushcontent:      "",
				Payload:          "",
				Ext:              "",
				Forcepushlist:    "",
				Forcepushcontent: "",
				Forcepushall:     "",
				Bid:              "",
				UseYidun:         0,
				MarkRead:         0,
				CheckFriend:      false,
			}},
			want:    MsgResult{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NetEaseIM{
				appKey:   tt.fields.appKey,
				secret:   tt.fields.secret,
				basePath: tt.fields.basePath,
			}
			got, err := n.SendMsg(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendMsg() got = %v, want %v", got, tt.want)
			}
		})
	}
}
