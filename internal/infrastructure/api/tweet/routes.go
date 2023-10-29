package tweetapi

import "github.com/gin-gonic/gin"

type Router struct {
	routeGroup gin.RouterGroup
	api        TweetAPI
}

func NewRouter(routeGroup gin.RouterGroup, api TweetAPI) *Router {
	return &Router{
		routeGroup: routeGroup,
		api:        api,
	}
}

func (r *Router) Register() {
	r.routeGroup.POST("/", r.api.CreateTweet)
}
