package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/EvatixGo/parse"
	"github.com/gocarina/gocsv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./CliRunner csv-string")
		os.Exit(1)
	}

	csv := os.Args[1]
	csv = strings.ReplaceAll(csv, "\\n", "\n")

	var cliRunners []parse.CliRunnerRecord
	gocsv.UnmarshalString(csv, &cliRunners)

	var wg sync.WaitGroup
	file, err := os.Create("./output.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	mutex := &sync.Mutex{}

	for _, c := range cliRunners {

		cmd := exec.Command("../CliStreamer/CliStreamer", c.CliStreamerRecordCsv())
		cmdReader, _ := cmd.StdoutPipe()
		scanner := bufio.NewScanner(cmdReader)

		// done := make(chan bool)

		wg.Add(1)
		go func() {
			defer wg.Done()
			for scanner.Scan() {
				txt := scanner.Text()

				mutex.Lock()
				fmt.Println(txt)
				file.WriteString(txt + "\n")
				mutex.Unlock()
			}
		}()

		cmd.Start()
	}

	wg.Wait()
}
