// +build wireinject

package api

import (
	"Go-000/Week04/service"
	"github.com/google/wire"
)

func InitializeServer() (*Server, func(), error) {
	wire.Build(NewServer, service.Provider)
	return nil, nil, nil
}