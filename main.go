package main

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/charmbracelet/huh"
	// "encoding/json"
	"errors"
	"os"
	"io/ioutil"
)

var (
	convertFrom string
	convertTo   string
	rateFrom    int
	rateTo      int
	amountStr   string
	converted   int
	result      string
)

func convertRates(amount int, currencyFrom string, currencyTo string) int {
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
					amount, e := strconv.Atoi(amountStr)
					if e == nil || amount <= 0 {
						fmt.Errorf("Please enter a numerical amount greater than 0")
					}
					return nil
				}),
		),
	) // end of form

	// Run the Terminal User Interface (TUI)
	err := form.Run()
	if err != nil {
		fmt.Println("Unable to run currency converter")
		os.Exit(1)
	}

	// Call currency conversion function
	amount, e := strconv.Atoi(amountStr)
	if e == nil {
		fmt.Println(amount, e)
	}

	converted := convertRates(amount, convertFrom, convertTo)
	result := fmt.Sprintf("%d", converted)

	// Run form to display result
	resultForm := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("Result").
				CharLimit(400).
				Value(&result),
		),
	) // end result form

	err1 := resultForm.Run()
	if err1 != nil {
		fmt.Println("Unable to produce result")
		os.Exit(1)
	}

}
