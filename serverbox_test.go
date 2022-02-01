package serverbox

import (
	. "github.com/ramdrjn/serverbox/internal"
	"github.com/ramdrjn/serverbox/pkgs/mux"
	"testing"
)

var sbcontext *SbContext

func TestInit(t *testing.T) {
	var err error
	sbcontext, err = Initialize(true, "./internal/sample.conf")
	if err != nil {
		t.Error(err)
	}
}

func TestRouteAttach(t *testing.T) {
	r:=new(mux.Router)
	err := AttachRouter(r, "web", sbcontext)
	if err != nil {
		t.Error(err)
	}
}

func TestRun(t *testing.T) {
	err := Run(sbcontext)
	if err != nil {
		t.Error(err)
	}
}

func TestShutDown(t *testing.T) {
	err := ShutDown(sbcontext)
	if err != nil {
		t.Error(err)
	}
}
