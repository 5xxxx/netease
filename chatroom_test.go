/*
 *
 * chatroom_test.go
 * netease
 *
 * Created by lintao on 2020/7/15 1:37 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package netease

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"go.uber.org/zap"
)

func TestNetEaseIM_AddRobot(t *testing.T) {
	type fields struct {
		appKey   string
		secret   string
		basePath string
		logger   *zap.Logger
		debug    bool
	}
	type args struct {
		r ChatroomRobot
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NetEaseIM{
				appKey:   tt.fields.appKey,
				secret:   tt.fields.secret,
				basePath: tt.fields.basePath,
				debug:    tt.fields.debug,
			}
			if err := n.AddRobot(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("AddRobot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNetEaseIM_CreateChatroom(t *testing.T) {

	type args struct {
		room Chatroom
	}
	tests := []struct {
		name    string
		args    args
		want    Chatroom
		wantErr bool
	}{
		{
			name: "创建聊天室",
			args: args{room: Chatroom{
				Creator:      "5efc4b1a301cf400019d53f4",
				Name:         "哈哈哈",
				Announcement: "欢迎来到我的直播间",
				Queuelevel:   1,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNetEaseIM(AppKey, Secret)
			n.SetDebug(true)

			got, err := n.CreateChatroom(tt.args.room)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateChatroom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("CreateChatroom() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestNetEaseIM_GetBatchChatroom(t *testing.T) {
	type fields struct {
		appKey   string
		secret   string
		basePath string
		logger   *zap.Logger
		debug    bool
	}
	type args struct {
		roomids             string
		needOnlineUserCount string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Chatroom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NetEaseIM{
				appKey:   tt.fields.appKey,
				secret:   tt.fields.secret,
				basePath: tt.fields.basePath,

				debug: tt.fields.debug,
			}
			got, err := n.GetBatchChatroom(tt.args.roomids, tt.args.needOnlineUserCount)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBatchChatroom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBatchChatroom() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNetEaseIM_GetChatroom(t *testing.T) {
	type fields struct {
		appKey   string
		secret   string
		basePath string
		logger   *zap.Logger
		debug    bool
	}
	type args struct {
		roomid              string
		needOnlineUserCount string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Chatroom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NetEaseIM{
				appKey:   tt.fields.appKey,
				secret:   tt.fields.secret,
				basePath: tt.fields.basePath,

				debug: tt.fields.debug,
			}
			got, err := n.GetChatroom(tt.args.roomid, tt.args.needOnlineUserCount)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChatroom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChatroom() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNetEaseIM_RequestAddr(t *testing.T) {
	type fields struct {
		appKey   string
		secret   string
		basePath string
		logger   *zap.Logger
		debug    bool
	}
	type args struct {
		r ChatroomRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NetEaseIM{
				appKey:   tt.fields.appKey,
				secret:   tt.fields.secret,
				basePath: tt.fields.basePath,

				debug: tt.fields.debug,
			}
			got, err := n.RequestAddr(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestAddr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestAddr() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNetEaseIM_SendChatRoomMsg(t *testing.T) {
	type args struct {
		r ChatroomMsg
	}
	tests := []struct {
		name string

		args    args
		wantErr bool
	}{
		{
			name: "发送聊天室消息",
			args: args{r: ChatroomMsg{
				Roomid:    "xx",
				FromAccid: "xx",
				MsgType:   0,
				Attach:    " {\n  \"msg\":\"哈哈哈\"\n}",
				MsgId:     fmt.Sprintf("%d", rand.Int63()),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := netease.SendChatRoomMsg(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("SendChatRoomMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNetEaseIM_SetMemberRole(t *testing.T) {
	type fields struct {
		appKey   string
		secret   string
		basePath string
		logger   *zap.Logger
		debug    bool
	}
	type args struct {
		r MemberRole
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NetEaseIM{
				appKey:   tt.fields.appKey,
				secret:   tt.fields.secret,
				basePath: tt.fields.basePath,

				debug: tt.fields.debug,
			}
			if err := n.SetMemberRole(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("SetMemberRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNetEaseIM_ToggleCloseStat(t *testing.T) {
	type fields struct {
		appKey   string
		secret   string
		basePath string
		logger   *zap.Logger
		debug    bool
	}
	type args struct {
		roomid   string
		operator string
		valid    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NetEaseIM{
				appKey:   tt.fields.appKey,
				secret:   tt.fields.secret,
				basePath: tt.fields.basePath,

				debug: tt.fields.debug,
			}
			if err := n.ToggleCloseStat(tt.args.roomid, tt.args.operator, tt.args.valid); (err != nil) != tt.wantErr {
				t.Errorf("ToggleCloseStat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNetEaseIM_UpdateChatroom(t *testing.T) {
	type fields struct {
		appKey   string
		secret   string
		basePath string
		logger   *zap.Logger
		debug    bool
	}
	type args struct {
		room Chatroom
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NetEaseIM{
				appKey:   tt.fields.appKey,
				secret:   tt.fields.secret,
				basePath: tt.fields.basePath,

				debug: tt.fields.debug,
			}
			if err := n.UpdateChatroom(tt.args.room); (err != nil) != tt.wantErr {
				t.Errorf("UpdateChatroom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
