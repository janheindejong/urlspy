package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	ApiHost       string
	SnapShotApi   string
	WaitDuration  time.Duration
	EmailHost     string
	EmailPort     int
	EmailAccount  string
	EmailPassword string
}

func LoadConfigFromEnv() *Config {
	config := Config{
		ApiHost:       getenv("APP_API_URL"),
		WaitDuration:  getWaitDuration(),
		EmailHost:     getenv("APP_EMAIL_HOST"),
		EmailPort:     getEmailPort(),
		EmailAccount:  getEmailAcount(),
		EmailPassword: getEmailPassword(),
	}
	log.Printf(`Loaded configuration: %+v`, config)
	return &config
}

func getWaitDuration() time.Duration {
	s := getenv("APP_WAIT_DURATION")
	duration, err := time.ParseDuration(s)
	if err != nil {
		log.Fatalf(`Couldn't parse environment variable APP_WAIT_DURATION "%s"`, s)
	}
	return duration
}

func getEmailAcount() string {
	filename := getenv("APP_EMAIL_ACCOUNT_FILE")
	return readSecret(filename)
}

func getEmailPassword() string {
	filename := getenv("APP_EMAIL_PASSWORD_FILE")
	return readSecret(filename)
}

func getEmailPort() int {
	s := getenv("APP_EMAIL_PORT")
	port, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Couldn't parse port %s", s)
	}
	return port
}

func getenv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("Environment variable %s not set", name)
	}
	return value
}

func readSecret(filename string) string {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Need to trim any trailing whitespace, since Docker secrets adds a newline to the file
	return strings.TrimSpace(string(b))
}
