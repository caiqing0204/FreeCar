package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/CyanAsterisk/FreeCar/server/cmd/car/global"
	"github.com/CyanAsterisk/FreeCar/server/cmd/car/initialize"
	car "github.com/CyanAsterisk/FreeCar/server/cmd/car/kitex_gen/car/carservice"
	"github.com/CyanAsterisk/FreeCar/server/cmd/car/tool/sim"
	"github.com/CyanAsterisk/FreeCar/server/cmd/car/tool/trip"
	"github.com/CyanAsterisk/FreeCar/server/cmd/car/tool/ws"
	"github.com/CyanAsterisk/FreeCar/shared/middleware"
	hzserver "github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
)

func main() {
	// initialization
	initialize.InitLogger()
	IP, Port := initialize.InitFlag()
	initialize.InitConfig()
	initialize.InitDB()
	initialize.InitMq()
	initialize.InitTrip()
	initialize.InitCar()

	r, info := initialize.InitRegistry(Port)
	tracerSuite, closer := initialize.InitTracer()
	defer closer.Close()

	// Create new server.
	srv := car.NewServer(new(CarServiceImpl),
		server.WithServiceAddr(utils.NewNetAddr("tcp", fmt.Sprintf("%s:%d", IP, Port))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithSuite(tracerSuite),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: global.ServerConfig.Name}),
	)

	// Use goroutine to listen for signal.
	go func() {
		err := srv.Run()
		if err != nil {
			klog.Fatal(err)
		}
	}()

	h := hzserver.Default(hzserver.WithHostPorts(global.ServerConfig.WsAddr))
	h.GET("/ws", ws.Handler(global.Subscriber))
	h.NoHijackConnPool = true
	go func() {
		klog.Infof("HTTP server started. addr: %s", global.ServerConfig.WsAddr)
		h.Spin()
	}()

	go trip.RunUpdater(global.Subscriber, global.TripClient)

	simController := sim.Controller{
		CarService: global.CarClient,
		Subscriber: global.Subscriber,
	}
	go simController.RunSimulations(context.Background())

	// receive termination signal
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := r.Deregister(info); err != nil {
		klog.Info("sign out failed")
	} else {
		klog.Info("sign out success")
	}
}