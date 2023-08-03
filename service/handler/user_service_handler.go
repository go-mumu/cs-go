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
	"github.com/go-mumu/cs-go/proto/pb"
	"github.com/go-mumu/cs-go/service/dal/dao"
)

type UserServiceHandler struct {
	pb.UnimplementedUserServiceServer
	WxuserDao *dao.WxuserDao
}

func NewUserServiceHandler(
	wxuserDao *dao.WxuserDao,
) *UserServiceHandler {
	return &UserServiceHandler{
		WxuserDao: wxuserDao,
	}
}

func (h *UserServiceHandler) IsVip(ctx context.Context, req *pb.IsVipReq) (*pb.IsVipRes, error) {

	userInfo := h.WxuserDao.GetUserInfoByMid(ctx, req.Mid)

	return &pb.IsVipRes{
		Vip7:        userInfo.Vip7,
		Overdue:     0,
		Type:        "",
		Viptime:     "2022",
		Vipvalidity: "",
	}, nil
}
