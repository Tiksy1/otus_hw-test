package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Tiksy1/otus_hw-test/hw12_13_14_15_calendar/internal/config"
	"github.com/Tiksy1/otus_hw-test/hw12_13_14_15_calendar/internal/logger"
	"github.com/Tiksy1/otus_hw-test/hw12_13_14_15_calendar/internal/server/http"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.json", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	cfg, err := config.NewConfig(configFile)
	if err != nil {
		log.Fatalf("can't get config: %v", err)
	}

	logg, err := logger.New(cfg.Logger.Level, cfg.Logger.FilePath)
	if err != nil {
		log.Fatalf("can't start logger %v\n", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := internalhttp.NewServer(internalhttp.NewTempPublic(logg), cfg.HTTPServer.Host, cfg.HTTPServer.Port, logg)

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals)

		<-signals
		signal.Stop(signals)
		cancel()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		logg.Info("stopping server...")
		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	logg.Info("starting server at " + server.Address)
	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		os.Exit(1) // nolint: gocritic
	}
}
