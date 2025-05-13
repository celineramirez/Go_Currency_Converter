package main
import (
	"fmt"
	// "net/http"
	// "github.com/charmbracelet/huh"
	// "encoding/json"
	"os"
)

var (
	base	 	string
	convert		string
	rateFrom	int
	rateTo		int
	converted	int
)

func main() {
	apiKey := os.Getenv("CURRENCY_CONV_API_KEY")
	
}