package main

import (
	"fmt"

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

func main() {
	var Peaks []Mountains
	browser := rod.New().NoDefaultDevice().MustConnect()
	defer browser.MustClose()
	page := browser.MustPage("https://nepalhimalpeakprofile.org/peak-profile")
	page.MustWindowFullscreen()
	page.MustWaitLoad()

	rows := page.MustElements("table#mountaintable tbody tr")
	for _, row := range rows {
		var peak Mountains

		// Assuming the correct order of columns in the table
		peak.Peak_id = len(Peaks) + 1 // Use a simple incremental ID
		peak.Alias = row.MustElement("td:nth-of-type(1)").MustText()
		peak.Name = row.MustElement("td:nth-of-type(2)").MustText()
		peak.Height = parseHeight(row.MustElement("td:nth-of-type(3)").MustText())
		peak.Peak_range = row.MustElement("td:nth-of-type(4)").MustText()
		peak.OpenToPublic = parseOpenToPublic(row.MustElement("td:nth-of-type(5)").MustText())

		// Append the populated mountain structure to the Peaks slice
		Peaks = append(Peaks, peak)
	}

	Peaks[0].Name = "Mount Everest"

	for _, peak := range Peaks {
		fmt.Printf("%+v\n", peak)
	}
}

// Helper function to parse the height string to a float32
func parseHeight(heightStr string) float32 {
	var height float32
	_, err := fmt.Sscanf(heightStr, "%f", &height)
	if err != nil {
		fmt.Println("Error parsing height:", err)
	}
	return height
}

// Helper function to parse the open-to-public status
func parseOpenToPublic(openStr string) bool {
	if openStr == "Opened" {
		return true
	} else {
		return false
	}
}
