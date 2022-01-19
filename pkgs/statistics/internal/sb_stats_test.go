package statistics

import (
	"context"
	pb "github.com/ramdrjn/serverbox/pkgs/statistics/pkgs/sb_stats_proto"
	"testing"
)

var s = &statisticsServer{}

func TestRegisterForStats(t *testing.T) {
	req := &pb.RegisterReq{}
	req.Uuid = "test"
	req.Type = pb.RegisterReq_SERVER
	s.RegisterForStats(context.TODO(), req)
}
