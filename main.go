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
		{
			Name:   "flush",
			Action: Flush,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "tube,t",
				},
			},
		},
		{
			Name:   "peek",
			Action: Peek,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "tube,t",
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
