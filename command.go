package main

import (
	"fmt"
	"github.com/urfave/cli"
	"time"
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

func Flush(ctx *cli.Context) {
	tubeName := ctx.String("tube")
	tubeSet := NewTubeSet(tubeName)
	defer func() {
		err := tubeSet.Conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	for {
		jobID, _, err := tubeSet.Reserve(time.Second)
		if err != nil {
			if err.Error() == "reserve-with-timeout: timeout" {
				break
			} else {
				panic(err)
			}
		}
		err = tubeSet.Conn.Delete(jobID)
		if err != nil {
			panic(err)
		}
	}
}
