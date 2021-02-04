package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/EvatixGo/parse"
	"github.com/gocarina/gocsv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./CliStreamer csv-string")
		os.Exit(1)
	}

	csv := os.Args[1]
	csv = strings.ReplaceAll(csv, "\\n", "\n")

	var cliStreamerRecords []parse.CliStreamerRecord
	gocsv.UnmarshalString(csv, &cliStreamerRecords)

	var wg sync.WaitGroup

	for _, c := range cliStreamerRecords {

		wg.Add(1)
		go func(c parse.CliStreamerRecord) {
			defer wg.Done()
			for i := 0; i < c.RunTimes; i++ {
				fmt.Printf("%s->%s\n", c.Title, c.Message1)
				time.Sleep(time.Duration(c.StreamDelay) * time.Second)
				fmt.Printf("%s->%s\n", c.Title, c.Message2)
			}
		}(c)

	}

	wg.Wait()
}
