package main

import (
	"fmt"
	sb "github.com/ramdrjn/serverbox"
	"github.com/ramdrjn/serverbox/pkgs/mux"
)

type msg struct {
	reply string
}

func testRouteHandler(args *mux.HandlerArgs) {
        s, ok := args.UserData.(msg)
        if ok {
                fmt.Fprintf(args.HttpRes, s.reply)
        }
}

func main() {
	ctx, err := sb.Initialize(true, "../../internal/sample.conf")
	if err != nil {
		ctx.Log.Error(err)
	}
	r := mux.NewRouter()
	r.RegisterRoute("/test", "GET", testRouteHandler, msg{"test-DONE"})
        err = sb.AttachRouter(r, "web", ctx)
        if err != nil {
		ctx.Log.Error(err)
        }
	err = sb.Run(ctx)
	if err != nil {
		ctx.Log.Error(err)
	}
	err = sb.ShutDown(ctx)
	if err != nil {
		ctx.Log.Error(err)
	}
}
