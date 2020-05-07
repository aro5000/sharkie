// SNI issues: https://github.com/golang/go/issues/22704

package main

import (
	"net/http"
	"log"
	"fmt"
	"time"
	"net"
	"strings"
	"context"
	"io/ioutil"
)


func MakeHTTPSRequest(server string, t targetdata){

	// custom dialer for HTTPS requests to control SNI
	dialer := &net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 10 * time.Second,
		DualStack: true,
	}

	for {
		req, _ := http.NewRequest("GET", t.url, nil)

		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
        	return http.ErrUseLastResponse
			},
			// We make a custom transport here because according to RFC the SNI field does not have to be the same as the Host header.
			// This transport forces the connection to go to the host we want.
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				// Use the port determined by the URL unless the server specifies otherwise
				if strings.Contains(server, ":"){
					addr = server
				}else { 
					addr = server + ":" + t.port
					}
				return dialer.DialContext(ctx, network, addr)
				},
			},
			Timeout: time.Second * 10,
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode == http.StatusOK{
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(server + " | " + fmt.Sprint(resp.StatusCode) + " | " + string(body))
		}else{
			fmt.Println(server + " |", resp.StatusCode)
		}
		time.Sleep((time.Duration(t.sleep * 1000)) * time.Millisecond)
	}
}


func MakeHTTPRequest(server string, t targetdata) {
	for {
		req, _ := http.NewRequest("GET", t.proto + server + t.path, nil)
		req.Host = t.host
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
        	return http.ErrUseLastResponse
			},
			Timeout: time.Second * 10,
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(server + " |", resp.StatusCode)
		time.Sleep((time.Duration(t.sleep * 1000)) * time.Millisecond)

	}
}