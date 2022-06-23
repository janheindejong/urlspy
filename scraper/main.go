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

	Config := LoadConfigFromEnv()

	// Create services
	ApiService := ApiService{host: Config.ApiHost}

	waitTime, err := time.ParseDuration(Config.WaitDuration)
	if err != nil {
		log.Fatalf(`Couldn't parse duration string "%s"`, Config.WaitDuration)
	}
	for {
		// Get all resources from api service
		resources, err := ApiService.GetResources()
		if err != nil {
			log.Fatal(err.Error())
		}

		// For every resource, launch goroutine that scrapes and posts to snapshot service
		// Note to self: I've investigated quite a bit how it works in Go with iterating over
		// pointers to slices and arrays. I came to the following conclusion:
		//
		// 1. You can range over a pointer to an array
		// 2. You can't range over a point to a slice
		// 3. Making a pointer to a slice is kind-of... well... pointless (punt intended),
		//    since a slice is already a pointer to an underlying array
		//
		// Case closed!

		var wg sync.WaitGroup
		for i := range resources {
			wg.Add(1)
			go SnapAndSave(&wg, &resources[i], &ApiService)
		}
		wg.Wait()
		time.Sleep(waitTime)
	}
}

// Takes snapshot of resource, and stores to snapshot database
func SnapAndSave(wg *sync.WaitGroup, resource *Resource, apiService *ApiService) error {
	defer wg.Done()

	// Make new snapshot
	snapshot, err := resource.Snap()
	if err != nil {
		log.Println(err)
		return err
	}

	// Compare
	if snapshot.Body == resource.LatestSnapshot.Body {
		log.Printf(`Body for url "%s" has not changed`, resource.Url)
	} else {
		log.Printf(`Body for url "%s" has changed`, resource.Url)
	}

	// Post
	err = apiService.PostSnapshot(resource, snapshot)
	if err != nil {
		log.Println(err)
	}

	return nil
}

type Config struct {
	ApiHost      string
	SnapShotApi  string
	WaitDuration string
}

func LoadConfigFromEnv() *Config {
	config := Config{
		ApiHost:      os.Getenv("APP_API_URL"),
		WaitDuration: os.Getenv("APP_WAIT_DURATION"),
	}
	log.Printf(`Loaded configuration: %+v`, config)
	return &config
}
