package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Jawadh-Salih/gn-lk-api/types"
	"github.com/gin-gonic/gin"
)

// pingHandler handles the /ping route
func pingHandler(c *gin.Context) {
	jsonFile, err := os.Open("gn-list.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	fmt.Println(string(byteValue))

	// Unmarshal the JSON data into a slice of GNDivisions
	var divisions []types.GNDivision
	err = json.Unmarshal(byteValue, &divisions)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "pong",
		"divisions": divisions, // Include the divisions data in the response
	})
}

func main() {
	// Read the JSON file

	router := gin.Default()

	// Define your API routes here
	router.GET("/ping", pingHandler)

	router.Run(":8080")
}
