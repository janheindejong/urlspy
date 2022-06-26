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
	dir := os.Getenv("APP_SECRET_DIR")
	b, err := os.ReadFile(dir + "/email_account.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func getEmailPassword() string {
	dir := os.Getenv("APP_SECRET_DIR")
	b, err := os.ReadFile(dir + "/email_password.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
