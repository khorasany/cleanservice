package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kianooshaz/cleanservice/internal/config"
	"github.com/kianooshaz/cleanservice/internal/pkg/logger/zap"
	"github.com/kianooshaz/cleanservice/internal/repository/mysql"
	"github.com/kianooshaz/cleanservice/internal/service/account"
	"github.com/kianooshaz/cleanservice/internal/transport/http/echo"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap/zapcore"
)

var serveCMD = &cli.Command{
	Name:    "serve",
	Aliases: []string{"s"},
	Usage:   "serve http",
	Action:  serve,
}

func serve(c *cli.Context) error {
	cfg := new(config.Config)
	config.ReadYAML("config.yaml", cfg)
	config.ReadEnv(cfg)

	f, err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	logger := zap.New(f, zapcore.ErrorLevel)

	mysqlRepo, err := mysql.New(cfg.Mysql, logger)
	if err != nil {
		return err
	}

	accSrv := account.New(cfg.Account, mysqlRepo, logger)

	restServer := echo.New(logger, accSrv)
	go func() {
		if err := restServer.Start(cfg.App.Address); err != nil {
			logger.Error(fmt.Sprintf("error happen while serving: %v", err))
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("\nReceived an interrupt, closing connections...")

	if err := restServer.Shutdown(); err != nil {
		fmt.Println("\nRest server doesn't shutdown in 10 seconds")
	}

	return nil
}
