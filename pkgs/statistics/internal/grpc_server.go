package statistics

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func InitializeGrpcServer(stc *StatsContext) error {
	host := fmt.Sprintf("%s:%d", stc.Conf.Host, stc.Conf.Port)
	return StartGrpcServer(host)
}

func StartGrpcServer(host string) error {
	lis, err := net.Listen("tcp", host)
	if err != nil {
		Log.Error("grpc server listen failed: %v", err)
		return err
	}
	grpcServer := grpc.NewServer()
	RegisterService_SB_Stats(grpcServer)
	grpcServer.Serve(lis)
	return nil
}
