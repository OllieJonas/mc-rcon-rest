package main

import (
	"flag"
	"net/http"
	"rest-rcon/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	port   string
	mode   string
	ttlStr string

	ttl time.Duration

	PossibleModes = map[string]bool{
		"debug":   true,
		"test":    true,
		"release": true,
	}

	Service *service.DispatchService
)

func init() {
	flag.StringVar(&port, "port", "8080", "Port to listen on")
	flag.StringVar(&mode, "mode", "release", "either debug, release or test")
	flag.StringVar(&ttlStr, "ttl", "2m", "Time to live. See https://pkg.go.dev/time#ParseDuration for format")
	flag.Parse()

	if !PossibleModes[mode] {
		logrus.Fatal("Mode must be either debug, release or test!")
	}

	var err error
	ttl, err = time.ParseDuration(ttlStr)

	if err != nil {
		logrus.Fatal("Malformed ttl! See https://pkg.go.dev/time#ParseDuration for more format!", err)
	}

	Service = service.NewDispatchService(ttl)
}

func main() {
	logrus.Info("Starting RCON REST Server on port " + port)
	gin.SetMode(mode)
	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.IndentedJSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Unknown route " + c.Request.URL.Path})
	})

	// health check
	router.GET("/health", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.POST("/dispatch", onDispatch)

	err := router.Run(":" + port)

	if err != nil {
		logrus.Fatal("Unable to listen and serve", err)
	}
}

func onDispatch(c *gin.Context) {
	var request service.DispatchRequest

	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Malformed JSON!"})
	} else {
		c.IndentedJSON(http.StatusOK, Service.DispatchCommands(&request))
	}
}
