package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// pingHandler handles the /ping route
func pingHandler(c *gin.Context, divisions []GNDivision) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "pong",
		"divisions": divisions, // Include the divisions data in the response
	})
}

func main() {
	// Read the JSON file
	jsonFile, err := os.Open("gn-list.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Unmarshal the JSON data into a slice of GNDivisions
	var divisions []GNDivision
	json.Unmarshal(byteValue, &divisions)

	router := gin.Default()

	// Define your API routes here
	router.GET("/ping", func(c *gin.Context) {
		pingHandler(c, divisions) // Pass divisions to pingHandler
	})

	router.Run(":8080")
}
