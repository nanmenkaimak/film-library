package applicator

import (
	"context"
	"github.com/nanmenkaimak/film_library/internal/config"
	"github.com/nanmenkaimak/film_library/internal/controller/http"
	"github.com/nanmenkaimak/film_library/internal/database/postgres"
	"github.com/nanmenkaimak/film_library/internal/repository"
	"github.com/nanmenkaimak/film_library/internal/usecase"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	logger *zap.SugaredLogger
	config *config.Config
}

func NewApp(logger *zap.SugaredLogger, config *config.Config) *App {
	return &App{
		logger: logger,
		config: config,
	}
}

func (a *App) Run() {
	var cfg = a.config
	var l = a.logger

	ctx, cancel := context.WithCancel(context.TODO())
	_ = ctx

	mainDB, err := postgres.New(cfg.Database.Main)
	if err != nil {
		l.Fatalf("cannot сonnect to mainDB '%s:%d': %v", cfg.Database.Main.Host, cfg.Database.Main.Port, err)
	}

	replicaDB, err := postgres.New(cfg.Database.Replica)
	if err != nil {
		l.Fatalf("cannot сonnect to replicaDB '%s:%d': %v", cfg.Database.Replica.Host, cfg.Database.Replica.Port, err)
	}

	repo := repository.NewRepository(mainDB, replicaDB)

	service := usecase.NewService(repo, cfg.Auth, l)

	endpointHandler := http.NewEndpointHandler(service, l)

	router := http.NewRouter(l)

	httpCfg := cfg.HttpServer

	server, err := http.NewServer(httpCfg.Port, httpCfg.ShutdownTimeout, router, l, endpointHandler)
	if err != nil {
		l.Fatalf("failed to create server err: %v", err)
	}

	server.Run()
	defer func() {
		if err := server.Stop(); err != nil {
			l.Panicf("failed close server err: %v", err)
		}
		l.Info("server closed")
	}()

	a.gracefulShutdown(cancel)
}

func (a *App) gracefulShutdown(cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch
	signal.Stop(ch)
	cancel()
}
