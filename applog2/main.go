package main

import (
	//	"bufio"
	"bytes"
	"fmt"
	//	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type LogItemMap map[string]map[string][]time.Time

func CallLogger(lis LogItemMap) {

	// Command to be executed
	cmd := exec.Command("wmctrl", "-lxG")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// 2nd separator for output
	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	// Split newlines
	output := strings.Split(out.String(), "\n")
	output = output[:len(output)-1]

	// Time
	t := time.Now()

	// Strip spaces and add to map
	for _, item := range output {
		strip := []string{}
		myitem := strings.Split(item, " ")
		for _, elem := range myitem {
			if elem != "" {
				strip = append(strip, elem)
			}
		}

		sepa := host + " "

		// Proc name and title
		name := strip[6]
		title := item[strings.Index(item, sepa)+len(sepa):]

		if lis[name] == nil {
			lis[name] = map[string][]time.Time{title: []time.Time{t}}
		} else {
			lis[name][title] = append(lis[name][title], t)
		}

	}
}

func main() {
	s := time.Duration(10)
	lim := make(LogItemMap)
	for true {
		CallLogger(lim)

		f, err := os.Create("/tmp/log.txt")
		check(err)
		defer f.Close()

		for keys, vals := range lim {
			_, err := f.Write([]byte(keys))
			check(err)
			fmt.Printf("%v:\n", keys)
			for key, val := range vals {
				_, err := f.Write([]byte(string(len(val) * int(s))))
				check(err)
				_, err = f.Write([]byte(key))
				check(err)
				fmt.Printf("\t%v\t%v\n", len(val)*int(s), key)
			}
		}
		fmt.Println(len(lim))
		//fmt.Println(lim)

		time.Sleep(s * time.Second)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
