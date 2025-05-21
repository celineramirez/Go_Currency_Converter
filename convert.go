package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getConversion() Currency {
	apiKey := os.Getenv("CURRENCY_CONV_API_KEY")

	url := "https://api.currencyfreaks.com/"
	response, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	response.Header.Add("cf-api-key", apiKey)

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var currency Currency
	err = json.Unmarshal(responseData, &currency)
	if err != nil {
		fmt.Println("Unable to unmarshal response")
	}

	return currency
}
