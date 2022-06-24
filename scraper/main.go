package main

func main() {

	config := LoadConfigFromEnv()

	scraper := Scraper{
		apiService:   &ApiService{host: config.ApiHost},
		waitDuration: config.WaitDuration,
	}

	scraper.RunForever()

}
