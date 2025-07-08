package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"log"
)

func getConversion() Currency {
	apiKey := os.Getenv("CURRENCY_CONV_API_KEY")

	url := "https://api.currencyfreaks.com/v2.0/rates/latest?apikey=" + apiKey + "&symbols=PKR,GBP,EUR,USD"
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// fmt.Println("Raw API Response:")
	// fmt.Println(
	// string(responseData))

	var currency Currency
	err = json.Unmarshal(responseData, &currency)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return currency
}
