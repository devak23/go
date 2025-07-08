package main

import "fmt"

type Scalar interface {
	~float32 | ~float64
}

func convertCelsiusToFahrenheit[T Scalar](fromTemp T) T {
	return fromTemp*T(9)/T(5) + T(32)
}

func convertFahrenheitToCelsius[T Scalar](fromTemp T) T {
	return (fromTemp - T(32)) * T(5) / T(9)
}

func convertCelsiusToKelvin[T Scalar](fromTemp T) T {
	return fromTemp + T(273.15)
}

func convertKelvinToCelsius[T Scalar](fromTemp T) T {
	return fromTemp - T(273.15)
}

func convertKelvinToFahrenheit[T Scalar](fromTemp T) T {
	return (fromTemp-T(273.15))*T(9)/T(5) + T(32)
}

func convertFahrenheitToKelvin[T Scalar](fromTemp T) T {
	return (fromTemp-T(32))*T(5)/T(9) + T(273.15)
}

type doConvert[T Scalar] func(T) T

func getConversions[T Scalar]() map[string]doConvert[T] {
	return map[string]doConvert[T]{
		"CF": convertCelsiusToFahrenheit[T],
		"FC": convertFahrenheitToCelsius[T],
		"CK": convertCelsiusToKelvin[T],
		"KC": convertKelvinToCelsius[T],
		"KF": convertKelvinToFahrenheit[T],
		"FK": convertFahrenheitToKelvin[T],
	}
}

func convert[T Scalar](a T, op string) T {

	if op, ok := getConversions[T]()[op]; ok {
		return op(a)
	} else {
		panic(fmt.Sprintf("Unsupported operation %s", op))
	}
}

func main() {

	fmt.Println("Fahrenheit = ", convert(91.0, "CF"))
	fmt.Println("Celsius = ", convert(195.8, "FC"))
	fmt.Println("Kelvin = ", convert(91.0, "FK"))
}
