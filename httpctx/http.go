package httpctx

import (
	"aos/pkg/setting"

	"github.com/gin-gonic/gin"
)

// http 请求载体
type HttpCtx struct {
	*gin.Context
	*setting.Logger
}

func New(ctx *gin.Context, logger *setting.Logger) *HttpCtx {
	httpCtx := HttpCtx{}
	httpCtx.Context = ctx
	httpCtx.Logger = logger
	return &httpCtx
}
