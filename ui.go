package main

import (
	"net/http"
	"fmt"
	"html/template"
	"encoding/json"
	"os"
)


func serveTemplate(w http.ResponseWriter, r *http.Request) {

	type templatedata struct {
		TD            targetdata
		TL            []tracking	
	}

	td := templatedata{TDATA, TRACKINGLIST}

	t, err := template.ParseFiles("./static/scan/index.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, td)
}


func scanInfo(w http.ResponseWriter, r *http.Request) {
	type Payload struct {
		Servers     []string `json:"servers,omitempty"`
		Url         string   `json:"url"`
		Expected    int      `json:"expected"`
	}
	var m Payload
	err := json.NewDecoder(r.Body).Decode(&m)
	if err!=nil{
		http.Error(w, "Message incorrectly formatted, or blank", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	TDATA.Url = m.Url
	TDATA.Expected = m.Expected
	w.WriteHeader(http.StatusOK)
	go worker(m.Servers)
}

func stop(w http.ResponseWriter, r *http.Request) {
	TDATA.Status = "STOP"
	w.WriteHeader(http.StatusOK)
}

func ui(){
	port := "localhost:5000"
	if os.Getenv("SHARKIE_PORT") != "" {
		port = os.Getenv("SHARKIE_PORT")
	}
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/scan", serveTemplate)
	http.HandleFunc("/data", scanInfo)
	http.HandleFunc("/stop", stop)
	fmt.Println("Server is starting at " + port + "...")
    http.ListenAndServe(port, nil)
}