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
	"github.com/go-mumu/cs-go/container/provider/dao_provider"
	"github.com/go-mumu/cs-go/proto/pb"
	"google.golang.org/grpc/metadata"
	"strconv"
)

type UserServiceHandler struct {
	pb.UnimplementedUserServiceServer
	Dao *dao_provider.Dao
}

func NewUserServiceHandler(Dao *dao_provider.Dao) *UserServiceHandler {
	return &UserServiceHandler{
		Dao: Dao,
	}
}

func (h *UserServiceHandler) IsVip(ctx context.Context, req *pb.IsVipReq) (*pb.IsVipRes, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	mid, err := strconv.ParseInt(md.Get("mid")[0], 10, 64)
	if err != nil {
		return nil, err
	}

	if ok {
		userInfo := h.Dao.WxuserDao.GetUserInfoByMid(ctx, mid)

		return &pb.IsVipRes{
			Vip7:        userInfo.Vip7,
			Overdue:     0,
			Type:        "",
			Viptime:     "2022",
			Vipvalidity: "",
		}, nil
	} else {
		return &pb.IsVipRes{}, nil
	}
}