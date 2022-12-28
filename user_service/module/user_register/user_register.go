package user_register

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
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
	result, err := u.store.CreateNewUser(ctx, createNewUserParams)
	if err != nil {
		log.Fatalln("Create new user failed")
		return err
	}
	lastUserId, err := result.LastInsertId()
	if err != nil {
		log.Fatalln("Get last user id fail")
		return err
	}
	CreateNewUserCredentialParams := db.CreateNewUserCredentialParams{
		UserID: lastUserId,
		PasswordHashed: sql.NullString{
			String: fmt.Sprint(sha256.Sum256([]byte(u.request.Password))),
			Valid:  true,
		},
	}
	_, err = u.store.CreateNewUserCredential(ctx, CreateNewUserCredentialParams)
	if err != nil {
		log.Fatalln("Create credential for user fail!")
		return err
	}
	return nil
}
