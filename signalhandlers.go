package serverbox

import (
	. "github.com/ramdrjn/serverbox/internal"
	"os"
	"os/signal"
	"syscall"
)

func SetupSignalHandlers(sbcontext *SbContext) {
	sbcontext.SignalChannel = make(chan os.Signal, 1)
	signal.Notify(sbcontext.SignalChannel, syscall.SIGINT, syscall.SIGTERM)
}

func BlockAndHandleSignal(sbcontext *SbContext) {
	sig := <-sbcontext.SignalChannel
	sbcontext.Log.Error("caught signal: ", sig)
	switch sig {
	case syscall.SIGINT:
		//handle SIGINT
		Abort(sbcontext)
	case syscall.SIGTERM:
		//handle SIGTERM
		ShutDown(sbcontext)
	}
}
