package main

import (
	"fmt"
	"flag"
)

type servers []string

func (x *servers) Set(value string) error {
	*x = append(*x, value)
	return nil
}

func (x *servers) String() string {
	return fmt.Sprint(*x)
}

var EMOJI map[string]string
var s servers

func main () {
	flag.Usage = func(){
		fmt.Println("Welcome to Sharkie! A CLI tool for tracking HTTP response codes.\nExample:\nsharkie -u example.com")
		fmt.Println("\n\nUse the -s flag to target multiple servers behind a load balancer with the same HTTP Host header:\nsharkie -u example.com -s 1.2.3.4 -s 3.4.5.6")
		fmt.Println("\n\nTrack success rate based on expected status code:\nsharkie -u https://example.com -s 1.2.3.4 -s 3.4.5.6 -e 200\n ")
		flag.PrintDefaults()
	}
	flag.Var(&s, "s", "Server IP or hostname (127.0.0.1, example.com)")
	flag.StringVar(&TDATA.Url, "u", "", "URL to target")
	flag.Float64Var(&TDATA.Sleep, "sleep", 1, "Time in seconds to sleep between requests")
	flag.IntVar(&TDATA.Expected, "e", 0, "Expected HTTP status code to generate success percentages. For example:\n-e 200 (200-299)\n-e 300 (300-399) etc.\nValid values are: 200, 300, 400, 500")
	flag.BoolVar(&TDATA.SkipTLS, "k", false, "Ignore invalid certificates for TLS connections. (Default is false)\nUsage: -k=true")
	flag.BoolVar(&TDATA.Emoji, "emoji", true, "Control whether emoji's display or not.\nUsage (to turn off): -emoji=false")
	flag.BoolVar(&TDATA.Ui, "ui", false, "Enable the UI and HTTP server with -ui=true")
	flag.Parse()
	EMOJI = setemoji()

	if !TDATA.Ui{
		worker(s)
	} else{
		ui()
	}
}