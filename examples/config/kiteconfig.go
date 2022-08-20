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
			ApiKey:       "",
			ApiSecret:    "",
			RequestToken: "",
			LoginId:      "",
			Passcode:     "",
			Zpin:         0,
		},
	}
}
