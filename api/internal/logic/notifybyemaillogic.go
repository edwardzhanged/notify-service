package logic

import (
	"context"

	"notify-service/api/internal/svc"
	"notify-service/api/internal/types"

	"log"
	pb "notify-service/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
)

type NotifyByEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNotifyByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyByEmailLogic {
	return &NotifyByEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NotifyByEmailLogic) NotifyByEmail(req *types.NotifyByEmailReq) (resp *types.NotifyByEmailResp, err error) {
	// todo: add your logic here and delete this line
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	// Call a method on the client
	user, err := c.GetUser(context.Background(), &pb.GetUserRequest{UserId: req.Msg})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("User: %s", user.Profile.Send)
	return
}
