package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Represents a resource, i.e. a URL that will be snapped
type Resource struct {
	Id             string   `json:"_id"`
	Url            string   `json:"url"`
	Email          string   `json:"email"`
	LatestSnapshot Snapshot `json:"latest_snapshot"`
}

// Take snapshot of resource
func (resource Resource) Snap() (*Snapshot, error) {
	log.Printf(`Snapping resource with url "%s"`, resource.Url)

	// Make HTTP call to API
	resp, err := http.Get(resource.Url)
	if err != nil {
		return nil, err
	}

	// Read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Create snapshot on heap
	snapshot := Snapshot{
		Datetime:   time.Now(),
		StatusCode: resp.StatusCode,
		Body:       string(body),
		ResourceId: resource.Id,
	}
	return &snapshot, nil
}

type Snapshot struct {
	Id         string    `json:"_id"`
	Datetime   time.Time `json:"datetime"`
	StatusCode int       `json:"status_code"`
	Body       string    `json:"body"`
	ResourceId string    `json:"resource_id"`
}

// ApiService to communicate with the resource DB api
type ApiService struct {
	host string
}

// Get list of all resources stored in DB
func (api ApiService) GetResources() ([]Resource, error) {
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
	return resources, err
}

// Post new snapshot
func (api ApiService) PostSnapshot(resource *Resource, snapshot *Snapshot) error {
	// Marshal snapshot
	body, err := json.Marshal(*snapshot)
	if err != nil {
		return err
	}

	// Post snapshot
	url := fmt.Sprintf("%s/resource/%s/snapshots", api.host, resource.Id)
	_, err = http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	log.Printf(`Successfully posted new snapshot to "%s"`, url)

	// Return nil if successful
	return nil
}
