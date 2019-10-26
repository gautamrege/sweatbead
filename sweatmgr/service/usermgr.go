package service

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc"

	usermgr "github.com/gautamrege/packt/sweatbead/proto/usermgr"
	"github.com/gautamrege/packt/sweatbead/sweatmgr/db"
	"github.com/gautamrege/packt/sweatbead/sweatmgr/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	serverAddr = "127.0.0.1:33011"
)

type UserMgrSvc interface {
	GetUser(id string) (user db.User, err error)
}

type userSvc struct {
}

func GetUserMgr() UserMgrSvc {
	return &userSvc{}
}

func (svc *userSvc) GetUser(id string) (user db.User, err error) {
	logger.Get().Info("Dialing UserMgr")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		logger.Get().Error("fail to dial: %v", err)
	}
	defer conn.Close()
	client := usermgr.NewUserMgrClient(conn)

	logger.Get().Info("Getting user %s", id)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if id == "" { // no user specified!
		logger.Get().Error("No user specified in request")
		err = errors.New("No user specified in request")
		return
	}

	req := usermgr.UserRequest{id}
	res, err := client.GetUser(ctx, &req)
	if err != nil {
		logger.Get().Error("%v.GetUser(_) = _, %v: ", client, err)
		return
	}

	if res.Count != 1 {
		logger.Get().Error("User not found")
		err = errors.New("User not found: " + id)
		return
	}

	first := res.Users[0]

	obj, _ := primitive.ObjectIDFromHex(first.Id)
	user = db.User{obj, first.Name, first.Device}

	logger.Get().Info("User found: %v", user)

	return
}
