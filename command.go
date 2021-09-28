package main

import (
	"fmt"
	"github.com/stevedomin/termtable"
	"github.com/urfave/cli"
	"sort"
	"time"
)

const DefaultTube = "default"

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
	sort.Strings(tubes)
	fmt.Println(DefaultTube, statTube(DefaultTube)["current-jobs-ready"])
	for _, tube := range tubes {
		if tube != DefaultTube {
			fmt.Println(tube, statTube(tube)["current-jobs-ready"])
		}
	}
}

func Stats(_ *cli.Context) {
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
	sort.Strings(tubes)
	var results []map[string]string
	results = append(results, statTube(DefaultTube))
	for _, tube := range tubes {
		if tube != DefaultTube {
			results = append(results, statTube(tube))
		}
	}
	printStats(results)
}

func statTube(tubeName string) map[string]string {
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
	return result
}

func printStats(stats []map[string]string) {
	table := termtable.NewTable(nil, &termtable.TableOptions{
		Padding:      1,
		UseSeparator: true,
	})
	table.SetHeader([]string{
		"Name", "Urgent", "Ready", "Reserved", "Delayed", "Buried",
	})
	keys := []string{
		"name",
		"current-jobs-urgent",
		"current-jobs-ready",
		"current-jobs-reserved",
		"current-jobs-delayed",
		"current-jobs-buried",
	}
	for _, stat := range stats {
		var row []string
		for _, key := range keys {
			row = append(row, stat[key])
		}
		table.AddRow(row)
	}
	fmt.Println(table.Render())
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

func Peek(ctx *cli.Context) {
	tubeName := ctx.String("tube")
	tubeSet := NewTubeSet(tubeName)
	defer func() {
		err := tubeSet.Conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	jobID, jobBody, err := tubeSet.Reserve(time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jobBody))
	err = tubeSet.Conn.Release(jobID, 0, 0)
	if err != nil {
		panic(err)
	}
}
