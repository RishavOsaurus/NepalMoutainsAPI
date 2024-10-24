package main

import (
	"fmt"
	"mountainsapi/scrapper"
	"time"
)

func main() {
	var Peaks []scrapper.Mountains
	for {
		// Create a ticker that triggers every 30 seconds
		timer1 := time.NewTicker(24 * time.Hour)

		scrapper.ScrapePeaks(&Peaks)
		scrapper.WriteToJson(Peaks)
		fmt.Println("Ok NEXT ITERATION")
		<-timer1.C
	}
}
