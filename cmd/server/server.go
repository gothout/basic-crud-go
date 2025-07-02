package server

import (
	adminHandler "basic-crud-go/internal/app/admin/handler"
	env "basic-crud-go/internal/configuration/env/enviroument"
	envServer "basic-crud-go/internal/configuration/env/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	router := gin.Default()
	if env.GetEnvironment() == "DEV" {
		gin.SetMode(gin.DebugMode)
	}
	if env.GetEnvironment() == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	// future -- CORS

	// register handlers
	adminHandler.RegisterAdminRoutes(router)
	return router

}

// StartServer inicialize server HTTP or HTTPs
func StartServer(router *gin.Engine) {
	// Start http
	address_http := fmt.Sprintf("%s:%s", envServer.GetListenServer(), envServer.GetHTTPPort())
	go func() {
		if err := router.Run(address_http); err != nil {
			log.Fatalf("error inicialize server http: %v", err)
		}
		log.Printf("Server listen on http://%s", address_http)
	}()
	// Start https
	address_https := fmt.Sprintf("%s:%s", envServer.GetListenServer(), envServer.GetHTTPSPort())
	if envServer.GetHTTPSuse() {
		go func() {
			httpsAddress := address_https
			if err := http.ListenAndServeTLS(httpsAddress, "./certificates/cert.crt", "./certificates/privkey.key", router); err != nil {
				log.Fatalf("error inicialize server https: %v", err)
			}
		}()
	}
	select {}
}
