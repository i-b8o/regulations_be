package app

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"prod_serv/internal/config"
	"prod_serv/pkg/logging"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	cfg    *config.Config
	logger *logging.Logger
}

func NewApp(config *config.Config, logger *logging.Logger) (App, error) {
	logger.Println("router initializing")
	router := httprouter.New()

	logger.Println("swagger docs initializing")
	// hosting swagger specification
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)
	return App{cfg: config, logger: logger}, nil
}

func (a *App) startHTTP() {
	a.logger.Info("start HTTP")

	// Define the listener (Unix or TCP)
	var listener net.Listener

	if a.cfg.Listen.Type == "sock" {
		// Determine the current dirrectory
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			a.logger.Fatal(err)
		}
		// Determine the socket path
		socketPath := path.Join(appDir, a.cfg.Listen.SocketFile)
		a.logger.Infof("socket path: %s", socketPath)
		a.logger.Infof("creaet and listen unix socket")

		listener, err = net.Listen("unix", socketPath)
		if err != nil {
			a.logger.Fatal(err)
		}
	} else {
		a.logger.Infof("bind application to host: %s and port: %s", a.cfg.Listen.BindIP, a.cfg.Listen.Port)
		var err error

		listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", a.cfg.Listen.BindIP, a.cfg.Listen.Port))
		if err != nil {
			a.logger.Fatal(err)
		}
	}

	c := cors.New(cors.Options{
		AllowedMethods:     []string{http.MethodGet, http.MethodPost},
		AllowedOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Authorization", "Location", "Charset", "Access-Control-Allow_Origin", "Content-Type"},
		OptionsPassthrough: true,
		ExposedHeaders:     []string{"Access-Token", "Refresh-Token", "Location", "Authorization", "Content-Disposition"},
		Debug:              false,
	})

	header := c.Handler(a.router)
}
