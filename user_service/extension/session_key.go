package extension

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Session struct {
	UserId     int64
	Username   string
	Email      string
	ExpireTime time.Time
}

func NewSession(userId int64, username string, email string, expireTime time.Time) *Session {
	return &Session{
		UserId:     userId,
		Username:   username,
		Email:      email,
		ExpireTime: expireTime,
	}
}

func (s *Session) GenKey() string {
	sessionToString := fmt.Sprintf("%s%s%s%s",
		strconv.FormatInt(s.UserId, 10),
		s.Username,
		s.Email,
		s.ExpireTime,
	)

	return fmt.Sprintf("%x", sha256.Sum256([]byte(sessionToString)))
}
