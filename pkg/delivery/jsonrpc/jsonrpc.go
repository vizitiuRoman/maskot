package jsonrpc

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/maskot/pkg/delivery/jsonrpc/api"
	"github.com/maskot/pkg/use_cases"
)

type Dependencies struct {
	UseCases *use_cases.UseCases
}

func NewJSONRpc(deps *Dependencies) (*mux.Router, error) {
	srv := rpc.NewServer()
	srv.RegisterCodec(json2.NewCodec(), "application/json")
	srv.RegisterCodec(json2.NewCodec(), "application/json;charset=UTF-8")

	err := srv.RegisterService(api.NewRpc(deps.UseCases.Seamless), "rpc")

	if err != nil {
		return nil, err
	}

	r := mux.NewRouter()

	r.Handle("/rpc", srv)

	return r, nil
}
