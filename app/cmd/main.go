package main

import (
	"log"

	"prod_serv/internal/app"
	"prod_serv/internal/config"
	"prod_serv/pkg/logging"
	// "github.com/i-b8o/regulations_be"
	// "github.com/i-b8o/regulations_be/pkg/handler"
	// "github.com/sirupsen/logrus"
)

// @title Regulations API
// @version 1.0
// @description API Server for Regulations

// @host 188.93.210.165:8080
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	log.Print("config initializing")
	cfg := config.GetConfig()
	log.Print("logger initializing")

	logger := logging.GetLogger(cfg.AppConfig.LogLevel)

	a, err := app.NewApp(cfg, &logger)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Running Application")

	a.Run()
	// logrus.SetFormatter(new(logrus.JSONFormatter))

	// handlers := handler.NewHandler()

	// // init server instance
	// srv := new(regulations_be.Server)
	// // run server
	// go func() {
	// 	if err := srv.Run(os.Getenv("REGULATIONS_API_SERVICE_PORT"), handlers.InitRoutes()); err != nil {
	// 		logrus.Fatalf("error occured while running http server: %s", err.Error())
	// 	}
	// }()
	// logrus.Print("REGULATIONS started")

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	// <-quit

	// logrus.Print("REGULATIONS Shutting Down")

	// if err := srv.Shutdown(context.Background()); err != nil {
	// 	logrus.Errorf("error occured on server shutting down: %s", err.Error())
	// }

}
