package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func CallLogger() {
	// Command to be executed
	cmd := exec.Command("wmctrl", "-lxG")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Parse output into lines
	output := strings.Split(out.String(), "\n")
	output = output[:len(output)-1]

	// Prepare time for parsing
	t := time.Now()

	// Separator to clean output
	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range output {
		strip := []string{}
		myitem := strings.Split(item, " ")
		for _, elem := range myitem {
			if elem != "" {
				strip = append(strip, elem)
			}
		}
		sepa := host + " "
		procname := strip[6]
		proctitle := item[strings.Index(item, sepa)+len(sepa):]
		fmt.Printf("%v %q %q\n", t, procname, proctitle)

	}
}

func main() {
	for true {
		CallLogger()
		time.Sleep(10 * time.Second)
	}
}
