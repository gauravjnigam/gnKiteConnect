package setup

import (
	"fmt"
	"gnKiteConnect/examples/config"

	kiteconnect "github.com/zerodhatech/gokiteconnect"
)

func SetupKiteConnection() *kiteconnect.Client {
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
		return nil
	}

	kc.SetAccessToken(session.AccessToken)

	return kc
}
