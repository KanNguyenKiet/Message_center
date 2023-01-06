package service

import (
	"context"
	"database/sql"
	"message-server/user_service/api"
	"message-server/user_service/module/user_login"
)

func (s *Service) LoginUser(ctx context.Context, request *api.LoginRequest) (*api.LoginResponse, error) {
	var isSuccess bool
	err := s.db.Transaction(func(tx *sql.Tx) (err error) {
		isSuccess, err = user_login.NewUserLoginProcessor(s.cfg, s.db.WithTx(tx), request).Process(ctx)
		return
	})
	if err != nil {
		return nil, err
	}
	if isSuccess {
		return &api.LoginResponse{
			Code:      200,
			Message:   "Login successfully!",
			IsSuccess: isSuccess,
		}, nil
	} else {
		return &api.LoginResponse{
			Code:      200,
			Message:   "Login failed!",
			IsSuccess: isSuccess,
		}, nil
	}
}
