package main

import (
	"log"
	"sync"
	"time"
)

// wg is used to ensure all resources have been handled before starting a new loop
var wg sync.WaitGroup

// Scraper is responsible for the business logic of the application -
// i.e. periodically taking snapshots from all resources in the DB
type Scraper struct {
	apiService   *ApiService
	emailService *EmailServiceSmtp
	waitDuration time.Duration
}

// RunForever starts the scraper
func (scraper Scraper) RunForever() {
	for {
		scraper.runOnce()
		log.Println("Waiting")
		time.Sleep(scraper.waitDuration)
	}
}

// runOnce takes and stores snapshots of all the resources in the DB
func (scraper Scraper) runOnce() {

	// Get all resources from api service
	resources, err := scraper.apiService.GetResources()
	if err != nil {
		log.Fatal(err.Error())
	}

	// For every resource, launch goroutine that scrapes and posts to snapshot service
	for i := range resources {
		wg.Add(1)
		go scraper.handleResource(&resources[i])
	}
	wg.Wait()

	// Note to self: I've investigated quite a bit how it works in Go with iterating over
	// pointers to slices and arrays. I came to the following conclusion:
	//
	// 1. You can range over a pointer to an array
	// 2. You can't range over a point to a slice
	// 3. Making a pointer to a slice is kind-of... well... pointless (punt intended),
	//    since a slice is already a pointer to an underlying array
}

// HandleResource stores a new snapshot of a resource, and verifies if the resource has changed
func (scraper Scraper) handleResource(resource *Resource) (err error) {
	defer wg.Done()

	// Make new snapshot
	snapshot, err := resource.Snap()
	if err != nil {
		log.Println(err)
		return
	}

	// Post snapshot to API
	err = scraper.apiService.PostSnapshot(resource, snapshot)
	if err != nil {
		log.Println(err)
	}

	// Send notification e-mail if resource has changed
	if resource.HasChanged(snapshot) {
		err = scraper.emailService.SendEmailResourceChanged(resource)
		if err != nil {
			log.Printf(
				`Resource with URL "%s" has changed, but received error while sending email to "%s": %s`,
				resource.Url, resource.Email, err,
			)
		} else {
			log.Printf(`Sent email to "%s", notifying of change in URL: "%s"`, resource.Email, resource.Url)
		}
	} else {
		log.Printf(`Body for url "%s" has not changed`, resource.Url)
	}

	return
}