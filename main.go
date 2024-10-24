package main

import gogin "mountainsapi/go-gin"

func main() {
	// var Peaks []scrapper.Mountains
	// for {

	// 	timer1 := time.NewTicker(20 * time.Second)

	// 	scrapper.ScrapePeaks(&Peaks)
	// 	scrapper.WriteToJson(Peaks)

	// 	<-timer1.C
	// }

	gogin.GinAPI()
}
