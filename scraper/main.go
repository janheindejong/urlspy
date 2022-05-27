/*
Webscraper that takes all resources stored in Resource database, takes snapshots of URLs, and stores in snapshot DB
*/

package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	// Get all Resources from resource api service
	ResourceApiService := ResourceApiService{host: "http://resources"}
	resources, err := ResourceApiService.GetResources()
	if err != nil {
		log.Fatal(err.Error())
	}

	// For every resource, launch goroutine that scrapes and posts to snapshot service
	var wg sync.WaitGroup
	for _, resource := range *resources {
		wg.Add(1)
		resource := resource
		go SnapAndSave(&wg, &resource)
	}
	wg.Wait()
}

// Takes snapshot of resource, and stores to snapshot database
func SnapAndSave(wg *sync.WaitGroup, resource *Resource) {
	defer wg.Done()
	snapshot, err := resource.Snap()
	if err != nil {
		log.Printf(`Couldn't snap resource with url "%s", error: %s`, resource.Url, err.Error())
		return
	}
	log.Printf(`Received statuscode "%v" from resource with path "%s"`, snapshot.StatusCode, resource.Url)
}

// Represents a resource, i.e. a URL that will be snapped
type Resource struct {
	Url string `json:"url"`
}

// Take snapshot of resource
func (resource Resource) Snap() (*SnapShot, error) {
	log.Printf(`Snapping resource with url "%s"`, resource.Url)

	// Make HTTP call to API
	resp, err := http.Get(resource.Url)
	if err != nil {
		return nil, err
	}

	// Read body
	body, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		return nil, err
	}

	// Create snapshot on heap
	snapshot := SnapShot{
		Url:        resource.Url,
		Datetime:   time.Now(),
		StatusCode: resp.StatusCode,
		Body:       string(body),
	}
	return &snapshot, nil
}

// Snapshot of given URL
type SnapShot struct {
	Url        string    `json:"url"`
	Datetime   time.Time `json:"datetime"`
	StatusCode int       `json:"response"`
	Body       string    `json:"body"`
}

// ResourceApiService to communicate with the resource DB api
type ResourceApiService struct {
	host string
}

// Get list of all resources stored in DB
func (api ResourceApiService) GetResources() (*[]Resource, error) {
	// Gets all the resources in the database
	log.Printf(`Getting resources from database at "%v"`, api.host)

	// Make HTTP call to API
	resp, err := http.Get(api.host + "/resource")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Parse response
	var resources []Resource
	err = json.NewDecoder(resp.Body).Decode(&resources)
	return &resources, err
}
