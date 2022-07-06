package app

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"time"

	_ "prod_serv/docs"
	postgressql "prod_serv/internal/adapters/db/postgresql"
	"prod_serv/internal/config"
	v1 "prod_serv/internal/controller/http/v1"
	"prod_serv/internal/domain/service"
	regUsecase "prod_serv/internal/domain/usecase/regulation"
	"prod_serv/pkg/client/postgresql"
	"prod_serv/pkg/logging"
	"prod_serv/pkg/metric"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	cfg        *config.Config
	logger     *logging.Logger
	router     *httprouter.Router
	httpServer *http.Server
}

func NewApp(config *config.Config, logger *logging.Logger) (App, error) {
	logger.Println("router initializing")
	router := httprouter.New()

	logger.Println("swagger docs initializing")
	// hosting swagger specification
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	logger.Println("heartbeat initializing")
	metricHandler := metric.Handler{}
	metricHandler.Register(router)

	pgConfig := postgresql.NewPgConfig(
		config.PostgreSQL.Username, config.PostgreSQL.Password,
		config.PostgreSQL.Host, config.PostgreSQL.Port, config.PostgreSQL.Database,
	)

	pgClient, err := postgresql.NewClient(context.Background(), 5, time.Second*5, pgConfig)
	if err != nil {
		logger.Fatal(err)
	}
	regAdapter := postgressql.NewRegulationStorage(pgClient, logger)
	regService := service.NewRegulationService(regAdapter)
	regUsecase := regUsecase.NewRegulationUsecase(regService)
	regHandler := v1.NewRegulationHandler(regUsecase)
	regHandler.Register(router)

	// regulationStorage := storage.NewRegulationStorage(pgClient, logger)
	// all, err := regulationStorage.All(context.Background())
	// if err != nil {
	// 	logger.Fatal(err)
	// }
	// logger.Fatal(all)

	return App{cfg: config, logger: logger, router: router}, nil
}

func (a *App) Run() {
	a.startHTTP()
}

func (a *App) startHTTP() {
	a.logger.Info("start HTTP")

	// Define the listener (Unix or TCP)
	var listener net.Listener

	if a.cfg.Listen.Type == config.LISTEN_TYPE_SOCK {
		// Determine the current dirrectory
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			a.logger.Fatal(err)
		}
		// Determine the socket path
		socketPath := path.Join(appDir, a.cfg.Listen.SocketFile)
		a.logger.Infof("socket path: %s", socketPath)
		a.logger.Infof("creaet and listen unix socket")

		// start up a unix socket listener
		listener, err = net.Listen("unix", socketPath)
		if err != nil {
			a.logger.Fatal(err)
		}
	} else {
		a.logger.Infof("bind application to host: %s and port: %s", a.cfg.Listen.BindIP, a.cfg.Listen.Port)
		var err error
		// start up a tcp listener
		listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", a.cfg.Listen.BindIP, a.cfg.Listen.Port))
		if err != nil {
			a.logger.Fatal(err)
		}
	}

	// create a new Cors handler
	c := cors.New(cors.Options{
		AllowedMethods:     []string{http.MethodGet, http.MethodPost},
		AllowedOrigins:     []string{"http://localhost:10000"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Content-Type"},
		OptionsPassthrough: true,
		ExposedHeaders:     []string{"Access-Token", "Refresh-Token", "Location", "Authorization", "Content-Disposition"},
		Debug:              false,
	})

	// apply the CORS specification on the request, and add relevant CORS headers
	handler := c.Handler(a.router)

	// define parameters for an HTTP server
	a.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	a.logger.Println("application initialized and started")

	// accept incoming connections on the listener, creating a new service goroutine for each
	if err := a.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			a.logger.Warn("server shutdown")

		default:
			a.logger.Fatal(err)
		}
	}
	err := a.httpServer.Shutdown(context.Background())
	if err != nil {
		a.logger.Fatal(err)
	}
}
