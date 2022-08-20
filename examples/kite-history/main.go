package main

import (
	"fmt"
	"gnKiteConnect/examples/config"
	"strconv"
	"time"

	kiteconnect "github.com/zerodhatech/gokiteconnect"
)

func main() {
	kc := kiteconnect.New(config.GetConfig().ApiConfig.ApiKey)

	// Login URL from which request token can be obtained
	fmt.Println(kc.GetLoginURL())

	// Obtained request token after Kite Connect login flow
	// simulated here by scanning from stdin
	var requestToken string
	fmt.Println("Enter request token:")
	fmt.Scanf("%s\n", &requestToken)

	config.GetConfig().ApiConfig.RequestToken = requestToken
	// Get user details and access token
	session, err := kc.GenerateSession(requestToken, config.GetConfig().ApiConfig.ApiSecret)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	kc.SetAccessToken(session.AccessToken)
	fromDate := "2022-08-19"
	toDate := "2022-08-20"

	data, err := GetHistoryData(kc, 256265 /*NIFTY*/, fromDate, toDate, 5)

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
