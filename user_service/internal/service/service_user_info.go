package service

import (
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"message-server/user_service/api"
	"message-server/user_service/module/user_info"
)

func (s *Service) GetUserInfo(ctx context.Context, request *api.GetUserInfoRequest) (response *api.GetUserInfoResponse, err error) {
	var userInfo *api.GetUserInfoResponse_UserInfo
	err = s.db.Transaction(func(tx *sql.Tx) (err error) {
		userInfo, err = user_info.NewGetUserInfoProcessor(s.cfg, s.db.WithTx(tx), request).Process(ctx)
		return
	})
	if err != nil {
		return nil, err
	}
	return &api.GetUserInfoResponse{
		Code:    int32(codes.OK),
		Message: "Success",
		Data:    userInfo,
	}, nil
}
