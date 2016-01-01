package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/danjac/podbaby/commands"
	"github.com/danjac/podbaby/config"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true

	cfg := config.New()

	app.Commands = []cli.Command{
		{
			Name:  "serve",
			Usage: "Run the server",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "url",
					EnvVar:      "DB_URL",
					Usage:       "Database connection URL",
					Destination: &cfg.DatabaseURL,
				},
				cli.StringFlag{
					Name:        "secret",
					EnvVar:      "SECRET_KEY",
					Usage:       "Secret key",
					Destination: &cfg.SecretKey,
				},
				cli.IntFlag{
					Name:        "port",
					Value:       5000,
					EnvVar:      "PORT",
					Usage:       "Server port",
					Destination: &cfg.Port,
				},
				cli.StringFlag{
					Name:        "env",
					Value:       "prod",
					Usage:       "Environment",
					Destination: &cfg.Env,
				},
				cli.StringFlag{
					Name:        "mail-addr",
					EnvVar:      "MAIL_ADDR",
					Value:       "",
					Usage:       "Email address",
					Destination: &cfg.Mail.Addr,
				},
				cli.StringFlag{
					Name:        "mail-host",
					EnvVar:      "MAIL_HOST",
					Value:       "",
					Usage:       "Email host",
					Destination: &cfg.Mail.Host,
				},
				cli.StringFlag{
					Name:        "mail-user",
					EnvVar:      "MAIL_USER",
					Value:       "",
					Usage:       "Email user",
					Destination: &cfg.Mail.User,
				},
				cli.StringFlag{
					Name:        "mail-password",
					EnvVar:      "MAIL_PASSWORD",
					Value:       "",
					Usage:       "Email password",
					Destination: &cfg.Mail.Password,
				},
				cli.StringFlag{
					Name:        "mail-id",
					EnvVar:      "MAIL_ID",
					Value:       "",
					Usage:       "Email identity",
					Destination: &cfg.Mail.ID,
				},
			},
			Action: func(c *cli.Context) {
				cfg.MustValidate()
				commands.Serve(cfg)
			},
		},
		{
			Name:  "fetch",
			Usage: "Fetch new podcasts",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "url",
					EnvVar:      "DB_URL",
					Usage:       "Database connection URL",
					Destination: &cfg.DatabaseURL,
				},
			},
			Action: func(c *cli.Context) {
				cfg.MustValidate()
				commands.Fetch(cfg)
			},
		},
	}

	app.Run(os.Args)

}
