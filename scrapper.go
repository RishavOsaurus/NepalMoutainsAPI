package main

import (
	"fmt"

	"github.com/go-rod/rod"
)

type Mountains struct {
	Peak_id      int
	Alias        string
	Height       float32
	Peak_range   string
	OpenToPublic bool
}

func main() {
	browser := rod.New().NoDefaultDevice().MustConnect()
	defer browser.MustClose()
	page := browser.MustPage("https://nepalhimalpeakprofile.org/peak-profile")
	page.MustWindowFullscreen()
	page.MustWaitLoad()

	rows := page.MustElements("table#mountaintable tbody tr")
	for _, row := range rows {
		id := row.MustElement("td:nth-of-type(1)").MustText()
		fmt.Println(id)

	}

}
