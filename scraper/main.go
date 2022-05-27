/*
Webscraper
*/

package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	println("Hello, world!")
	// Get all Resources from resource api service
	ResourceApiService := ResourceApiService{host: "http://resources"}
	resources, err := ResourceApiService.GetResources()
	if err != nil {
		log.Fatal(err.Error())
	}

	// For every resource, launch goroutine that scrapes and posts to snapshot service
	for _, resource := range *resources {
		println(resource.Url)
		snapshot, _ := resource.Snap()
		log.Println(*snapshot)
	}
}

// Snapshot service to communicate with that

type Resource struct {
	Url string `json:"url"`
}

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
