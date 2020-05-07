package main

import (
	"fmt"
	"flag"
	"strings"
	"sync"
)

type Servers []string

func (x *Servers) Set(value string) error {
	*x = append(*x, value)
	return nil
}

func (x *Servers) String() string {
	return fmt.Sprint(*x)
}

var servers Servers
var wg sync.WaitGroup

type targetdata struct {
	url   string
	host  string
	path  string
	proto string
	port  string
	sleep float64
}

func main () {
	var t targetdata
	flag.Var(&servers, "s", "Server IP or hostname")
	flag.StringVar(&t.url, "u", "", "URL to target")
	flag.Float64Var(&t.sleep, "sleep", 1, "Time in seconds to sleep between requests")
	flag.Parse()
	fmt.Println("Servers:", servers)
	
	// Get the host header from the URL
	urlstr := strings.Split(t.url, "://")

	// assuming the URL string has http:// or https:// we will then split the rest to get the path
	if len(urlstr) > 1 {
		t.proto = urlstr[0] + "://"
		urlstr = strings.Split(urlstr[1], "/")
	} else {
		// if there is no '://' then we can just assume http://
		t.proto = "http://"
		urlstr = strings.Split(urlstr[0], "/")
	}

	// reconstruct full path
	if len(urlstr) > 1 {
		t.path = ""
		for i := 1; i < len(urlstr); i++{
			t.path += "/" + urlstr[i]
		}
	} else {
		// if there is not a '/' in the URL, then we can just assume it is the root.
		t.path = "/"
	}

	// Splitting on ':' incase a port number was specified.
	urlstr = strings.Split(urlstr[0], ":")
	t.host = urlstr[0]
	// get the port used
	if len(urlstr) > 1 {
		t.port = urlstr[1]
	} else {
		if t.proto == "https://"{
			t.port = "443"
		} else {
			t.port = "80"
		}
	}

	fmt.Println("URL:", t.url)
	fmt.Println("Host:", t.host)
	fmt.Println("Path:", t.path)
	fmt.Println("Proto:", t.proto)
	fmt.Println("Port:", t.port)
	fmt.Println("Sleep Time:", t.sleep)

	wg.Add(len(servers))
	for _, i := range servers{
		go MakeRequest(i, t)
	}

	wg.Wait()
}