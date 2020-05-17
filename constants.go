package main
import (
	"strings"
	"strconv"
)

type targetdata struct {
	Url            string
	Host           string
	Path           string
	Proto          string
	Port           string
	Expected       int
	DisplaySuccess bool
	SkipTLS        bool
	Emoji          bool
	Sleep          float64
}

var TDATA targetdata

type tracking struct {
	Twohundreds     int
	Threehundreds   int
	Fourhundreds    int
	Fivehundreds    int
	Failed          int
	Total           int
	Server          string
}

var TRACKINGLIST []tracking

// Set up emoji's
func setemoji() map[string]string {
	var emoji map[string]string
	if TDATA.Emoji {
		shark, _ := strconv.ParseInt(strings.TrimPrefix("\\U1F988", "\\U"), 16, 32)
		thumbup, _ := strconv.ParseInt(strings.TrimPrefix("\\U1F44D", "\\U"), 16, 32)
		thumbdown, _ := strconv.ParseInt(strings.TrimPrefix("\\U1F44E", "\\U"), 16, 32)
		sad, _ := strconv.ParseInt(strings.TrimPrefix("\\U1F629", "\\U"), 16, 32)
		eyebrow, _ := strconv.ParseInt(strings.TrimPrefix("\\U1F928", "\\U"), 16, 32)
		neutral, _ := strconv.ParseInt(strings.TrimPrefix("\\U1F610", "\\U"), 16, 32)
		emoji = map[string]string {"shark":string(shark), 
									"thumbup":string(thumbup), 
									"thumbdown":string(thumbdown), 
									"sad":string(sad),
									"eyebrow":string(eyebrow),
									"neutral":string(neutral)}
	} else {
		// If emojis are disabled, we will return a map with empty strings
		emoji = map[string]string {"shark":"", 
									"thumbup":"", 
									"thumbdown":"", 
									"sad":"",
									"eyebrow":"",
									"neutral":""}

	}
	return emoji
}

