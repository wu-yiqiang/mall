// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"mall/service/order/api/internal/config"
	"mall/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRPC userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRPC: userclient.NewUser(zrpc.MustNewClient(c.UserRPC)),
		// zrpc.WithUnaryClientInterceptor(interceptor)
	}
}

//func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.Invoker) (interface{}, error) {
//	fmt.Println("拦截前")
//	md := metadata.Pairs("authorization", ctx.Value("token").(string))
//	err := invoker(ctx, method, req, reply, cc, opts...)
//	//fmt.Println("拦截后")
//	//return nil, err
//	return handler(ctx, req)
//}
