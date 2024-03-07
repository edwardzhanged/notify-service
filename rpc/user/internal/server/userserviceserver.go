// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"notify-service/rpc/user/internal/logic"
	"notify-service/rpc/user/internal/svc"
	"notify-service/rpc/user/user"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

// 查询用户信息
func (s *UserServiceServer) GetUser(ctx context.Context, in *user.GetUserRequest) (*user.GetUserResponse, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}