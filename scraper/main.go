/*
Webscraper that takes all resources stored in Resource database, takes snapshots of URLs, and stores in snapshot DB
*/

package main

import (
	"log"
	"sync"
)

func main() {
	// Create services
	ResourceApiService := ResourceApiService{host: "http://resources"}
	SnapShotApiService := SnapShotApiService{host: "http://snapshots"}

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
