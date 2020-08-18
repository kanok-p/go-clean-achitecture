package main

import (
	"github.com/kanok-p/go-clean-achitecture/app"
	"github.com/kanok-p/go-clean-achitecture/config"
	userStore "github.com/kanok-p/go-clean-achitecture/repository/users/store"
	userService "github.com/kanok-p/go-clean-achitecture/service/users"
)

func newApp(config *config.Config) *app.App {
	usrRepo := userStore.New(config.MongoDBEndpoint, config.MongoDBName, config.MongoDBCollUser)
	usrService := userService.New(usrRepo)
	return app.New(usrService)
}
