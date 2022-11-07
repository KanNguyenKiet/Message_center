package main

import (
	"fmt"
	"log"
	"message-server/api_handler"
	"message-server/config"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	if err := runServer(os.Args); err != nil {
		log.Fatal(err)
	}
}

var cfg *config.ServerConfig

func runServer(args []string) error {
	var err error
	cfg, err = config.DefaultLoad()
	if err != nil {
		log.Fatal(err)
		return err
	}

	if cfg.Env == "local" {
		fmt.Printf("%+v\n", cfg)
	}

	app := &cli.App{
		Name:  "message-server",
		Usage: "send and receive message between two user",
	}

	app.Commands = []*cli.Command{
		{
			Name:   "server",
			Usage:  "Run message center server",
			Action: run,
		},
	}

	if err = app.Run(os.Args); err != nil {
		return err
	}

	return nil
}

func run(cliCtx *cli.Context) error {
	muxServer := http.NewServeMux()
	muxServer.HandleFunc("/get-env", api_handler.GetServerName)

	if err := http.ListenAndServe(cfg.Host+":"+cfg.Port, muxServer); err != nil {
		return err
	}
	return nil
}
