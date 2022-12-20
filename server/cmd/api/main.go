package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/CyanAsterisk/FreeCar/server/cmd/api/global"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/initialize"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/initialize/rpc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertztracer "github.com/hertz-contrib/tracer/hertz"
)

func main() {
	// initialize
	initialize.InitLogger()
	initialize.InitConfig()
	r, info := initialize.InitRegistry()
	tracer, closer := initialize.InitTracer()
	defer closer.Close()
	rpc.Init()
	// create a new server
	h := server.New(
		server.WithHostPorts(fmt.Sprintf(":%d", global.ServerConfig.Port)),
		server.WithTracer(hertztracer.NewTracer(tracer, func(c *app.RequestContext) string {
			return "FreeCar.server" + "::" + c.FullPath()
		})),
		server.WithRegistry(r, info),
	)

	h.Use(hertztracer.ServerCtx())
	register(h)
	// Use goroutine to listen for signal.
	go func() {
		if err := h.Run(); err != nil {
			hlog.Fatalf("start error:", err.Error())
		}
	}()

	//  receive termination signal
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := r.Deregister(info); err != nil {
		hlog.Info("sign out failed")
	} else {
		hlog.Info("sign out success")
	}
	h.Spin()
}
