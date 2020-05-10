package main

import(
	"fmt"
	"os"
	"os/exec"
	"text/tabwriter"
	"time"
	"strconv"
	"strings"
)

func update()  {
	for {
		cmd := exec.Command("clear") 
        cmd.Stdout = os.Stdout
		cmd.Run()

		// Get a shark emoji :)
		emoji, _ := strconv.ParseInt(strings.TrimPrefix("\\U1F988", "\\U"), 16, 32)

		fmt.Printf("===========\nSHARKIE  %s\n===========\n", string(emoji))
		fmt.Println("URL:", TDATA.Url)
		fmt.Println("Host:", TDATA.Host)
		fmt.Println("Path:", TDATA.Path)
		fmt.Println("Proto:", TDATA.Proto)
		fmt.Println("Port:", TDATA.Port)
		fmt.Println("Sleep Time:", TDATA.Sleep)

		// Loop through TRACKINGLIST and print all the current values
		w := new(tabwriter.Writer)
		// minwidth, tabwidth, padding, padchar, flags
		w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t", "Server", "200s", "300s", "400s", "500s", "Failed", "Total")
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t", "------", "----", "----", "----", "----", "------", "-----")
	
		for _, i := range TRACKINGLIST {
 	       	fmt.Fprintf(w, "\n %s\t%d\t%d\t%d\t%d\t%d\t%d\t", i.Server, i.Twohundreds, i.Threehundreds, i.Fourhundreds, i.Fivehundreds, i.Failed, i.Total)
		}
		w.Flush()
		fmt.Println("\n\n")
		time.Sleep((time.Duration(500)) * time.Millisecond)
	}
}