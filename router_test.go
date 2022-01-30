package serverbox

import (
	"fmt"
	"net/http"
	"testing"
)

type msg struct {
	reply string
}

func routeHandler(userdata interface{}, res http.ResponseWriter,
	req *http.Request) {
	s, ok := userdata.(msg)
	if ok {
		fmt.Fprintf(res, s.reply)
	}
}

func TestRegisterRoute(t *testing.T) {
	r := NewRouter()
	r.RegisterRoute("/test", "get", routeHandler, msg{"test-DONE"})
}
