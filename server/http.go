package server

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	docs "github.com/hamdiBouhani/GopherNet-golang/docs" //Swagger definition
	"github.com/hamdiBouhani/GopherNet-golang/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GopherNet API
// @version 1.0
// @description GopherNet API

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// HttpService provides a HTTP Rest endpoint for the loading functions
type HttpService struct {
	router *gin.Engine

	ns         *gin.RouterGroup
	port       string
	corsConfig cors.Config

	BurrowServiceInstance *services.BurrowService

	devMode  bool
	testMode bool
}

func NewHttpService(burrowServiceInstance *services.BurrowService) *HttpService {
	svc := &HttpService{
		BurrowServiceInstance: burrowServiceInstance,
	}

	return svc
}

// Start the http service with given listeners and then listen on port
func (svc *HttpService) Start() error {

	var port = flag.String("port", "8080", "Http port to serve application on")
	flag.Parse()

	svc.port = fmt.Sprintf(":%s", *port)

	svc.corsConfig = cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	svc.devMode = os.Getenv("RUN_MODE") == "DEV"
	svc.testMode = os.Getenv("TEST_HTTP") == "TRUE"
	if svc.devMode {
		log.Println("-DEV_MODE enabled-")
	}

	err := svc.registerRoutes()
	if err != nil {
		return err
	}
	if !svc.testMode {
		return svc.router.Run(svc.port) // listen & serve
	}
	return nil
}

func (svc *HttpService) registerRoutes() error {
	if svc.testMode {
		gin.SetMode(gin.TestMode)
	}
	docs.SwaggerInfo.BasePath = "/api/v1"
	svc.router = gin.Default()
	svc.router.Use(cors.New(svc.corsConfig))

	svc.ns = svc.router.Group("/api/v1")

	svc.Routes()
	return nil
}

func (svc *HttpService) Routes() {
	svc.ns.GET("/ping", svc.ping) //PING/PONG
	svc.ns.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	svc.ns.GET("/burrows", svc.BurrowStatus)
	svc.ns.PUT("/rent-burrow/:id", svc.RentBurrow)
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Router /ping [get]
func (svc *HttpService) ping(c *gin.Context) {
	c.JSON(200, "pong")
}
