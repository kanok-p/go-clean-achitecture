package users

import (
	"github.com/kanok-p/go-clean-architecture/service/users"
)

type Controller struct {
	service users.Service
}

func New(service users.Service) (ctrl *Controller) {
	return &Controller{service: service}
}
