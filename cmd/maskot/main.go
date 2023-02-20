package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/maskot/internal/handler"
	"github.com/maskot/internal/infra/postgres"
	pgrepos "github.com/maskot/internal/repo/postgres"
	gb "github.com/maskot/internal/use_cases/seamless/get_balance"
	rb "github.com/maskot/internal/use_cases/seamless/rollback_transaction"
	wad "github.com/maskot/internal/use_cases/seamless/withdraw_and_deposit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initLogger(lvl string) *zap.Logger {
	logLvls := map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"error": zapcore.ErrorLevel,
	}

	c := zap.NewProductionConfig()
	c.Development = true
	c.EncoderConfig.EncodeTime = zapcore.EpochTimeEncoder

	logLvl, ok := logLvls[lvl]
	if !ok {
		panic("newLog: invalid lvl given")
	}
	c.Level = zap.NewAtomicLevelAt(logLvl)

	log, err := c.Build()
	if err != nil {
		panic(err)
	}

	return log
}

func main() {
	log := initLogger(os.Getenv("LOG_LEVEL"))

	db, closeDB, err := postgres.NewPool(&postgres.Config{
		DBName:   os.Getenv("DB_DATABASE"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  "disable",
	})
	if err != nil {
		log.Panic("failed to create db pool", zap.Error(err))
	}
	defer func() {
		if err := closeDB(); err != nil {
			log.Panic("failed to close db pool", zap.Error(err))
		}
	}()

	repos := pgrepos.NewRepos(db)

	wadUseCase := wad.NewUseCase(repos.Balance, repos.Transaction)
	rtUseCase := rb.NewUseCase(repos.Balance, repos.Transaction)
	gbUseCase := gb.NewUseCase(repos.Balance)

	rpcHandler := handler.NewRPC(&handler.RPCConfig{
		GbUseCase:  gbUseCase,
		WadUseCase: wadUseCase,
		RtUseCase:  rtUseCase,
	})
	mux, err := mustInitMux(rpcHandler, log)
	if err != nil {
		log.Panic("failed to init mux", zap.Error(err))
	}

	srv := http.Server{Addr: ":8080", Handler: mux}
	go func() {
		log.Info("server is starting on port: 8080")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panic("failed to ListenAndServe :8080", zap.Error(err))
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh

	shutdownServer(&srv, logWithServer(log, "server"))
}

func logWithServer(log *zap.Logger, srv string) *zap.Logger {
	return log.With(zap.String("server", srv))
}

func shutdownServer(srv *http.Server, log *zap.Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to shutdown server", zap.Error(err))
		return
	}

	log.Info("server is shut down")
}
