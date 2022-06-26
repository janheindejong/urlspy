package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// DbService to communicate with the resource DB api
type DbService struct {
	host string
}

// Get list of all resources stored in DB
func (db DbService) GetResources() ([]Resource, error) {
	// Gets all the resources in the database
	log.Printf(`Getting resources from database at "%v"`, db.host)

	// Make HTTP call to API
	resp, err := http.Get(db.host + "/resource")
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
func (db DbService) PostSnapshot(resource *Resource, snapshot *Snapshot) error {
	// Marshal snapshot
	body, err := json.Marshal(*snapshot)
	if err != nil {
		return err
	}

	// Post snapshot
	url := fmt.Sprintf("%s/resource/%s/snapshots", db.host, resource.Id)
	_, err = http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	log.Printf(`Successfully posted new snapshot to "%s"`, url)

	// Return nil if successful
	return nil
}
