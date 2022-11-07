package api_handler

import (
	"fmt"
	"io"
	"message-server/user_service/config"
	"net/http"
)

func GetServerName(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /get Env request\n")
	cfg, _ := config.LoadConfig()
	io.WriteString(w, fmt.Sprintf("You are running app on %s\n", cfg.Env))
}
