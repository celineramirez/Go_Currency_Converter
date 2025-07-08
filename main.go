package main

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/huh"
	"os"
	"strconv"
	"log"
)

var (
	convertFrom string
	convertTo   string
	amountStr   string
)

// Retrieve rates from the API and calculate the conversion
func convertRates(amount float64, currencyFrom string, currencyTo string) float64 {
	data := getConversion()

	// fmt.Printf("Date: %s, Base: %s\n", data.Date, data.Base)
	// fmt.Printf("Rates map: %+v\n", data.Rates)

	// fill from user input decided rate
	rateFrom := data.Rates[currencyFrom]
	rateTo := data.Rates[currencyTo]

	// Parse the string
	convFromRate, err := strconv.ParseFloat(rateFrom, 64)
	if err != nil {
		log.Fatalf("Error parsing rate from: %v", err)
	}

	convToRate, err2 := strconv.ParseFloat(rateTo, 64)
	if err2 != nil {
		log.Fatalf("Error parsing rate from: %v", err2)
	}

	// Calculate the conversion
	converted := amount * (convToRate / convFromRate)
	return converted
}

// Currency Menu
var currencySelect = []huh.Option[string]{
	huh.NewOption("USD", "USD"),
	huh.NewOption("GBP", "GBP"),
	huh.NewOption("EUR", "EUR"),
	huh.NewOption("PKR", "PKR"),
}

func main() {

	// Make TUI form
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What is your base currency?").
				Options(currencySelect...).
				Value(&convertFrom),
		),

		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What currency do you want to convert to?").
				Options(currencySelect...).
				Validate(func(x string) error {
					if x == convertFrom {
						return errors.New("cannot choose the same currency you are converting from")
					}
					return nil
				}).
				Value(&convertTo),
		),

		huh.NewGroup(
			huh.NewInput().
				Title("How much to convert?").
				Value(&amountStr).
				Validate(func(x string) error {
					amount, e := strconv.ParseFloat(x, 64)
					if e != nil || amount <= 0 {
						fmt.Errorf("please enter a numerical amount greater than 0")
					}
					return nil
				}),
		),
	)

	// Run the Terminal User Interface (TUI)
	err1 := form.Run()
	if err1 != nil {
		fmt.Println("Unable to run currency converter")
		os.Exit(1)
	}

	// Convert user input amount string to a float
	amount, e := strconv.ParseFloat(amountStr, 64)
	if e != nil {
		fmt.Printf("Invalid input %q, please enter a numerical value greater than 0\n", amountStr)
	}

	// Call the conversion function and get result
	converted := convertRates(amount, convertFrom, convertTo)
	result := fmt.Sprintf("%.2f", converted)

	switch convertTo {
	case "USD":
		result = "$" + result
	case "GBP":
		result = "£" + result
	case "EUR":
		result = "€" + result
	case "PKR":
		result = "Rs " + result
	}

	// Run form to display result
	resultForm := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("Result").
				CharLimit(400).
				Value(&result),
		),
	)

	// If unable to run resultForm
	err2 := resultForm.Run()
	if err2 != nil {
		fmt.Println("Unable to produce result")
		os.Exit(1)
	}

}
