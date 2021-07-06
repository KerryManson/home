package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"user/handler"
	"user/proto/user"
	"user/model"
)
var consulReg registry.Registry

func main() {


	model.InitRedis()
	consulReg = consul.NewRegistry()


	// Create service
	srv := micro.NewService(
		micro.Address("127.0.0.1:12342" ),
		micro.Name("service.user"),
		micro.Version("latest"),
		micro.Registry(consulReg),
	)

	// Register handler
	user.RegisterSendSmsHandler(srv.Server(), new(handler.User))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
