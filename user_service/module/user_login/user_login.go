package user_login

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"message-server/user_service/api"
	"message-server/user_service/config"
	"message-server/user_service/extension"
	"message-server/user_service/internal/db"
	"time"
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
	// Check username is existed or not?
	user, err := u.store.GetUserByUsername(ctx, sql.NullString{
		Valid:  true,
		String: u.request.Username,
	})
	if err != nil {
		return false, err
	}

	// Get user's credential
	pwdHashed, err := u.store.GetCrendentailByUserId(ctx, user.ID)
	if err != nil {
		return false, err
	}

	// Check password
	if fmt.Sprintf("%x", sha256.Sum256([]byte(u.request.Password))) != pwdHashed.String {
		return false, nil
	}

	// Generate session key for current user
	// If user still in valid session, do nothing
	if !user.SessionKey.Valid {
		sessionExpiredAt := time.Now().Add(time.Hour)
		sessionKey := extension.NewSession(user.ID, user.UserName.String, user.Email, sessionExpiredAt).GenKey()
		_, err = u.store.UpdateSessionInfo(ctx, db.UpdateSessionInfoParams{
			SessionKey: sql.NullString{
				String: sessionKey,
				Valid:  true,
			},
			SessionExpired: sql.NullTime{
				Time:  sessionExpiredAt,
				Valid: true,
			},
			ID: user.ID,
		})
		if err != nil {
			return false, fmt.Errorf("%s", "Session init failed")
		}
	}

	return true, nil
}
