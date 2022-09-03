package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"vyatta-cfg-go/internal/controller"
	"vyatta-cfg-go/internal/service/logic/activecfg"
)

var tree *activecfg.Node

func init() {
	tree = &activecfg.Node{}
}

func main() {

	s := g.Server()

	s.Use(ghttp.MiddlewareHandlerResponse, ghttp.MiddlewareCORS)
	//对象注册
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Bind(controller.Interfaces{})
		group.Bind(controller.System{})
		group.Bind(controller.Service{})
		group.Bind(controller.Protocols{})
		group.Bind(controller.Show{})
	})
	s.Run()
}
