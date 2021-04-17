package main

import (
	"fmt"
	"github.com/urfave/cli"
)

func Stat(_ *cli.Context) {
	conn := NewConn()
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	tubes, err := conn.ListTubes()
	if err != nil {
		panic(err)
	}
	for _, tube := range tubes {
		statTube(tube)
	}
}

func statTube(tubeName string) {
	tube := NewTube(tubeName)
	defer func() {
		err := tube.Conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	result, err := tube.Stats()
	if err != nil {
		panic(err)
	}
	readyCount := result["current-jobs-ready"]
	fmt.Println(tubeName, readyCount)
}
