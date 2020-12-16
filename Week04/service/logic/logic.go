package logic

import (
	pb "Go-000/Week04/pub/idl/grpc/testcase"
	"Go-000/Week04/service/common"
	"Go-000/Week04/service/model"
	"context"
	"fmt"
	"github.com/google/wire"
)

type Service struct {
	usr model.UserDB
}

var Provider = wire.NewSet(NewService, model.Provider)

func NewService(usr model.UserDB) *Service {
	return &Service{usr}
}

// Echo ...
func (s *Service) Echo(ctx context.Context, req *pb.EchoReq) *pb.EchoRes {
	fun := "Service.Echo -->"
	return &pb.EchoRes{
		Errinfo:              fmt.Sprintf("%s -- %s", fun, common.ErrNotSupported),
		Data:                 nil,
	}
}

// HotFix ...
func (s *Service) HotFix(ctx context.Context, req *pb.HotFixReq) *pb.HotFixRes {
	fun := "Service.HotFix -->"
	return &pb.HotFixRes{
		Errinfo:              fmt.Sprintf("%s -- %s", fun, common.ErrNotSupported),
		Data:                 nil,
	}
}