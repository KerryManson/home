package main

import (
	"getCaptcha/handler"
	pb "getCaptcha/proto"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/consul/v2"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
)
var consulReg registry.Registry

func main() {
	consulReg = consul.NewRegistry()


	// Create service
	srv := micro.NewService(
		micro.Address("127.0.0.1:12341"),
		micro.Name("service.getCaptcha"),
		micro.Version("latest"),
		micro.Registry(consulReg),
	)

	// Register handler
	pb.RegisterGetCaptchaHandler(srv.Server(), new(handler.GetCaptcha))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
