package logic

import (
	"context"
	"errors"

	"go-cloud-disk/core/internal/svc"
	"go-cloud-disk/core/internal/types"
	"go-cloud-disk/core/models"
	"go-cloud-disk/helper"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// 1. 从数据库中查询当前用户
	user := new(models.UserBasic)
	has, err := models.Engine.Where("name = ? AND password = ?", req.Name, helper.Md5(req.Password)).Get(user)
	if err != nil {
		return
	}
	// 用户不存在
	if !has {
		err = errors.New("用户名或密码错误")
		return
	}
	// 2. 生成 token
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name)
	if err != nil {
		return
	}
	resp = new(types.LoginReply)
	resp.Token = token

	return
}
