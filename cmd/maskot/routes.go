package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/maskot/internal/handler"
	"github.com/maskot/internal/handler/middleware"
	"go.uber.org/zap"
)

func mustInitMux(r *handler.RPC, log *zap.Logger) (*mux.Router, error) {
	srv := rpc.NewServer()

	srv.RegisterCodec(NewUpCodec(), "application/json")
	srv.RegisterCodec(NewUpCodec(), "application/json;charset=UTF-8")

	err := srv.RegisterService(r, "rpc")

	if err != nil {
		return nil, err
	}

	m := mux.NewRouter()

	logMiddleware := middleware.NewLogger(log)

	m.Handle("/rpc", logMiddleware.Log(srv))

	return m, nil
}
