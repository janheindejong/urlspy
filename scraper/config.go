package main

import (
	"log"
	"os"
	"time"
)

type Config struct {
	ApiHost       string
	SnapShotApi   string
	WaitDuration  time.Duration
	EmailAccount  string
	EmailPassword string
}

func LoadConfigFromEnv() *Config {
	config := Config{
		ApiHost:       getApiHost(),
		WaitDuration:  getWaitDuration(),
		EmailAccount:  getEmailAcount(),
		EmailPassword: getEmailPassword(),
	}
	log.Printf(`Loaded configuration: %+v`, config)
	return &config
}

func getApiHost() string {
	s := os.Getenv("APP_API_URL")
	if s == "" {
		log.Fatal(`Environment variable "APP_API_URL" not set`)
	}
	return s
}

func getWaitDuration() time.Duration {
	s := os.Getenv("APP_WAIT_DURATION")
	duration, err := time.ParseDuration(s)
	if err != nil {
		log.Fatalf(`Couldn't parse environment variable APP_WAIT_DURATION "%s"`, s)
	}
	return duration
}

func getEmailAcount() string {
	filename := os.Getenv("APP_EMAIL_ACCOUNT_FILE")
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func getEmailPassword() string {
	filename := os.Getenv("APP_EMAIL_PASSWORD_FILE")
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
