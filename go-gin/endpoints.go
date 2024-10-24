package gogin

import (
	"bytes"
	"encoding/json"
	"mountainsapi/scrapper"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GinAPI() {
	fileBytes, err := os.ReadFile("peaks.json")
	if err != nil {
		panic(err)
	}

	var Peaks []scrapper.Mountains

	err = json.Unmarshal(fileBytes, &Peaks)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/api/v1", func(c *gin.Context) {

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, fileBytes, "", "  "); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to format JSON"})

		}

		c.Data(http.StatusOK, "application/json", fileBytes)

	})

	r.GET("/api/v1/peak/:name", func(c *gin.Context) {
		var query = c.Param("name")
		if query == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name query parameter is required"})
		}

		var foundPeaks []scrapper.Mountains

		for _, peak := range Peaks {
			if query == peak.Name {
				foundPeaks = append(foundPeaks, peak)
			}
		}

		if len(foundPeaks) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Peak not found"})
		} else {

			prettyFoundPeaks, err := json.MarshalIndent(foundPeaks, "", "  ")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to format JSON"})
				return
			}
			c.Data(http.StatusOK, "application/json", prettyFoundPeaks)
		}
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
		}

		prettyFoundPeak, err := json.MarshalIndent(foundPeak, "", "  ")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to format JSON"})
			return
		}
		c.Data(http.StatusOK, "application/json", prettyFoundPeak)

	})

	r.Run("localhost:8080")
}
