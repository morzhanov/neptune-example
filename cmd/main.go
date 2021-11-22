package main

import (
	"log"

	"github.com/morzhanov/neptune-example/internal/config"
	"github.com/morzhanov/neptune-example/internal/db"
	"github.com/morzhanov/neptune-example/internal/logger"
	"github.com/morzhanov/neptune-example/internal/service"
	"go.uber.org/zap"
)

func failOnError(l *zap.Logger, step string, err error) {
	if err != nil {
		l.Fatal("initialization error", zap.Error(err), zap.String("step", step))
	}
}

func main() {
	l, err := logger.NewLogger()
	if err != nil {
		log.Fatal("initialization error during logger setup")
	}
	c, err := config.NewConfig()
	failOnError(l, "config", err)
	d, err := db.NewDB(c)
	failOnError(l, "neptune database", err)

	srv := service.NewService(d, l)
	if err := srv.Run(); err != nil {
		failOnError(l, "service", err)
	}
}
