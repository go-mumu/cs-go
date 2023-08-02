/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-27
 * File: user.go
 * Desc:
 */

package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-mumu/cs-go/api/client"
	"github.com/go-mumu/cs-go/proto/pb"
	"net/http"
	"time"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) IsVip(c *gin.Context) {
	mid := c.GetInt64("mid")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5000*time.Millisecond)
	defer cancel()

	cli := new(client.Clients)

	res, _ := cli.UserClient(ctx).IsVip(ctx, &pb.IsVipReq{Mid: mid})
	/*var isVipReq pb.IsVipReq

	_ = c.ShouldBind(&isVipReq)

	ctx, cancel := context.WithTimeout(context.Background(), 5000 * time.Millisecond)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("mid", "1598606668982307"))

	cli := new(client.Clients)

	res, _:= cli.UserClient(ctx).IsVip(ctx, &pb.IsVipReq{Mid: isVipReq.Mid})*/

	c.JSON(http.StatusOK, gin.H{
		"res": res,
	})
}
