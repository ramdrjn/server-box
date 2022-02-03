package serverbox

import (
	. "github.com/ramdrjn/serverbox/internal"
	"testing"
	"github.com/ramdrjn/serverbox/pkgs/mux"
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
	r := mux.NewRouter()
	err := RouterAttach(r, "web", sbcontext)
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
