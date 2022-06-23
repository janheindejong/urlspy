package main

import (
	"log"
	"sync"
)

// Main function, that is repeated in intervals of WaitDuration
// Note to self: I've investigated quite a bit how it works in Go with iterating over
// pointers to slices and arrays. I came to the following conclusion:
//
// 1. You can range over a pointer to an array
// 2. You can't range over a point to a slice
// 3. Making a pointer to a slice is kind-of... well... pointless (punt intended),
//    since a slice is already a pointer to an underlying array
func RunOnce(apiService *ApiService) {

	// Get all resources from api service
	resources, err := apiService.GetResources()
	if err != nil {
		log.Fatal(err.Error())
	}

	// For every resource, launch goroutine that scrapes and posts to snapshot service
	var wg sync.WaitGroup
	for i := range resources {
		wg.Add(1)
		go SnapAndSave(&wg, &resources[i], apiService)
	}
	wg.Wait()
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

	// Compare to see if anything has changed
	if snapshot.Body == resource.LatestSnapshot.Body {
		log.Printf(`Body for url "%s" has not changed`, resource.Url)
	} else {
		log.Printf(`Body for url "%s" has changed`, resource.Url)
	}

	// Post snapshot to API
	err = apiService.PostSnapshot(resource, snapshot)
	if err != nil {
		log.Println(err)
	}

	return nil
}
