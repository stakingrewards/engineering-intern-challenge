package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Jeffail/gabs"

	"strconv"
	"strings"

	"github.com/go-rod/rod"
)

func main() {

	//send request
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/ardor")
	//check errors
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	//read responce
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//reaponce to json object
	jsonObj, err := gabs.ParseJSON(body)
	if err != nil {
		panic(err)
	}
	var price float64

	price = jsonObj.Path("market_data.current_price.usd").Data().(float64)

	page := rod.New().MustConnect().MustPage("https://www.ardorportal.org/monitor")
	fees := page.MustElement("#page > div.col-md-4 > div:nth-child(1) > div > table > tbody > tr:nth-child(5) > td").MustText()

	if feesInFloat, err := strconv.ParseFloat(strings.Fields(fees)[0], 64); err == nil {

		writeFile(annualEarnings(feesInFloat, price))
	}

}

func annualEarnings(fees float64, price float64) (formula float64) {
	formula = fees * 365 * price

	return
}

func writeFile(data float64) {

	f, err := os.OpenFile("history.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	output := fmt.Sprintf("%f \n", data)
	if _, err := f.WriteString(output); err != nil {
		log.Println(err)
	}
}
