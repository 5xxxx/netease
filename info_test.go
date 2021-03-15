package netease

import (
	"testing"
)

func TestNetEaseIM_GetUinfo(t *testing.T) {

	type args struct {
		accids string
	}
	tests := []struct {
		name string

		args    args
		want    []Account
		wantErr bool
	}{
		{
			name:    "查询用户信息",
			args:    args{accids: "5ff57bbdc608c5a2ea08c91f"},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNetEaseIM(AppKey, Secret)
			n.SetDebug(true)
			_, err := n.GetUinfo(tt.args.accids)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUinfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("GetUinfo() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
