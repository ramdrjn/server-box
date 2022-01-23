package state

import (
	"context"
	pb "github.com/ramdrjn/serverbox/pkgs/state/pkgs/sb_state_proto"
	"google.golang.org/grpc"
)

type stateServer struct {
	pb.UnimplementedStateServer
}

func RegisterService_SB_State(grpcServer grpc.ServiceRegistrar) {
	pb.RegisterStateServer(grpcServer, &stateServer{})
}

func (*stateServer) RegisterForState(ctx context.Context, req *pb.RegisterReq) (res *pb.RegisterRes, err error) {
	r := &pb.RegisterRes{}
	return r, nil
}

func (*stateServer) ReportState(ctx context.Context, req *pb.ReportReq) (res *pb.ReportRes, err error) {
	r := &pb.ReportRes{}
	return r, nil
}
