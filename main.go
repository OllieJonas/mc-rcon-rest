package main

import (
	"flag"
	"net/http"

	"rest-rcon/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	address string
	port    string
	mode    string

	PossibleModes = map[string]bool{
		"debug":   true,
		"test":    true,
		"release": true,
	}

	Service = service.NewDispatchService()
)

func init() {
	flag.StringVar(&address, "address", "localhost", "Address to listen on")
	flag.StringVar(&port, "port", "8080", "Port to listen on")
	flag.StringVar(&mode, "mode", "release", "either debug, release or test")
	flag.Parse()

	if !PossibleModes[mode] {
		logrus.Fatal("Mode must be either debug, release or test!")
	}
}

func main() {
	logrus.Info("Starting RCON REST Server on port " + port)
	gin.SetMode(mode)
	router := gin.Default()

	// paths

	// health check
	router.GET("/health", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.POST("/dispatch", onDispatch)

	err := router.Run(":" + port)

	if err != nil {
		logrus.Fatalf("Unable to listen and serve ", err)
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
