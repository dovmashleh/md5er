package handlers

import "github.com/gin-gonic/gin"

func (r *Routing) health(ctx *gin.Context) {
	ctx.JSON(200, struct {
		Health string `json:"health"`
	}{
		"ok",
	})
}
