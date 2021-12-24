package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	address string
	port string
	mode string

	PossibleModes = map[string]bool{
		"debug": true,
		"test": true,
		"release": true,
	}
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
	router.GET("/health", healthCheck)

	err := router.Run(address + ":" + port)

	if err != nil {
		logrus.Fatalf("Unable to listen and serve ", err)
	}
}


func healthCheck(ctx *gin.Context) {

	response := struct {
		Status string `json:"status"`
	}{}

	response.Status = "OK"

	ctx.IndentedJSON(http.StatusOK, response)
}
