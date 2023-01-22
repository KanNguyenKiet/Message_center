package service

import (
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"message-server/user_service/api"
	"message-server/user_service/module/user_logout"
)

func (s *Service) LogoutUser(ctx context.Context, request *api.LogoutRequest) (response *api.LogoutResponse, err error) {
	var isSuccess bool
	err = s.db.Transaction(func(tx *sql.Tx) (err error) {
		isSuccess, err = user_logout.NewLogoutUserProcessor(s.cfg, s.db.WithTx(tx), request).Process(ctx)
		return
	})
	if err != nil {
		return &api.LogoutResponse{
			Code:      int32(codes.Internal),
			Message:   err.Error(),
			IsSuccess: false,
		}, nil
	}

	return &api.LogoutResponse{
		Code:      int32(codes.OK),
		Message:   "Success",
		IsSuccess: isSuccess,
	}, nil
}
