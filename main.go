package main

import (
	"github.com/gin-gonic/gin"

	"github.com/kanok-p/go-clean-achitecture/config"
)

const (
	port = ":8080"
)

func main() {
	conf := config.Get()

	router := gin.New()
	router.Use(gin.Recovery())

	_ = newApp(conf).Register(router)

	_ = router.Run(port)
}
