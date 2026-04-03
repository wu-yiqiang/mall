// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"mall/service/shortener/api/internal/config"
	"mall/service/shortener/api/internal/handler"
	"mall/service/shortener/api/internal/svc"
)

var configFile = flag.String("f", "etc/shortener-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
