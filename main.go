package main
import (
	"fmt"
	// "net/http"
	"github.com/charmbracelet/huh"
	// "encoding/json"
	"os"
)

var (
	base	 	string
	convertTo	string
	rateFrom	int
	rateTo		int
	converted	int
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
	)

	huh.NewGroup(
	huh.NewSelect[string]().
	Title("What currency do you want to convert to?").
	Options(
			huh.NewOption("USD", "usd"),
			huh.NewOption("GBP", "gbp"),
			huh.NewOption("EUR", "eur"),
			huh.NewOption("JPY", "jpy"),
		).
	Value(&convertTo),

	huh.NewInput().
		Title("How much to convert?").
		Value(&rateFrom),
		Validate(func(amt int) error{
			if amt <= 0 {
                return fmt.Errorf("Please enter a value greater than 0")
            }
            return nil
		}),
	)

	// Run the Terminal User Interface (TUI)
	err := form.Run()
	if err != nil {
		fmt.Println("Unable to run currency converter")
		os.Exit(1)
	}

}