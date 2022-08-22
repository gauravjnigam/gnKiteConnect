package main

import (
	"fmt"
	"gnKiteConnect/examples/setup"
	"strconv"
	"time"

	kiteconnect "github.com/zerodhatech/gokiteconnect"
)

func main() {

	kc := setup.SetupKiteConnection()

	fromDate := "2022-08-19"
	toDate := "2022-08-20"

	data, err := GetHistoryData(kc, 256265 /*NIFTY*/, fromDate, toDate, 2)
	if err != nil {
		fmt.Printf("Error while getting history data - %v", err)
	}

	fmt.Printf("%v\n", data)
}

func GetHistoryData(kc *kiteconnect.Client, instrumentCode int, fromDate string,
	toDate string /*YYYY-MM-DD*/, intervalInMin int) ([]kiteconnect.HistoricalData, error) {
	fmt.Printf("Calling history data from %s - %s\n", fromDate, toDate)

	from_date := fromDate
	to_date := toDate
	fromDateTime, err := time.Parse("2006-01-02", from_date)
	toDateTime, err := time.Parse("2006-01-02", to_date)
	if err != nil {
		fmt.Errorf("Error while fetching history data - %v\n", err)
	}

	interval := strconv.Itoa(intervalInMin) + "minute"
	fmt.Printf("Interval : %s\n", interval)

	return kc.GetHistoricalData(instrumentCode, interval, fromDateTime, toDateTime, false)
}
