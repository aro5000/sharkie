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
		w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	
		// Check if we should display success rates depending on the expected value:
		if TDATA.DisplaySuccess{
			fmt.Println("Expected Response:", TDATA.Expected)
			fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t", "Server", "200s", "300s", "400s", "500s", "Failed", "Total", "Success %")
			fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t", "------", "----", "----", "----", "----", "------", "-----", "---------")
			for _, i := range TRACKINGLIST {
				var percent float64
				switch TDATA.Expected {
				case 200: percent = float64(i.Twohundreds) / float64(i.Total) * float64(100)
				case 300: percent = float64(i.Threehundreds) / float64(i.Total) * float64(100)
				case 400: percent = float64(i.Fourhundreds) / float64(i.Total) * float64(100)
				case 500: percent = float64(i.Fivehundreds) / float64(i.Total) * float64(100)
				}
				// Figure out which emoji to use
				var emoji string
				switch {
				case (percent == 100): emoji = EMOJI["thumbup"]
				case (50.0 <= percent && percent <= 100.0): emoji = EMOJI["thumbdown"] 
				case (percent < 50): emoji = EMOJI["sad"]
				default: emoji = ""
				}
 		       	fmt.Fprintf(w, "\n %s\t%d\t%d\t%d\t%d\t%d\t%d\t%.2f\t%s\t", i.Server, i.Twohundreds, i.Threehundreds, i.Fourhundreds, i.Fivehundreds, i.Failed, i.Total, percent, emoji)
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
