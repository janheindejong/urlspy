package main

func main() {

	config := LoadConfigFromEnv()
	apiService := ApiService{host: config.ApiHost}
	emailService := EmailServiceSmtp{
		host:    "smtp.gmail.com",
		port:    587,
		account: config.EmailAccount,
		pass:    config.EmailPassword,
	}

	scraper := Scraper{
		apiService:   &apiService,
		emailService: &emailService,
		waitDuration: config.WaitDuration,
	}

	scraper.RunForever()

}
