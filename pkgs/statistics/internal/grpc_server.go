package statistics

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

var grpcServer *grpc.Server

func InitializeGrpcServer(stc *StatsContext) error {
	host := fmt.Sprintf("%s:%d", stc.Conf.Host, stc.Conf.Port)

	lis, err := net.Listen("tcp", host)
	if err != nil {
		stc.Log.Errorf("grpc server on %s listen failed: %s",
			host, err)
		return err
	}

	grpcServer = grpc.NewServer()

	RegisterService_SB_Stats(grpcServer)

	stc.Log.Infof("starting grpc server on %s", host)
	grpcServer.Serve(lis)

	return nil
}

func ShutDownGrpcServer(stc *StatsContext) error {
	stc.Log.Info("stopping grpc server")

	grpcServer.GracefulStop()

	return nil
}

func AbortGrpcServer(stc *StatsContext) error {
	stc.Log.Info("aborting grpc server")

	grpcServer.Stop()

	return nil
}
