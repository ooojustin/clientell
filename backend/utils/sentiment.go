package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type SentimentDocument struct {
	ID       int    `json:"id"`
	Language string `json:"language"`
	Text     string `json:"text"`
}

type SentimentRequest struct {
	Documents []SentimentDocument `json:"documents"`
}

type ConfidenceScores struct {
	Positive float64 `json:"positive"`
	Neutral  float64 `json:"neutral"`
	Negative float64 `json:"negative"`
}

type Sentiment struct {
	Sentiment        string           `json:"sentiment"`
	ConfidenceScores ConfidenceScores `json:"confidenceScores"`
}

type SentimentResponse struct {
	Documents []Sentiment `json:"documents"`
}

func AnalyzeSentiment(text string) (*Sentiment, error) {

	url := "https://clientell.cognitiveservices.azure.com/text/analytics/v3.0/sentiment"
	apiKey := "910496a9030c4c5596bafa5d3058f115"

	// determine data to send to server
	sreq := SentimentRequest{
		Documents: []SentimentDocument{
			SentimentDocument{
				ID:       1,
				Language: "en",
				Text:     text,
			},
		},
	}
	bjson, _ := json.Marshal(sreq)

	// create client and initialize request
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bjson))

	// set api key in header
	req.Header.Set("Ocp-Apim-Subscription-Key", apiKey)

	// execute web request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// read response from request
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// deserialize response data
	var sresp SentimentResponse
	if err := json.Unmarshal(body, &sresp); err != nil {
		return nil, err
	}

	// return first document (resolved sentiment data) if provided
	if len(sresp.Documents) > 0 {
		sentiment := &sresp.Documents[0]
		return sentiment, nil
	} else {
		return nil, errors.New("Failed to analyze sentiment of text.")
	}

}
