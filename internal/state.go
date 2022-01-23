package serverbox

import (
	"context"
	"errors"
	pb "github.com/ramdrjn/serverbox/pkgs/state/pkgs/sb_state_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type State struct {
	uuid    string
	conn    *grpc.ClientConn
	state   pb.StateClient
	enabled bool
}

func InitializeState(uuid string, host string, state *State) error {
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

	cli := pb.NewStateClient(conn)

	state.uuid = uuid
	state.conn = conn
	state.state = cli
	state.enabled = true

	return err
}

func ShutDownState(state *State) error {
	if state.enabled == false {
		Log.Debugln("shut: state not configured")
		return nil
	}

	state.conn.Close()

	return nil
}

func (s *State) RegisterForState() error {
	req := &pb.RegisterReq{Uuid: s.uuid, Type: pb.RegisterReq_SERVER}
	ctx := context.TODO()
	_, err := s.state.RegisterForState(ctx, req)
	if err != nil {
		Log.Error("registration failed for server")
	}
	return err
}

func convertState(state string) (pb.ReportReq_State, error) {
	switch state {
	case "up":
		return pb.ReportReq_UP, nil
	case "down":
		return pb.ReportReq_DOWN, nil
	case "maintanence":
		return pb.ReportReq_MAINTANENCE, nil
	}
	return pb.ReportReq_MAINTANENCE, errors.New("invalid state")
}

func (s *State) ReportState(state string) error {
	stateVal, err := convertState(state)
	if err != nil {
		Log.Error(err)
	}
	req := &pb.ReportReq{TargetUuid: s.uuid, State: stateVal}
	ctx := context.TODO()
	_, err = s.state.ReportState(ctx, req)
	if err != nil {
		Log.Error("report failed for server")
	}
	return err
}
