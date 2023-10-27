package auth

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	routeGroup gin.RouterGroup
	api        Auth
}

func NewRouter(routeGroup gin.RouterGroup, api Auth) *Router {
	return &Router{
		routeGroup: routeGroup,
		api:        api,
	}
}

func (r *Router) Register() {
	r.routeGroup.POST("/register", r.api.CreateUser)
}
