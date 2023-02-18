package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/maskot/internal/handler"
)

func mustInitMux(r *handler.Rpc) (*mux.Router, error) {
	srv := rpc.NewServer()

	srv.RegisterCodec(NewUpCodec(), "application/json")
	srv.RegisterCodec(NewUpCodec(), "application/json;charset=UTF-8")

	err := srv.RegisterService(r, "rpc")

	if err != nil {
		return nil, err
	}

	m := mux.NewRouter()

	m.Handle("/rpc", srv)

	return m, nil
}
