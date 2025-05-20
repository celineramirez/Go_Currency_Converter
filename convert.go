package main

import (
	"encoding/json"
	"fmt"
	// "io"
	// "log"
	"net/http"
	"os"
)

func getConversion() Response {
	apiKey := os.Getenv("CURRENCY_CONV_API_KEY")

	// api call
	url := "https://api.currencyfreaks.com/"
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Print(err.Error())
    }

	req.Header.Add("cf-api-key", apiKey)

	res, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Print(err.Error())
    }

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
        fmt.Print(err.Error())
    }
    fmt.Println(string(body))
}