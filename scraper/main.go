/*
Webscraper
*/

package main

import (
	"encoding/json"
	"log"
	"net/http"
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
	}
}

// Snapshot service to communicate with that

type Resource struct {
	Url string `json:"url"`
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
