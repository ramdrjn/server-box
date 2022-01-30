package serverbox

import (
	"fmt"
	"testing"
	"net/http"
)

func routeHandler(userdata interface{}, res http.ResponseWriter,
	req *http.Request) {
	s:=string(userdata)
	fmt.Fprintf(res, s)
}

func TestRegisterRoute(t *testing.T) {
	r:Router{}
	r.RegisterRoute("/test", "get", routeHandler, "test-DONE")
	if err != nil {
		t.Error(err)
	}

}
