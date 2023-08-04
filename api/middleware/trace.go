/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-03
 * File: trace.go
 * Desc: trace middleware
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-mumu/cs-go/library/common/consts"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := metadata.AppendToOutgoingContext(c.Request.Context(), consts.TraceId, uuid.New().String())

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
