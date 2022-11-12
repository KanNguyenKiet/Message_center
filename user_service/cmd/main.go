package main

import (
	"fmt"
	"log"
	"message-server/common/database/migration"
	"message-server/user_service/config"
	"message-server/user_service/internal/service"
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
	cfg, err = config.LoadConfig()
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
		{
			Name:        "migrate",
			Usage:       "Migrate database",
			Subcommands: migration.CliCommand(cfg.MigrateFolder, cfg.Database.String()),
		},
	}

	if err = app.Run(os.Args); err != nil {
		return err
	}

	return nil
}

func run(cliCtx *cli.Context) error {
	err := service.CreateServer(cfg)
	if err != nil {
		log.Fatalln("Create server error", err)
		return err
	}
	return nil
}
