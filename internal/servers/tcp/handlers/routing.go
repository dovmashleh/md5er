package handlers

import (
	"github.com/gin-gonic/gin"
	"md5er/internal/md5"
)

type Routing struct {
	Md5service *md5.MD5service
	Router     *gin.Engine
}

func NewRouting() *Routing {
	routing := &Routing{
		Md5service: md5.New(),
	}
	routing.makeRouter()
	return routing
}

func (r *Routing) makeRouter() *gin.Engine {
	router := gin.New()
	r.Router = router
	api := router.Group("/api/v1")
	api.POST("/md5", r.postMD5)
	api.GET("/md5", r.getMD5)
	api.GET("/health", r.health)
	return router
}
