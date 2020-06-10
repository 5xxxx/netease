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

import "testing"

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
				ToAccids:  []string{"9527111"},
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
