// SNI issues: https://github.com/golang/go/issues/22704

package main

import (
	"net/http"
	"time"
	"net"
	"strings"
	"context"
	"crypto/tls"
)


func MakeHTTPSRequest(server string, index int){
	for {
		req, _ := http.NewRequest("GET", TDATA.Url, nil)
		req.Header.Set("User-Agent", "sharkie")

		client := setTlsClient(server)
		resp, err := client.Do(req)
		if err != nil {
			evaluate(0, index)
		} else{
			evaluate(resp.StatusCode, index)
		}
		time.Sleep((time.Duration(TDATA.Sleep * 1000)) * time.Millisecond)
	}
}


func MakeHTTPRequest(server string, index int) {
	for {
		req, _ := http.NewRequest("GET", TDATA.Proto + server + TDATA.Path, nil)
		req.Header.Set("User-Agent", "sharkie")
		req.Host = TDATA.Host
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
        	return http.ErrUseLastResponse
			},
			Timeout: time.Second * 10,
		}
		resp, err := client.Do(req)
		if err != nil {
			evaluate(0, index)
		} else{
			evaluate(resp.StatusCode, index)
		}

		time.Sleep((time.Duration(TDATA.Sleep * 1000)) * time.Millisecond)
	}
}


func evaluate(response_code int, index int){
	switch  {
	case (200 <= response_code && response_code <= 299):
		TRACKINGLIST[index].Twohundreds += 1
	case (300 <= response_code && response_code <= 399):
		TRACKINGLIST[index].Threehundreds += 1
	case (400 <= response_code && response_code <= 499):
		TRACKINGLIST[index].Fourhundreds += 1
	case (500 <= response_code && response_code <= 599):
		TRACKINGLIST[index].Fivehundreds += 1
	default:
		TRACKINGLIST[index].Failed += 1
	}
	TRACKINGLIST[index].Total += 1
}


func setTlsClient(server string) *http.Client {
	// custom dialer for HTTPS requests to control SNI
	dialer := &net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 10 * time.Second,
		DualStack: true,
	}
	// if the flag to skip TLS validation was set, we'll return a client with that set
	if TDATA.SkipTLS {
		return &http.Client{
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
							addr = server + ":" + TDATA.Port
							}
						return dialer.DialContext(ctx, network, addr)
						},
						TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
					},
					Timeout: time.Second * 10,
				}
			} else{
				return &http.Client{
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
							addr = server + ":" + TDATA.Port
							}
						return dialer.DialContext(ctx, network, addr)
						},
					},
					Timeout: time.Second * 10,
				}
			}
		}