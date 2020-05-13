package main

import (
	"fmt"
	"flag"
	"strings"
	"sync"
)

type servers []string

func (x *servers) Set(value string) error {
	*x = append(*x, value)
	return nil
}

func (x *servers) String() string {
	return fmt.Sprint(*x)
}

var s servers
var wg sync.WaitGroup

func main () {
	flag.Var(&s, "s", "Server IP or hostname")
	flag.StringVar(&TDATA.Url, "u", "", "URL to target")
	flag.Float64Var(&TDATA.Sleep, "sleep", 1, "Time in seconds to sleep between requests")
	flag.Parse()
	
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