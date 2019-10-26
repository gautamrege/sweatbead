package service

import (
	"context"

	gp_empty "github.com/golang/protobuf/ptypes/empty"

	pb "github.com/gautamrege/packt/sweatbead/proto/usermgr"
	"github.com/gautamrege/packt/sweatbead/usermgr/db"
	"github.com/gautamrege/packt/sweatbead/usermgr/logger"
)

type GrpcServer struct {
	DB db.Storer
}

func (s *GrpcServer) GetUser(ctx context.Context, req *pb.UserRequest) (res *pb.UserResponse, err error) {
	logger.Get().Info("grpc: Getting user")
	user, err := s.DB.ByID(ctx, req.Userid)
	if err != nil {
		logger.Get().Info("Error fetching data", err)
	}

	res = &pb.UserResponse{
		Count: 1,
		Users: []*pb.User{{req.Userid, user.Name, user.Device}},
	}

	return
}

func (s *GrpcServer) ListUsers(ctx context.Context, e *gp_empty.Empty) (res *pb.UserResponse, err error) {
	users, err := s.DB.List(ctx)
	if err != nil {
		logger.Get().Info("Error fetching data", err)
	}

	res = &pb.UserResponse{
		Count: int32(len(users)),
		Users: []*pb.User{},
	}

	res.Count = int32(len(users))
	for _, u := range users {
		tmp := pb.User{
			Id:     u.ID.Hex(),
			Name:   u.Name,
			Device: u.Device,
		}
		res.Users = append(res.Users, &tmp)
	}

	return
}
