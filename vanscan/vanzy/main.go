package main

import (
	"go_scan/vanzy/getfinger"
	"go_scan/vanzy/myflag"
	"bufio"
	"flag"
	"os"
	"fmt"
)

var URL = flag.String("url", "", "input url")
var Urllist = flag.String("file", "", "input path to urllist")
var Fi1e = flag.Bool("f", false, "See the result in the file")
func main(){
	myflag.Banner()
	flag.Parse()
	if *URL != "" {
		getfinger.Run(*URL,*Fi1e)
	}
	if *Urllist != "" {
		file, err := os.Open(*Urllist)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		for _, line := range lines {
			go getfinger.Run(line,*Fi1e)
		}
	}
	if *URL == "" && *Urllist == "" {
		fmt.Println("--url  for the single target")
		fmt.Println("--file  for lots of targets")
		os.Exit(0)
	}
}
