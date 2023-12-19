package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/obadoraibu/wordle_plus_backend.git/internal/config"
	"github.com/obadoraibu/wordle_plus_backend.git/internal/domain"
	"github.com/obadoraibu/wordle_plus_backend.git/internal/repository"
	"github.com/obadoraibu/wordle_plus_backend.git/internal/service"
	"github.com/obadoraibu/wordle_plus_backend.git/internal/transport/rest"
	"github.com/obadoraibu/wordle_plus_backend.git/internal/transport/rest/handler"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func Run(configPath string) error {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		logrus.Error("cannot create new config")
		return err
	}

	repo, err := repository.NewRepository()
	if err != nil {
		logrus.Fatalf("error creating repository, %s", err)
	}

	service := service.NewService(service.Dependencies{
		Repo: repo,
	})
	handler := handler.NewHandler(handler.Dependencies{
		Service: service,
	})

	server := rest.NewServer()

	go func() {
		if err := server.Start(cfg.Port, handler.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	c := cron.New()
	_, err = c.AddFunc("@midnight", func() {
		req := &domain.NewWordRequest{Length: 5}
		resp, err := service.NewWord(&gin.Context{}, req)
		repo.Mutex.Lock()
		if err != nil {
			repo.DailyWord = "apple"
		} else {
			repo.DailyWord = resp.Word
		}
		repo.Mutex.Unlock()
	})
	if err != nil {
		return err
	}
	c.Start()

	logrus.Printf("service started on port %s", cfg.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("service shutting down")

	if err := server.Stop(context.Background()); err != nil {
		logrus.Errorf("error on server shutting down: %s", err.Error())
	}

	return nil
}
