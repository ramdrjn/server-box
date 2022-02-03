package main

import (
	"fmt"
	sb "github.com/ramdrjn/serverbox"
)

type msg struct {
	reply string
}

func testRouteHandler(args *sb.HandlerArgs) {
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
	r := sb.NewRouter()
	r.RegisterRoute("/test", "get", testRouteHandler, msg{"test-DONE"})
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
