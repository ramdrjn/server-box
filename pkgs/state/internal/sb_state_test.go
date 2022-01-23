package state

import (
	"context"
	pb "github.com/ramdrjn/serverbox/pkgs/state/pkgs/sb_state_proto"
	"testing"
)

var s = &stateServer{}

func TestRegisterForState(t *testing.T) {
	req := &pb.RegisterReq{}
	req.Uuid = "test"
	req.Type = pb.RegisterReq_SERVER
	s.RegisterForState(context.TODO(), req)
}
