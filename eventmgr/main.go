package main

import (
	"os"

	"github.com/gautamrege/sweatbead/eventmgr/app"
	"github.com/gautamrege/sweatbead/eventmgr/config"
	"github.com/gautamrege/sweatbead/eventmgr/server"
	"github.com/urfave/cli"
)

func main() {
	config.Load()
	app.Init()
	defer app.Close()

	cliApp := cli.NewApp()
	cliApp.Name = "SweatBead EventMgr"
	cliApp.Version = "1.0.0"
	cliApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start server",
			Action: func(c *cli.Context) error {
				server.StartAPIServer()
				return nil
			},
		},
	}
	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}
