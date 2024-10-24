package scrapper

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-rod/rod"
)

type Mountains struct {
	Peak_id      int
	Name         string
	Alias        string
	Height       float32
	Peak_range   string
	OpenToPublic bool
}

func ScrapePeaks(Peaks *[]Mountains) {
	*Peaks = []Mountains{}
	browser := rod.New().NoDefaultDevice().MustConnect()
	defer browser.MustClose()
	page := browser.MustPage("https://nepalhimalpeakprofile.org/peak-profile")
	page.MustWindowFullscreen()
	page.MustWaitLoad()

	rows := page.MustElements("table#mountaintable tbody tr")

	for _, row := range rows {
		var peak Mountains

		peak.Peak_id = len(*Peaks) + 1
		peak.Alias = row.MustElement("td:nth-of-type(1)").MustText()
		peak.Name = row.MustElement("td:nth-of-type(2)").MustText()
		peak.Height = parseHeight(row.MustElement("td:nth-of-type(3)").MustText())
		peak.Peak_range = row.MustElement("td:nth-of-type(4)").MustText()
		peak.OpenToPublic = parseOpenToPublic(row.MustElement("td:nth-of-type(5)").MustText())

		*Peaks = append(*Peaks, peak)
	}

	if len(*Peaks) > 0 {
		(*Peaks)[0].Name = "Mount Everest"
	}

}

func parseHeight(heightStr string) float32 {
	var height float32
	_, err := fmt.Sscanf(heightStr, "%f", &height)
	if err != nil {
		fmt.Println("Error parsing height:", err)
	}
	return height
}

func parseOpenToPublic(openStr string) bool {
	return openStr == "Opened"
}

func WriteToJson(Peaks []Mountains) {
	jsonData, err := json.MarshalIndent(Peaks, "", "")
	if err != nil {
		fmt.Println(err)
	}
	file, err := os.Create("peaks.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println(err)
	}

}
