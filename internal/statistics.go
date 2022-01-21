package serverbox

import (
	pb "github.com/ramdrjn/serverbox/pkgs/statistics/pkgs/sb_stats_proto"
	"google.golang.org/grpc"
)

type Statistics struct {
	conn  *grpc.ClientConn
	stats pb.StatisticsClient
}

func getStatsClient(conn *grpc.ClientConn) (pb.StatisticsClient, error) {
	cli := pb.NewStatisticsClient(conn)
	return cli, nil
}

func InitializeStatistics(sbc *SbContext) (statistics *Statistics, err error) {
	if sbc.Conf.Stats.Enabled == false {
		sbc.Log.Debugln("statistics not configured")
		return nil, nil
	}

	host := fmt.Sprintf("%s:%d", sbc.Conf.Stats.Host, sbc.Conf.Stats.Port)

	conn, err := grpc.Dial(host, nil)
	if err != nil {
		sbc.Log.Error("server connect fail: %v", err)

		//TODO - Retry connection.

		return nil, err
	}

	cli, _ := getStatsClient(conn)

	statistics = &Statistics{conn, cli}

	return statistics, err
}

func ShutDownStatistics(sbc *SbContext) (err error) {
	if sbc.Conf.Stats.Enabled == false {
		sbc.Log.Debugln("statistics not configured")
		return nil, nil
	}

	sbc.Stats.conn.Close()

	return nil
}
