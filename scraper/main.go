package main

import (
	"time"
)

func main() {

	config := LoadConfigFromEnv()
	apiService := ApiService{host: config.ApiHost}

	for {
		RunOnce(&apiService)
		time.Sleep(config.WaitDuration)
	}
}
