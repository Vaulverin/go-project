package main

import "fmt"

func getParams() (string, float) {

}

func convertCurrencyValue(currency string, value float) (string, value) {

}

func main() {
  currency, value := getParams()
  convertedCurrency, convertedValue := convertCurrencyValue(currency, value)
  fmt.Println(convertedCurrency, " = ", convertedValue)
}
