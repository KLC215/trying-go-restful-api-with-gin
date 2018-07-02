package main

import (
	"apiserver/config"
	"apiserver/router"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"

	"github.com/spf13/pflag"

	"github.com/gin-gonic/gin"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	// Initial config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// Create Gin engine
	g := gin.New()

	// Set gin run mode: debug, release, test
	gin.SetMode(viper.GetString("runmode"))

	// Define Gin middlewares
	middlewares := []gin.HandlerFunc{}

	// Define routes
	router.Load(
		// Cores
		g,
		// Middlewares
		middlewares...,
	)

	// Ping server to make sure the routing is working
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming request on http address: %s", viper.GetString("port"))
	log.Printf(http.ListenAndServe(viper.GetString("port"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")

		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router.")
}
