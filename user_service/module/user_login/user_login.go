package user_login

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"message-server/user_service/api"
	"message-server/user_service/config"
	"message-server/user_service/internal/db"
)

type UserLoginProcessor struct {
	cfg     *config.ServerConfig
	store   db.StoreQuerier
	request *api.LoginRequest
}

func NewUserLoginProcessor(cfg *config.ServerConfig, store db.StoreQuerier, request *api.LoginRequest) *UserLoginProcessor {
	return &UserLoginProcessor{
		cfg:     cfg,
		store:   store,
		request: request,
	}
}

func (u *UserLoginProcessor) Process(ctx context.Context) (bool, error) {
	// Check username is exist or not?
	user_id, err := u.store.GetUserByUsername(ctx, sql.NullString{
		Valid:  true,
		String: u.request.Username,
	})
	if err != nil {
		return false, err
	}

	// Get user's credential
	pwd_hashed, err := u.store.GetCrendentailByUserId(ctx, user_id)
	if err != nil {
		return false, err
	}

	// Check password
	if fmt.Sprint(sha256.Sum256([]byte(u.request.Password))) != pwd_hashed.String {
		return false, nil
	}

	return true, nil
}
