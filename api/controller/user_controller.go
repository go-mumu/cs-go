/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-27
 * File: user.go
 * Desc: user controller
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-mumu/cs-go/api/client"
	"github.com/go-mumu/cs-go/library/config"
	"github.com/go-mumu/cs-go/library/log"
	"github.com/go-mumu/cs-go/proto/pb"
	"net/http"
)

type UserController struct {
	*client.ServiceClient
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) IsVip(c *gin.Context) {
	mid := c.GetInt64("mid")

	log.Cli.Info("remote config", "msg", config.V.GetString("interest.msg"))

	ctx := c.Request.Context()

	res, _ := u.UserClient().IsVip(ctx, &pb.IsVipReq{Mid: mid})

	c.JSON(http.StatusOK, gin.H{
		"res": res,
	})
}
