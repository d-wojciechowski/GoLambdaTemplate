package main

import (
	"context"
	"reflect"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	type args struct {
		ctx   context.Context
		event MyEvent
	}
	tests := []struct {
		name    string
		args    args
		want    *MyResponse
		wantErr bool
	}{
		{
			name:    "Should return OK when its done correctly",
			args:    args{event: MyEvent{Url: "http://www.google.com"}},
			want:    &MyResponse{Message: "200 OK"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleRequest(tt.args.ctx, tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
