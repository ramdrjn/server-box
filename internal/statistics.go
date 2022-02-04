package serverbox

import (
	"context"
	"errors"
	pb "github.com/ramdrjn/serverbox/pkgs/statistics/pkgs/sb_stats_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Statistics struct {
	uuid       string
	conn       *grpc.ClientConn
	statistics pb.StatisticsClient
	enabled    bool
}

func InitializeStatistics(uuid string, host string, stats *Statistics) error {
	var opts []grpc.DialOption

	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())

	Log.Debug("dialling: ", host)
	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		Log.Error("server connect fail: ", err)
		return err
	}

	cli := pb.NewStatisticsClient(conn)

	stats.uuid = uuid
	stats.conn = conn
	stats.statistics = cli
	stats.enabled = true

	return err
}

func ShutDownStatistics(stats *Statistics) error {
	if stats.enabled == false {
		Log.Debugln("shut: statistics not configured")
		return nil
	}

	stats.conn.Close()

	return nil
}

func (s *Statistics) RegisterForStats() error {
	if s.enabled == false {
		return nil
	}
	req := &pb.RegisterReq{Uuid: s.uuid, Type: pb.RegisterReq_SERVER}
	ctx := context.TODO()
	res, err := s.statistics.RegisterForStats(ctx, req)
	if err != nil {
		Log.Error("registration failed for server")
	}
	if err == nil && res.Enrolled == false {
		Log.Error("registration not enrolled for server")
		err = errors.New("registration not enrolled")
	}
	return err
}
