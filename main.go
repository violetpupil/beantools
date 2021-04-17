package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:   "stat",
			Action: Stat,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
