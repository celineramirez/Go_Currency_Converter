package main

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"strconv"
	"errors"
	"os"
)

var (
	convertFrom string
	convertTo   string
	amountStr   string
)

func convertRates(amount float64, currencyFrom string, currencyTo string) float64 {
	data := getConversion()
	convFrom := data.Rates[currencyFrom]
	convTo := data.Rates[currencyTo]
	converted := amount * (convTo / convFrom)
	return converted
}

func main() {

	// Make TUI form
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What is your base currency?").
				Options(
					huh.NewOption("USD", "usd"),
					huh.NewOption("GBP", "gbp"),
					huh.NewOption("EUR", "eur"),
					huh.NewOption("JPY", "jpy"),
				).
				Value(&convertFrom),
		),

		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What currency do you want to convert to?").
				Options(
					huh.NewOption("USD", "usd"),
					huh.NewOption("GBP", "gbp"),
					huh.NewOption("EUR", "eur"),
					huh.NewOption("JPY", "jpy"),
				).
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
					amount, e := strconv.ParseFloat(amountStr,64)
					if e == nil || amount <= 0 {
						fmt.Errorf("Please enter a numerical amount greater than 0")
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

	// Call currency conversion function
	amount, e := strconv.ParseFloat(amountStr,64)
	if e == nil {
		fmt.Println(amount, e)
	}

	converted := convertRates(amount, convertFrom, convertTo)
	result := fmt.Sprintf("%.2f", converted)

	// Run form to display result
	resultForm := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("Result").
				CharLimit(400).
				Value(&result),
		),
	)

	err2 := resultForm.Run()
	if err2 != nil {
		fmt.Println("Unable to produce result")
		os.Exit(1)
	}

}
