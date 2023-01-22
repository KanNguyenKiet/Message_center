package user_info

import (
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc/metadata"
	"message-server/user_service/api"
	"message-server/user_service/config"
	"message-server/user_service/internal/db"
	"time"
)

type GetUserInfoProcessor struct {
	cfg     *config.ServerConfig
	store   db.StoreQuerier
	request *api.GetUserInfoRequest
}

func NewGetUserInfoProcessor(cfg *config.ServerConfig, store db.StoreQuerier, request *api.GetUserInfoRequest) *GetUserInfoProcessor {
	return &GetUserInfoProcessor{
		cfg:     cfg,
		store:   store,
		request: request,
	}
}

func (g *GetUserInfoProcessor) Process(ctx context.Context) (*api.GetUserInfoResponse_UserInfo, error) {
	var sessionKey string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		sessionKey = md.Get("_session")[0]
	}
	userInfo, err := g.store.GetUserInfoBySessionKey(ctx, sql.NullString{
		String: sessionKey,
		Valid:  true,
	})
	if err != nil {
		return nil, err
	}
	if !userInfo.SessionExpired.Valid {
		return nil, errors.New("user is not logged in")
	}
	if userInfo.SessionExpired.Time.Before(time.Now()) {
		return nil, errors.New("session expired")
	}

	return &api.GetUserInfoResponse_UserInfo{
		LastName:  userInfo.LastName.String,
		FirstName: userInfo.FirstName.String,
		Phone:     userInfo.Phone.String,
		Email:     userInfo.Email,
	}, nil
}
