package server

import (
	enterpriseHandler "basic-crud-go/internal/app/admin/enterprise/handler"
	userHandler "basic-crud-go/internal/app/admin/user/handler"
	env "basic-crud-go/internal/configuration/env/environment"
	envServer "basic-crud-go/internal/configuration/env/server"
	"fmt"
	"github.com/gin-contrib/cors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	router := gin.Default()
	devCors, prodCors := envServer.GetCorsOrigins()
	if env.GetEnvironment() == "DEV" {
		gin.SetMode(gin.DebugMode)
		// Cors set
		config := cors.Config{
			AllowOrigins: []string{
				fmt.Sprintf("http://%s", devCors),
				fmt.Sprintf("https://%s", devCors),
			},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			AllowCredentials: true,
		}
		router.Use(cors.New(config))
	}
	if env.GetEnvironment() == "PROD" {
		gin.SetMode(gin.ReleaseMode)
		// Cors set
		config := cors.Config{
			AllowOrigins:     []string{fmt.Sprintf("https://%s", prodCors)},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			AllowCredentials: true,
		}
		router.Use(cors.New(config))
	}

	// register handlers
	userHandler.RegisterUserRoutes(router)
	enterpriseHandler.RegisterEnterpriseRoutes(router)
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
