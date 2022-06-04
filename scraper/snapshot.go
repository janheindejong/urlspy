package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// Snapshot of given URL
type SnapShot struct {
	Url        string    `json:"url"`
	Datetime   time.Time `json:"datetime"`
	StatusCode int       `json:"response"`
	Body       string    `json:"body"`
}

// SnapShotApiService is utility class to communicate with snapshot DB
type SnapShotApiService struct {
	host string
}

func (api SnapShotApiService) Create(snapshot *SnapShot) error {

	// Marshal snapshot
	body, err := json.Marshal(*snapshot)
	if err != nil {
		return err
	}

	// Post snapshot
	_, err = http.Post(
		api.host+"/snapshot",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	// Return nil if successful
	return nil
}
