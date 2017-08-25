package main

import (
  "flag"
  "fmt"
  "os"
  "net/http"
  "io/ioutil"
  "regexp"
  "strconv"
)
func getParams() (string, float64) {
  listCommand := flag.NewFlagSet("list", flag.ExitOnError)
  // List subcommand flag pointers
  currencyPtr := listCommand.String("currency", "", "Currency - RUB or EUR. (Required)")
  valuePtr := listCommand.Float64("value", 1, "Value to parse, can be float. (Required)")
  listCommand.Parse(os.Args[1:])
  if listCommand.Parsed() == false {
    fmt.Println("No arguments.")
    os.Exit(1)
  }
  if *currencyPtr != "EUR" && *currencyPtr != "RUB" {
	listCommand.PrintDefaults()
    os.Exit(1)
  }
  return *currencyPtr, *valuePtr
}

func convertCurrencyValue(currency string, value float64) (string, float64) {
  resp, err := http.Get("http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
  if err != nil {
    fmt.Errorf("GET error: %v", err)
    os.Exit(1)
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    fmt.Errorf("Status error: %v", resp.StatusCode)
    os.Exit(1)
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Errorf("Read body: %v", err)
	os.Exit(1)
  }
  re := regexp.MustCompile("RUB.+rate='([0-9\\.]+)'")
  matches:= re.FindStringSubmatch(string(data))
  if len(matches) < 2 {
	fmt.Errorf("No RUB currency")
	os.Exit(1)
  }
  rate, err := strconv.ParseFloat(matches[1], 64)
  if err != nil {
    fmt.Errorf("Parse error. String to float")
	os.Exit(1)
  }
  if currency == "EUR" {
	return "RUB", value * rate
  }
  return "EUR", value / rate
}

func main() {
  currency, value := getParams()
  convertedCurrency, convertedValue := convertCurrencyValue(currency, value)
  fmt.Println(convertedCurrency, " = ", convertedValue)
}
