package main

import (
	"fmt"
	"flag"
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
	debug:=flag.Bool("debug", false, "set true to enable debug mode")
	confFile:=flag.String("conf", "./sample.conf", "path of configuration file")
	ctx, err := sb.Initialize(*debug, *confFile)
	if err != nil {
		ctx.Log.Error(err)
	}
	sb.SetupSignalHandlers(ctx)
	r := mux.NewRouter()
	r.RegisterRoute("/test", "get", testRouteHandler, msg{"test-DONE"})
	err = sb.AttachRouter(r, "web", ctx)
	if err != nil {
		ctx.Log.Error(err)
	}
	err = sb.Run(ctx)
	if err != nil {
		ctx.Log.Error(err)
	}
	sb.BlockAndHandleSignal(ctx)
}
