package logic

import (
	"context"
	"errors"

	"zero-demo/mall/order/api/internal/svc"
	"zero-demo/mall/order/api/internal/types"
	"zero-demo/mall/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
////////////////////
func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.IdRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	if user.Name != "test" {
		return nil, errors.New("用户不存在")
	}
	return &types.OrderReply{
		Id:   req.Id,
		Name: "test order",
	}, nil
}
