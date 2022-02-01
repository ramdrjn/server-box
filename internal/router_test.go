package serverbox

import (
	"fmt"
	"testing"
)

type msg struct {
	reply string
}

func testRouteHandler(args *HandlerArgs) {
	s, ok := args.UserData.(msg)
	if ok {
		fmt.Fprintf(args.HttpRes, s.reply)
	}
}

func TestRegisterRoute(t *testing.T) {
	r := NewRouter()
	r.RegisterRoute("/test", "get", testRouteHandler, msg{"test-DONE"})
}
