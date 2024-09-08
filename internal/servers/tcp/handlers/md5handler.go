package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"unsafe"
)

type MD5PostRequest struct {
	DataString string `json:"text" validate:"required|string"`
	DataBase64 []byte `json:"base64" validate:"sometimes"`
}
type MD5Response struct {
	Hash   string `json:"hash"`
	Base64 []byte `json:"base64"`
}

func (r *Routing) postMD5(ctx *gin.Context) {
	req := &MD5PostRequest{}
	if err := ctx.ShouldBindBodyWithJSON(req); err != nil {
		r.errorResponse(ctx, 400, "json", err)
		return
	}
	if !r.validate(ctx, req) {
		return
	}
	dataSlice := unsafe.Slice(unsafe.StringData(req.DataString), len(req.DataString))
	hash := r.Md5service.AsByteArray(dataSlice)

	resp := &MD5Response{
		Hash:   fmt.Sprintf("%x", hash),
		Base64: hash[:],
	}
	ctx.JSON(200, resp)
}
func (r *Routing) getMD5(ctx *gin.Context) {

}
