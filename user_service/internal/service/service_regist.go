package service

import (
	"context"
	"database/sql"
	"message-server/user_service/api"
	"message-server/user_service/module/user_register"
)

// RegisterUser Assign API Handler for Server
func (s *Service) RegisterUser(ctx context.Context, request *api.User) (*api.RegisterUserResponse, error) {
	err := s.db.Transaction(func(tx *sql.Tx) (err error) {
		err = user_register.NewUserRegisterProcessor(s.cfg, s.db.WithTx(tx), request).Process(ctx)
		return
	})
	if err != nil {
		return nil, err
	}

	return &api.RegisterUserResponse{
		Code:    200,
		Message: "Register new account successfully!",
	}, nil
}
