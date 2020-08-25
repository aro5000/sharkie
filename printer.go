package main

import(
	"fmt"
	"os"
	"os/exec"
	"text/tabwriter"
	"time"
)

func update()  {
	for {
		cmd := exec.Command("clear") 
        cmd.Stdout = os.Stdout
		cmd.Run()

		fmt.Printf("===========\nSHARKIE  %s\n===========\n", EMOJI["shark"])
		fmt.Println("URL:", TDATA.Url)
		fmt.Println("Host:", TDATA.Host)
		fmt.Println("Path:", TDATA.Path)
		fmt.Println("Proto:", TDATA.Proto)
		fmt.Println("Port:", TDATA.Port)
		fmt.Println("Sleep Time:", TDATA.Sleep)

		// Loop through TRACKINGLIST and print all the current values
		w := new(tabwriter.Writer)
		// minwidth, tabwidth, padding, padchar, flags
		w.Init(os.Stdout, 8, 8, 1, '\t', 0)
	
		// Check if we should display success rates depending on the expected value:
		if TDATA.DisplaySuccess{
			fmt.Println("Expected Response:", TDATA.Expected)
			fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t", "Server", "200s", "300s", "400s", "500s", "Failed", "Total", "Success %")
			fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t", "------", "----", "----", "----", "----", "------", "-----", "---------")
			for _, i := range TRACKINGLIST {
 		       	fmt.Fprintf(w, "\n %s\t%d\t%d\t%d\t%d\t%d\t%d\t%.2f\t%s\t", i.Server, i.Twohundreds, i.Threehundreds, i.Fourhundreds, i.Fivehundreds, i.Failed, i.Total, i.Percent, i.Emoji)
			}
		} else {
			fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t", "Server", "200s", "300s", "400s", "500s", "Failed", "Total")
			fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t", "------", "----", "----", "----", "----", "------", "-----")
			for _, i := range TRACKINGLIST {
 		       	fmt.Fprintf(w, "\n %s\t%d\t%d\t%d\t%d\t%d\t%d\t", i.Server, i.Twohundreds, i.Threehundreds, i.Fourhundreds, i.Fivehundreds, i.Failed, i.Total)
			}
		}
		w.Flush()
		fmt.Println("\n\n")
		time.Sleep((time.Duration(500)) * time.Millisecond)
	}
}
