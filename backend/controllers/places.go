package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/go-querystring/query"
)

type FindPlaceData struct {
	Input     string `url:"input"`
	InputType string `url:"inputtype"`
	Fields    string `url:"fields"`
	Key       string `url:"key"`
}

func SearchPlaces(c *gin.Context) {

	// query parameters to pass to google places api
	apiKey := "AIzaSyAeN85xIwoCXFUd3ZcSLLWVf0_LNPOm6Jo"
	params := FindPlaceData{
		Input:     c.Query("query"),
		InputType: "textquery",
		Fields:    "name,formatted_address,place_id",
		Key:       apiKey,
	}

	// establish url from encoded query parameters
	qvalues, _ := query.Values(params)
	url := "https://maps.googleapis.com/maps/api/place/findplacefromtext/json?" + qvalues.Encode()

	// create client and initialize request
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	// execute web request
	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}

	// read response from request
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// convert body bytes to json data
	var jsonMap map[string]interface{}
	json.Unmarshal(body, &jsonMap)

	// make sure response from google indicates success
	if jsonMap["status"] != "OK" {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    jsonMap["candidates"],
	})

}
