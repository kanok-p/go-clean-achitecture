package main

import (
	"github.com/kanok-p/go-clean-architecture/app"
	userController "github.com/kanok-p/go-clean-architecture/app/users"
	"github.com/kanok-p/go-clean-architecture/config"
	userStore "github.com/kanok-p/go-clean-architecture/repository/users/store"
	userService "github.com/kanok-p/go-clean-architecture/service/users"
	"github.com/kanok-p/go-clean-architecture/util/validate"
)

var formats = map[string][]string{
	"password": {`[0-9]`, `[a-z]`, `[A-Z]`, `^[0-9a-zA-Z]+$`},
}

func newApp(config *config.Config) *app.App {
	usrRepo, err := userStore.New(config)
	if err != nil {
		panic(err)
	}

	validator := validate.New(usrRepo, config, formats)
	usrService := userService.New(usrRepo, validator)
	usrController := userController.New(usrService)
	return app.New(*usrController)
}
