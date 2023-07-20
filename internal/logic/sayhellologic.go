package logic

import (
	"context"
	"os"

	"demo-rpc/demo"
	"demo-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SayHelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHelloLogic) SayHello(in *demo.InMsg) (*demo.OutMsg, error) {
	getenv := os.Getenv("POD_NAME")

	return &demo.OutMsg{Msg: "From " + getenv + ",you say " + in.Msg}, nil
}
