package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func compare(x []int, e int) bool {
	for _, i := range x {
		if i == e {
			return true
		}
	}
	return false
}

func parse(s []string) []string {

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
		for i := 1; i < len(urlstr); i++ {
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
		if TDATA.Proto == "https://" {
			TDATA.Port = "443"
		} else {
			TDATA.Port = "80"
		}
	}

	// If no servers were specified, just set the host as the server target
	if len(s) < 1 {
		// if urlstr is greater than 1, there is a port defined, and we'll use that as the server. Otherwise, just use the Host that we already parsed out.
		if len(urlstr) > 1 {
			s = append(s, urlstr[0]+":"+urlstr[1])
		} else {
			s = append(s, TDATA.Host)
		}
	} else {
		// If the port is not 80 or 443, we should attach the globally set port to the server name for the appropriate requests to be sent.
		if TDATA.Port != "80" && TDATA.Port != "443" {
			for index, i := range s {
				s[index] = i + ":" + TDATA.Port
			}
		}
	}
	return s
}

func worker(s []string) {
	// Set status as RUNNING, this variable is used to stop the loop of requests being made.
	TDATA.Status = "RUNNING"
	// Check if there is a url defined, otherwise print the usage
	if TDATA.Url == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Check if the expected value is valid
	if TDATA.Expected == 0 {
		TDATA.DisplaySuccess = false
	} else {
		expectedValues := []int{200, 300, 400, 500}
		display := compare(expectedValues, TDATA.Expected)
		if display {
			TDATA.DisplaySuccess = true
		} else {
			fmt.Println("[!] Invalid value with the '-e' flag!")
			flag.Usage()
			os.Exit(1)
		}
	}

	s = parse(s)

	wg.Add(len(s))

	for index, i := range s {
		TRACKINGLIST = append(TRACKINGLIST, tracking{0, 0, 0, 0, 0, 0, i, 0.0, ""})
		if TDATA.Proto == "https://" {
			go MakeHTTPSRequest(i, index, &wg)
		} else {
			go MakeHTTPRequest(i, index, &wg)
		}
	}
	// Update the terminal if it is not in UI mode
	if !TDATA.Ui {
		go update()
	}
	wg.Wait()

	if TDATA.Ui {
		fmt.Println("[!] Stopping threads and starting over.\n ")
	}
	cleanup()
}

func cleanup() {
	TRACKINGLIST = nil
	TDATA.Url = ""
	TDATA.Expected = 0
}
