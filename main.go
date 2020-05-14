package main

import (
	"fmt"
	"flag"
	"strings"
	"sync"
	"os"
)

type servers []string

func (x *servers) Set(value string) error {
	*x = append(*x, value)
	return nil
}

func (x *servers) String() string {
	return fmt.Sprint(*x)
}

func compare(x []int, e int) bool{
	for _, i := range(x){
		if i == e {
			return true
		}
	}
	return false
	}

var s servers
var wg sync.WaitGroup

func main () {
	flag.Var(&s, "s", "Server IP or hostname")
	flag.StringVar(&TDATA.Url, "u", "", "URL to target")
	flag.Float64Var(&TDATA.Sleep, "sleep", 1, "Time in seconds to sleep between requests")
	flag.IntVar(&TDATA.Expected, "e", 0, "Expected HTTP status code to generate success percentages. For example:\n-e 200 (200-299)\n-e 300 (300-399) etc.\nValid values are: 200, 300, 400, 500")
	flag.Parse()

	// Check if there is a url defined, otherwise print the usage
	if TDATA.Url == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Check if the expected value is valid
	if TDATA.Expected == 0 {
		TDATA.DisplaySuccess = false
	} else {
		expectedValues := []int{200,300,400,500}
		display := compare(expectedValues, TDATA.Expected)
		if display{
			TDATA.DisplaySuccess = true
		} else{
			fmt.Println("[!] Invalid value with the '-e' flag!")
			flag.Usage()
			os.Exit(1)
		}
	}

	// Get the host header from the URL
	urlstr := strings.Split(TDATA.Url, "://")

	// assuming the URL string has http:// or https:// we will then split the rest to get the path
	if len(urlstr) > 1 {
		TDATA.Proto = urlstr[0] + "://"
		urlstr = strings.Split(urlstr[1], "/")
	} else {
		// if there is no '://' then we can just assume http://
		TDATA.Proto = "http://"
		urlstr = strings.Split(urlstr[0], "/")
	}

	// reconstruct full path
	if len(urlstr) > 1 {
		TDATA.Path = ""
		for i := 1; i < len(urlstr); i++{
			TDATA.Path += "/" + urlstr[i]
		}
	} else {
		// if there is not a '/' in the URL, then we can just assume it is the root.
		TDATA.Path = "/"
	}

	// Splitting on ':' incase a port number was specified.
	urlstr = strings.Split(urlstr[0], ":")
	TDATA.Host = urlstr[0]
	// get the port used
	if len(urlstr) > 1 {
		TDATA.Port = urlstr[1]
	} else {
		if TDATA.Proto == "https://"{
			TDATA.Port = "443"
		} else {
			TDATA.Port = "80"
		}
	}

	// If no servers were specified, just set the host as the server target
	if len(s) < 1 {
		s = append(s, TDATA.Host)
	}

	wg.Add(len(s) + 1)
	for index, i := range s{
		TRACKINGLIST = append(TRACKINGLIST, tracking{0,0,0,0,0,0,i})
		if TDATA.Proto == "https://"{
			go MakeHTTPSRequest(i, index)
		}else {
			go MakeHTTPRequest(i, index)
		}
	}
	// Update the terminal
	go update()
	wg.Wait()
}