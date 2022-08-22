package config

type Config struct {
	ApiConfig *APIConfig
}

type APIConfig struct {
	ApiKey       string
	ApiSecret    string
	RequestToken string
	KiteLoginUrl string
	LoginId      string
	Passcode     string
	Zpin         int
}

func GetConfig() *Config {
	return &Config{
		ApiConfig: &APIConfig{
			ApiKey:       "exas5p7qo7tula9z",
			ApiSecret:    "741kozxapef490wabcokhs1d4giuhiw3",
			RequestToken: "",
			LoginId:      "",
			Passcode:     "",
			Zpin:         0,
		},
	}
}
