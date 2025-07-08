# Go_Currency_Converter

This is a currency converter in Golang that retrieves data rate values from an API: https://currencyfreaks.com/<br>
The charmbracelet/huh TUI framework is implemented for displaying the menu options for the currency converter: https://github.com/charmbracelet/huh<br>

Upon running the program, the user is prompted to input the currency they will be converting from:<br>
![base currency](screenshots/base_cur.png)

After entering the base currency, the user is prompted to input the currency they want to convert to:<br>
![conversion currency](screenshots/convert_to.png)

However, if the currency that the user wants to convert to is the same as their base currency, they will see the following error message:<br>
![same currency error](screenshots/same_currency.png)

If the user correctly inputs a different currency to convert to, they will be prompted to enter the amount they would like to convert from their base currency:<br>
![enter value](screenshots/enter_Val.png)

If the user inputs an invalid value (non-numerical or numerical value less than or equal to 0), they will see the following error message:<br>
![invalid input error](screenshots/invalid_input.png)<br>
![invalid input error](screenshots/invalid_input2.png)

If a valid numerical value is entered, the user will get their converted currency value:<br>
![result](screenshots/result.png)
