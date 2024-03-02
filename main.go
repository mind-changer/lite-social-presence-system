package main

import (
	"os"

	"github.com/lite-social-presence-system/config"
	"github.com/lite-social-presence-system/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	var err error
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	cfg := &config.Config{}
	err = config.LoadConfigFromEnv(cfg, ".env")
	if err != nil {
		logrus.Fatal("SERVER QUIT ERROR- Error while loading env config ", err.Error())
		return
	}
	go server.RunGRPCServer(cfg)
	server.RunHTTPServer(cfg)
}
