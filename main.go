package main

import (
	"fmt"
	"strconv"
	// "net/http"
	"github.com/charmbracelet/huh"
	// "encoding/json"
	"os"
)

var (
	base      string
	convertTo string
	rateFrom  int
	rateTo    int
	converted int
)

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
						return errors.New("Cannot choose the same currency you are converting from")
					}

					return nil
				}).
				Value(&convertTo),
		),

		huh.NewGroup(
			huh.NewInput().
				Title("How much to convert?").
				Validate(func(x string) error {
					if _, err := strconv.ParseFloat(x, 64); err != nil {
						return fmt.Errorf("Please enter a valid value of currency")
					}
					return nil
				}),
			Value(&rateFrom),
		),
	) // end of form

	// Run the Terminal User Interface (TUI)
	err := form.Run()
	if err != nil {
		fmt.Println("Unable to run currency converter")
		os.Exit(1)
	}

}
