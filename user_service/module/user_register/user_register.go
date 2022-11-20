package user_register

import (
	"context"
	"database/sql"
	"log"
	"message-server/user_service/api"
	"message-server/user_service/config"
	"message-server/user_service/internal/db"
)

type UserRegisterProcessor struct {
	cfg     *config.ServerConfig
	store   db.StoreQuerier
	request *api.User
}

func NewUserRegisterProcessor(cfg *config.ServerConfig, store db.StoreQuerier, request *api.User) *UserRegisterProcessor {
	return &UserRegisterProcessor{
		cfg:     cfg,
		store:   store,
		request: request,
	}
}

func (u *UserRegisterProcessor) Process(ctx context.Context) error {
	createNewUserParams := db.CreateNewUserParams{
		FirstName: sql.NullString{
			String: u.request.FirstName,
			Valid:  true,
		},
		LastName: sql.NullString{
			String: u.request.LastName,
			Valid:  true,
		},
		Email: u.request.Email,
		Phone: sql.NullString{
			String: u.request.Phone,
			Valid:  true,
		},
		UserName: sql.NullString{
			String: u.request.UserName,
			Valid:  true,
		},
	}
	log.Println(createNewUserParams)
	_, err := u.store.CreateNewUser(ctx, createNewUserParams)
	if err != nil {
		log.Fatalln("Create new user failed")
		return err
	}
	return nil
}
