//		Simple Admin
//
//		This is simple admin api doc
//
//		Schemes: http, https
//		Host: localhost:9100
//		BasePath: /
//		Version: 1.2.0
//		Contact:
//		SecurityDefinitions:
//		  Token:
//		    type: apiKey
//		    name: Authorization
//		    in: header
//		Security:
//		  - Token: []
//	    Consumes:
//		  - application/json
//
//		Produces:
//		  - application/json
//
// swagger:meta
package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/config"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/handler"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/k8s_terminal"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/core.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	server := rest.MustNewServer(c.RestConf, rest.WithCors(c.CROSConf.Address))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	go func() {
		http.HandleFunc("/ws", k8s_terminal.Terminal.WsHandler)
		http.ListenAndServe(":8081", nil)
	}()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
