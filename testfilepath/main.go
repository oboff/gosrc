package main

import (
	"math/rand"
	//"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var wallpapers []string

func printFile(dirpath string, info os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}
	suf := []string{"jpg", "png", "gif", "jpeg"}
	for _, s := range suf {
		if strings.HasSuffix(dirpath, s) || strings.HasSuffix(dirpath, strings.ToUpper(s)) {
			wallpapers = append(wallpapers, dirpath)
		}
	}

	//fmt.Println(wallpapers)
	return nil
}

func rotate() error {
	w := wallpapers[rand.Intn(len(wallpapers))]
	cmd := exec.Command("feh", "--bg-fill", w)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func main() {
	//time.Sleep(time.Second * 5)
	log.SetFlags(log.Lshortfile)
	if len(os.Args) > 1 {
		dir := os.Args[1]
		err := filepath.Walk(dir, printFile)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(wallpapers)
		for {
			//fmt.Printf("%v\n", w)
			rotate()
			time.Sleep(time.Minute * 15)
		}
	}

}
