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
	"github.com/go-mumu/cs-go/service/dal/third/central/interest"
	"strconv"
	"time"
)

type UserServiceHandler struct {
	pb.UnimplementedUserServiceServer
	interest.Interest
	WxuserDao *dao.WxuserDao
}

func NewUserServiceHandler(wxuserDao *dao.WxuserDao) *UserServiceHandler {
	return &UserServiceHandler{WxuserDao: wxuserDao}
}

func (h *UserServiceHandler) IsVip(ctx context.Context, req *pb.IsVipReq) (*pb.IsVipRes, error) {

	userInfo := h.WxuserDao.GetUserInfoByMid(ctx, req.Mid)

	userInterest := h.UserInterest(ctx, strconv.FormatInt(req.Mid, 10))

	var overdue int32 = 0
	if parseTime, _ := time.ParseInLocation(time.DateTime, userInterest["interest_end"]+" 23:59:59", time.Local); parseTime.After(time.Now()) {
		overdue = 1
	}

	return &pb.IsVipRes{
		Vip7:        userInfo.Vip7,
		Overdue:     overdue,
		Type:        "",
		Viptime:     userInfo.Createtime.String(),
		Vipvalidity: userInterest["interest_end"] + " 23:59:59",
	}, nil
}
