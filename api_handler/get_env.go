package api_handler

import (
	"fmt"
	"io"
	"message-server/config"
	"net/http"
)

func GetServerName(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /get Env request\n")
	cfg, _ := config.DefaultLoad()
	io.WriteString(w, fmt.Sprintf("You are running app on %s\n", cfg.Env))
}
