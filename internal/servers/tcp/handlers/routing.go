package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"md5er/internal/md5"
	"net/http"
)

var ErrValidation = errors.New("validation error")

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

type Response struct {
	Status string        `json:"status"`
	Data   *ResponseData `json:"data,omitempty"`
}

type ResponseData struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Content any    `json:"content,omitempty"`
}

func (r *Routing) successResponse(ctx *gin.Context, status int, data ...any) {
	res := Response{Status: "ok"}

	if len(data) > 0 {
		res.Data = &ResponseData{Content: data[0]}
	}

	ctx.JSON(status, res)
}

func (r *Routing) errorResponse(ctx *gin.Context, status int, errorCode string, err error, errorsData ...any) {

	res := Response{
		Status: "error",
		Data: &ResponseData{
			Code:    errorCode,
			Message: err.Error(),
		},
	}

	if len(errorsData) > 0 {
		res.Data.Content = errorsData[0]
	}

	ctx.JSON(status, res)
}

func (r *Routing) validationErrorResponse(ctx *gin.Context, errs validate.Errors) {
	r.errorResponse(ctx, http.StatusBadRequest, "validation", ErrValidation, errs)
}

func (r *Routing) validate(ctx *gin.Context, req any) bool {
	v := validate.Struct(req)
	if !v.Validate() {
		r.validationErrorResponse(ctx, v.Errors)
		return false
	}

	return true
}
