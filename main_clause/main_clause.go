package main_clause

import (
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

type DataResponse struct {
	Beef map[string]int `json:"beef"`
}

func fetchData() (string, error) {
	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
func cleanText(text string) []string {
	re := regexp.MustCompile(`[^a-zA-Z\s-]`)
	cleanedText := re.ReplaceAllString(text, "")
	return strings.Fields(strings.ToLower(cleanedText))
}

func countBeefTypes(text []string) map[string]int {
	beefTypes := []string{
		"t-bone", "fatback", "pastrami", "pork", "meatloaf", "jowl", "enim", "bresaola",
	}
	beefCount := make(map[string]int)
	for _, word := range text {
		for _, beef := range beefTypes {
			if word == beef {
				beefCount[beef]++
			}
		}
	}

	return beefCount
}
func getDataBeef(c *gin.Context) {

	data, err := fetchData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	cleanedText := cleanText(data)
	beefCount := countBeefTypes(cleanedText)
	c.JSON(http.StatusOK, DataResponse{
		Beef: beefCount,
	})
}

func Server() {
	router := gin.Default()
	router.GET("/food/mainclause", getDataBeef)
	if err := router.Run(":3000"); err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
