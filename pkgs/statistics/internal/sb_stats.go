package statistics

import (
	"context"
	"google.golang.org/grpc"
	pb "github.com/ramdrjn/serverbox/pkgs/statistics/pkgs/sb_stats_proto"
)

type statisticsServer struct {
	pb.UnimplementedStatisticsServer
}

func RegisterService_SB_Stats(grpcServer grpc.ServiceRegistrar) {
	pb.RegisterStatisticsServer(grpcServer, &statisticsServer{})
}

func (*statisticsServer) RegisterForStats(ctx context.Context, req *pb.RegisterReq) (res *pb.RegisterRes, err error) {
	r := &pb.RegisterRes{}
	r.Enrolled=true
	return r, nil
}
