package main

import (
	"fmt"
)

var resultMessage string

func convertKelvin(temperatureInput float64) (float64, float64) {
	//Insert your code here
	resultFahrenheit := ((9.0 / 5.0) * temperatureInput) - 459.67
	resultCelsius := (5.0 / 9.0) * (resultFahrenheit - 32.0)
	//Do not remove this line
	return resultFahrenheit, resultCelsius
}

func convertCelsius(temperatureInput float64) (float64, float64) {
	//Insert your code here
	resultFahrenheit := (9.0/5.0)*(temperatureInput) + 32.0
	resultKelvin := (5.0 / 9.0) * (resultFahrenheit + 459.67)

	//Do not remove this line
	return resultFahrenheit, resultKelvin
}

func convertFahrenheit(temperatureInput float64) (float64, float64) {
	//Insert your code here
	resultKelvin := (5.0 / 9.0) * (temperatureInput + 459.67)
	resultCelsius := (5.0 / 9.0) * (temperatureInput - 32.0)

	//Do not remove this line
	return resultKelvin, resultCelsius
}

func main() {
	var temperatureChoice int
	var temperatureInput float64
	fmt.Println("Enter your temperature format :1 for Kelvin, 2 for Celsius, 3 for Fahrenheit: ")
	fmt.Scanln(&temperatureChoice)
	fmt.Println("Enter the temperature: ")
	fmt.Scanln(&temperatureInput)

	if temperatureChoice == 1 {
		//Insert Code here

		resultFahrenheit, resultCelsius := convertKelvin(temperatureInput)

		// resultMessage = "Temperature in farenheit is " + strconv.FormatFloat(output1, 'f', 2, 64) + " and temperature in celsius is " + strconv.FormatFloat(output2, 'f', 2, 64)

		//DO not remove this
		fmt.Printf("Fahrenheit: %.02f, Celsius: %.02f", resultFahrenheit, resultCelsius)

	} else if temperatureChoice == 2 {
		//Insert Code here

		resultFahrenheit, resultKelvin := convertCelsius(temperatureInput)

		//DO not remove this
		fmt.Printf("Fahrenheit: %.2f, Celcius: %.02f", resultFahrenheit, resultKelvin)

	} else if temperatureChoice == 3 {
		//Insert Code here

		resultKelvin, resultCelsius := convertFahrenheit(temperatureInput)

		//DO not remove this
		fmt.Printf("Kelvin: %.2f, Celcius %.02f", resultKelvin, resultCelsius)

	} else {
		fmt.Println("No Conversion")
	}

}
