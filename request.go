// SNI issues: https://github.com/golang/go/issues/22704

package main

import (
	"net/http"
	"log"
	"fmt"
	"time"
	"net"
	//"strings"
	"context"
)

func MakeRequest(server string, t targetdata){

	dialer := &net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 10 * time.Second,
		DualStack: true,
	}

	for {
		//req, _ := http.NewRequest("GET", t.proto + server + t.path, nil)
		req, _ := http.NewRequest("GET", t.url, nil)
		//req.Header.Add("Host", t.host)
		req.Host = t.host
		//req.URL.Host = server

		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
        	return http.ErrUseLastResponse
			},
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				// redirect all connections to 127.0.0.1
				//addr = server + addr[strings.LastIndex(addr, ":"):]
				addr = server + ":" + t.port
				return dialer.DialContext(ctx, network, addr)
				},
			},
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(server + " |", resp.StatusCode)
		time.Sleep((time.Duration(t.sleep * 1000)) * time.Millisecond)
	}
}
