package main

import (
	"os"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"
	"github.com/taadis/cms-api/internal/posts/handler"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.cms"),
		micro.Registry(etcd.NewRegistry(registry.Addrs(os.Getenv("MICRO_REGISTRY_ADDRESS")))),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	micro.RegisterHandler(service.Server(), handler.NewPostsHandler())

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
