package service

import (
	"context"
	"database/sql"
	"log"
	"message-server/user_service/api"
	"message-server/user_service/module/user_login"

	"google.golang.org/grpc/codes"
)

func (s *Service) LoginUser(ctx context.Context, request *api.LoginRequest) (*api.LoginResponse, error) {
	var isSuccess bool
	var sessionKey string
	_ = s.db.Transaction(func(tx *sql.Tx) (err error) {
		isSuccess, sessionKey, err = user_login.NewUserLoginProcessor(s.cfg, s.db.WithTx(tx), request).Process(ctx)
		if err != nil {
			log.Println(err)
		}
		return
	})
	if isSuccess {
		return &api.LoginResponse{
			Code:       int32(codes.OK),
			Message:    "Login successfully!",
			IsSuccess:  isSuccess,
			SessionKey: sessionKey,
		}, nil
	} else {
		return &api.LoginResponse{
			Code:      int32(codes.OK),
			Message:   "Login failed!",
			IsSuccess: isSuccess,
		}, nil
	}
}
