package main

import (
	"context"
	"flag"
	"go-gin/pkg/logger"
	"go-gin/until"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"

	"go-gin/internal/app"
)

var VERSION = "4.0.0"

var (
	configFile string
	modelFile  string
	wwwDir     string
	swaggerDir string
)

func init() {
	flag.StringVar(&configFile, "c", "./config/config.toml", "配置文件(.json,.yaml,.toml)")
	//flag.StringVar(&wwwDir, "www", "", "静态站点目录")
	//flag.StringVar(&swaggerDir, "swagger", "", "swagger目录")
}

func main() {

	flag.Parse()
	if configFile == "" {
		panic("请使用-c指定配置文件")
	}
	var state int32 = 1
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	ctx := logger.NewTraceIDContext(context.Background(), until.MustUUID())
	span := logger.StartSpanWithCall(ctx)
	call := app.Init(ctx,
		app.SetConfigFile(configFile),
		app.SetModelFile(modelFile),
		app.SetWWWDir(wwwDir),
		app.SetVersion(VERSION))
	select {
	case sig := <-sc:
		atomic.StoreInt32(&state, 0)
		span().Printf("获取到退出信号[%s]", sig.String())
	}

	if call != nil {
		call()
	}
	span().Printf("服务退出")

	os.Exit(int(atomic.LoadInt32(&state)))
}
