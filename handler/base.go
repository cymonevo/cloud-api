package handler

import "github.com/cymonevo/cloud-api/internal/router"

type BaseHandler interface {
	Register() router.Router
}
