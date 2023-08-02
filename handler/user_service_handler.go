/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-26
 * File: user_service.go
 * Desc: user service handler
 */

package handler

import (
	"context"
	"github.com/go-mumu/cs-go/container/provider"
	"github.com/go-mumu/cs-go/proto/pb"
)

type UserServiceHandler struct {
	pb.UnimplementedUserServiceServer
	Dao *provider.Dao
}

func (h *UserServiceHandler) IsVip(ctx context.Context, req *pb.IsVipReq) (*pb.IsVipRes, error) {

	userInfo := h.Dao.WxuserDao.GetUserInfoByMid(ctx, req.Mid)

	return &pb.IsVipRes{
		Vip7:        userInfo.Vip7,
		Overdue:     0,
		Type:        "",
		Viptime:     "2022",
		Vipvalidity: "",
	}, nil
}
