package service

import (
	"context"

	google_protobuf "github.com/golang/protobuf/ptypes"

	pb "github.com/gautamrege/packt/sweatbead/proto/sweatmgr"
	"github.com/gautamrege/packt/sweatbead/sweatmgr/db"
	"github.com/gautamrege/packt/sweatbead/sweatmgr/logger"
)

type GrpcServer struct {
}

func (s *GrpcServer) GetSweatStats(ctx context.Context, req *pb.SweatStatsRequest) (res *pb.SweatStatsResponse, err error) {
	ctx = context.WithValue(ctx, "UserID", req.Userid)

	sweats, err := db.ListUserSweat(ctx)
	if err != nil {
		logger.Get().Info("Error fetching data", err)
	}

	res = &pb.SweatStatsResponse{
		Sweat: []*pb.Sweat{},
	}

	res.Userid = req.Userid
	for _, sw := range sweats {
		tmp := pb.Sweat{
			Glucose:         sw.Glucose,
			Chloride:        sw.Chloride,
			Sodium:          sw.Sodium,
			Potassium:       sw.Potassium,
			Magnesium:       sw.Magnesium,
			Calcium:         sw.Calcium,
			Humidity:        sw.Humidity,
			RoomTemperature: sw.RoomTemperatureF,
			BodyTemperature: sw.BodyTemperatureF,
			Heartbeat:       sw.HeartBeat,
		}
		tmp.CreatedAt, _ = google_protobuf.TimestampProto(sw.CreatedAt) // ignore conv. error (default: nil timestamp)
		res.Sweat = append(res.Sweat, &tmp)
	}

	return
}
