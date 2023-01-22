package user_logout

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"message-server/user_service/api"
	"message-server/user_service/config"
	"message-server/user_service/internal/db"
)

type LogoutUserProcessor struct {
	cfg     *config.ServerConfig
	store   db.StoreQuerier
	request *api.LogoutRequest
}

func NewLogoutUserProcessor(cfg *config.ServerConfig, store db.StoreQuerier, request *api.LogoutRequest) *LogoutUserProcessor {
	return &LogoutUserProcessor{
		cfg:     cfg,
		store:   store,
		request: request,
	}
}

func (l *LogoutUserProcessor) Process(ctx context.Context) (bool, error) {
	userSession, err := l.store.GetUserSession(ctx, sql.NullString{
		String: l.request.Username,
		Valid:  true,
	})
	if err != nil {
		log.Println(err)
		return false, errors.New("fail to get user's session")
	}
	if !userSession.SessionKey.Valid {
		return false, errors.New("user is not logged in")
	}

	_, err = l.store.ClearUserSession(ctx, userSession.ID)
	if err != nil {
		return false, errors.New("logout failed")
	}

	return true, nil
}
