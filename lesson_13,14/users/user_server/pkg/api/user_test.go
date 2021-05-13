package api

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/bukhavtsov/go-training-spring-2021/lesson_13,14/users/proto/go_proto"
)

func TestUserServer_HelloUser(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *pb.HelloRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.HelloResponse
		wantErr bool
	}{
		{
			name: "should work without error",
			args: args{request: &pb.HelloRequest{UserName: "test username"}},
			want: &pb.HelloResponse{Message: "Hello test username"},
			wantErr: false,
		},
		{
			name: "should now work, error",
			args: args{request: &pb.HelloRequest{UserName: ""}},
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserServer{}
			got, err := u.HelloUser(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("HelloUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HelloUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
