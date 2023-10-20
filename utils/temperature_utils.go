// Utility function to handle conversion between Celsius and Fahrenheit.

package utils

func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}
