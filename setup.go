package main

import (
	"github.com/kanok-p/go-clean-achitecture/app"
	"github.com/kanok-p/go-clean-achitecture/config"
	userStore "github.com/kanok-p/go-clean-achitecture/repository/users/store"
	userService "github.com/kanok-p/go-clean-achitecture/service/users"
	"github.com/kanok-p/go-clean-achitecture/util/validate"
)

var formats = map[string][]string{
	"password": {`[0-9]`, `[a-z]`, `[A-Z]`, `^[0-9a-zA-Z]+$`},
}

func newApp(config *config.Config) *app.App {
	usrRepo := userStore.New(config)
	validator := validate.New(usrRepo, config, formats)
	usrService := userService.New(usrRepo, validator)
	return app.New(usrService)
}
