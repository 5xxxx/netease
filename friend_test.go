package netease

import (
	"reflect"
	"testing"
)

func TestNetEaseIM_ListBlackAndMuteList(t *testing.T) {

	type args struct {
		accid string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		want1   []string
		wantErr bool
	}{
		{
			name: "查询黑名单",
			args: args{
				accid: "5e836916e63913000134799b",
			},
			want:    nil,
			want1:   nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNetEaseIM(AppKey, Secret)
			n.SetDebug(true)

			got, got1, err := n.ListBlackAndMuteList(tt.args.accid)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListBlackAndMuteList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListBlackAndMuteList() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListBlackAndMuteList() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
