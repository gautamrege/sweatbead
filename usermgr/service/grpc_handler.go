package service

import (
	"context"

	google_protobuf "github.com/golang/protobuf/ptypes"

	pb "github.com/gautamrege/packt/sweatbead/proto/usermgr"
	"github.com/gautamrege/packt/sweatbead/usermgr/db"
	"github.com/gautamrege/packt/sweatbead/usermgr/logger"
)

type GrpcServer struct {
	DB db.Storer
}

func (s *GrpcServer) GetUser(ctx context.Context, req *pb.UserRequest) (res *pb.UserResponse, err error) {
	user, err := s.DB.ByID(ctx, req.Userid)
	if err != nil {
		logger.Get().Info("Error fetching data", err)
	}

	res = &pb.UserResponse{
		Id:     req.Userid,
		Name:   user.Name,
		Device: user.Device,
	}

	return
}

func (s *GrpcServer) ListUsers(ctx context.Context, e *google_protobuf.Empty) (res *pb.UserResponse, err error) {
	users, err := s.DB.List(ctx)
	if err != nil {
		logger.Get().Info("Error fetching data", err)
	}

	res = &pb.UserResponse{
		Count: len(users),
		Users: []*pb.User{},
	}

	res.Userid = req.Userid
	for _, u := range users {
		tmp := pb.User{
			Id:     u.Userid,
			Name:   u.Name,
			Device: u.Device,
		}
		res.Users = append(res.Users, &tmp)
	}

	return
}
