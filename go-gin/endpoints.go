package gogin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mountainsapi/scrapper"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var Peaks []scrapper.Mountains

func loadPeaksFromFile() {
	fileBytes, err := os.ReadFile("peaks.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &Peaks)
	if err != nil {
		panic(err)
	}
}

func scrapeAndUpdatePeaks() {
	scrapper.ScrapePeaks(&Peaks)
	scrapper.WriteToJson(Peaks)
	loadPeaksFromFile()
	fmt.Println("Data updated")
}

func GinAPI() {

	scrapeAndUpdatePeaks()

	r := gin.Default()

	r.GET("/api/v1", func(c *gin.Context) {
		fileBytes, err := os.ReadFile("peaks.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read JSON file"})
			return
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, fileBytes, "", "  "); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to format JSON"})
			return
		}

		c.Data(http.StatusOK, "application/json", fileBytes)
	})

	r.GET("/api/v1/peak/:name", func(c *gin.Context) {
		query := c.Param("name")

		var foundPeaks []scrapper.Mountains
		for _, peak := range Peaks {
			if query == peak.Name {
				foundPeaks = append(foundPeaks, peak)
			}
		}

		if len(foundPeaks) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Peak not found"})
			return
		}

		prettyFoundPeaks, err := json.MarshalIndent(foundPeaks, "", "  ")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to format JSON"})
			return
		}
		c.Data(http.StatusOK, "application/json", prettyFoundPeaks)
	})

	r.GET("/api/v1/search", func(c *gin.Context) {
		idParam := c.Query("id")
		if idParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id query parameter is required"})
			return
		}

		peakID, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
			return
		}

		var foundPeak scrapper.Mountains
		var peakFound bool
		for _, peak := range Peaks {
			if peak.Peak_id == peakID {
				foundPeak = peak
				peakFound = true
				break
			}
		}

		if !peakFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Peak not found"})
			return
		}

		prettyFoundPeak, err := json.MarshalIndent(foundPeak, "", "  ")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to format JSON"})
			return
		}
		c.Data(http.StatusOK, "application/json", prettyFoundPeak)
	})

	go r.Run("localhost:8080")

	for {
		time.Sleep(20 * time.Hour)
		scrapeAndUpdatePeaks()
	}
}
