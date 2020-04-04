package rpc

import (
	"context"
	"testing"
	"time"

	"github.com/txze/visitorpb/go"
	"google.golang.org/grpc"
)

//该测试依赖服务启动 只用于个人测试使用 结果不能作为单元测试参照
func TestService_AddVisitor(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	c := visitorpb.NewVisitorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	reply, err := c.AddVisitor(ctx, &visitorpb.AddVisitorRequest{
		Id:       "ddd",
		Unixnano: 1,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(reply)
}
