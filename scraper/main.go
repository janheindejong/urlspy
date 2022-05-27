/*
Webscraper that takes all resources stored in Resource database, takes snapshots of URLs, and stores in snapshot DB
*/

package main

import (
	"log"
	"os"
	"sync"
	"time"
)

func main() {

	config := LoadConfigFromEnv()

	// Create services
	ResourceApiService := ResourceApiService{host: config.ResourceApi}
	SnapShotApiService := SnapShotApiService{host: config.SnapShotApi}

	waitTime, err := time.ParseDuration(config.WaitDuration)
	if err != nil {
		log.Fatalf(`Couldn't parse duration string "%s"`, getEnvWithDefault("APP_CONFIG_WAIT_TIME", "30s"))
	}
	for {
		// Get all Resources from resource api service
		resources, err := ResourceApiService.GetResources()
		if err != nil {
			log.Fatal(err.Error())
		}

		// For every resource, launch goroutine that scrapes and posts to snapshot service
		var wg sync.WaitGroup
		for _, resource := range *resources {
			wg.Add(1)
			resource := resource
			go SnapAndSave(&wg, &resource, &SnapShotApiService)
		}
		wg.Wait()
		time.Sleep(waitTime)
	}
}

// Takes snapshot of resource, and stores to snapshot database
func SnapAndSave(wg *sync.WaitGroup, resource *Resource, snapShotApiService *SnapShotApiService) {
	defer wg.Done()
	snapshot, err := resource.Snap()
	if err != nil {
		log.Printf(`Couldn't snap resource with url "%s", error: %s`, resource.Url, err.Error())
		return
	}
	log.Printf(`Received statuscode %v from resource with path "%s"`, snapshot.StatusCode, resource.Url)
	err = snapShotApiService.Create(snapshot)
	if err != nil {
		log.Printf(`Couldn't store snap of resource with url "%s", error: %s`, resource.Url, err.Error())
		return
	}
	log.Printf(`Successfully stored snapshot of resource with URL "%s"`, resource.Url)
}

type Config struct {
	ResourceApi  string
	SnapShotApi  string
	WaitDuration string
}

func LoadConfigFromEnv() *Config {
	config := Config{
		ResourceApi:  getEnvWithDefault("APP_CONFIG_RESOURCE_API", "http://resources"),
		SnapShotApi:  getEnvWithDefault("APP_CONFIG_SNAPSHOT_API", "http://snapshots"),
		WaitDuration: getEnvWithDefault("APP_CONFIG_WAIT_DURATION", "30s"),
	}
	log.Printf(`Loaded configuration: %+v`, config)
	return &config
}

func getEnvWithDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}
