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
	"os"
	"testing"
)

func init() {
	AppKey = os.Getenv("app_key")
	Secret = os.Getenv("secret")
}

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
				FromAccid: "13802426870",
				ToAccids:  []string{"15675132016"},
				Msgtype:   0,
				Attach:    "{`heihei`}",
				Save:      2,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNetEaseIM("1823d2e114e071e7670cdd520e76cecf", "8bdf74978155")
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
				From: "5e82ddaf7da7a00001c12ef8",
				Ope:  0,
				To:   "5e84658857b24c00011fc724",
				Type: 100,
				Body: "{\n  \"msg\":\"傻屌\"//消息内容\n}",
			}},
			want:    MsgResult{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNetEaseIM(AppKey, Secret)
			_, err := n.SendMsg(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestNetEaseIM_SendAttachMsg(t *testing.T) {
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
			name: "",
			args: args{msg: AttachMsg{
				From:    "13802426870",
				Msgtype: 0,
				Save:    2,
				To:      "15675132016",
				Attach:  "{`sd`:`哈哈哈`}",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNetEaseIM(AppKey, Secret)
			if err := n.SendAttachMsg(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("SendAttachMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
