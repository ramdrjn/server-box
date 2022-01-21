package serverbox

import (
	"errors"
	"fmt"
	pb "github.com/ramdrjn/serverbox/pkgs/statistics/pkgs/sb_stats_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Statistics struct {
	conn  *grpc.ClientConn
	stats pb.StatisticsClient
}

func getStatsClient(conn *grpc.ClientConn) (pb.StatisticsClient, error) {
	if conn == nil {
		Log.Error("statistics server connection not valid")
		return nil, errors.New("connection not valid")
	}
	cli := pb.NewStatisticsClient(conn)
	return cli, nil
}

func InitializeStatistics(sbc *SbContext) (statistics *Statistics, err error) {
	var opts []grpc.DialOption
	if sbc.Conf.Statistics.Enabled == false {
		sbc.Log.Debugln("init: statistics not configured")
		return nil, nil
	}

	host := fmt.Sprintf("%s:%d", sbc.Conf.Statistics.Host,
		sbc.Conf.Statistics.Port)

	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	sbc.Log.Debug("dialling: ", host)
	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		sbc.Log.Error("server connect fail: ", err)

		//TODO - Retry connection.

		return nil, err
	}

	cli, _ := getStatsClient(conn)

	statistics = &Statistics{conn, cli}

	return statistics, err
}

func ShutDownStatistics(sbc *SbContext) (err error) {
	if sbc.Conf.Statistics.Enabled == false {
		sbc.Log.Debugln("shut: statistics not configured")
		return nil
	}

	sbc.Stats.conn.Close()

	return nil
}
