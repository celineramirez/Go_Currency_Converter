package main

import (
	"fmt"
	"strconv"
	// "net/http"
	"github.com/charmbracelet/huh"
	// "encoding/json"
	"errors"
	"os"
)

var (
	base      string
	convertTo string
	rateFrom  int
	rateTo    int
	amount    string
	converted int
	result    string
)

func convertRates(convFrom int, currencyFrom string, currencyTo string) int {
	convFrom := data.Rates[currencyFrom]
	convTo := data.Rates[currencyTo]
	converted := amount * (convTo / convFrom)
	return converted
}

func main() {
	// apiKey := os.Getenv("CURRENCY_CONV_API_KEY")

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
				Value(&base),
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
					if x == base {
						return errors.New("cannot choose the same currency you are converting from")
					}

					return nil
				}).
				Value(&convertTo),
		),

		huh.NewGroup(
			huh.NewInput().
				Title("How much to convert?").
				Value(&amount).
				Validate(func(x string) error {
					if x <= 0 {
						return fmt.Errorf("please enter a valid amount of currency")
					}
					return nil
				}),
		),
	) // end of form

	// Do the currency conversion
	amount, e := strconv.Atoi(amount)
	if e == nil {
        fmt.Println(amount, e)
    }
	converted := convertRates(amount, base, convertTo)
	result := fmt.Sprintf("%d", converted)

	// Run the Terminal User Interface (TUI)
	err := form.Run()
	if err != nil {
		fmt.Println("Unable to run currency converter")
		os.Exit(1)
	}

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
