package main

import (
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
