package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/kanok-p/go-clean-architecture/config"
)

const (
	port = ":8080"
)

func main() {
	conf, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()
	router.Use(gin.Recovery())

	_ = newApp(conf).Register(router)

	_ = router.Run(port)
}
