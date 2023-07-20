// Code generated by goctl. DO NOT EDIT.
// Source: mydemo.proto

package server

import (
	"context"

	"demo-rpc/demo"
	"demo-rpc/internal/logic"
	"demo-rpc/internal/svc"
)

type DemoServer struct {
	svcCtx *svc.ServiceContext
	demo.UnimplementedDemoServer
}

func NewDemoServer(svcCtx *svc.ServiceContext) *DemoServer {
	return &DemoServer{
		svcCtx: svcCtx,
	}
}

func (s *DemoServer) SayHello(ctx context.Context, in *demo.InMsg) (*demo.OutMsg, error) {
	l := logic.NewSayHelloLogic(ctx, s.svcCtx)
	return l.SayHello(in)
}