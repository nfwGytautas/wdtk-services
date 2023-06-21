package main

import (
	"errors"

	"github.com/nfwGytautas/wdtk-go-backend/microservice"
)

// Single entry inside locator table
type locatorEntry struct {
	ip string
}

// Table of location for services
type LocatorTable struct {
	entries map[string]locatorEntry
}

// Load locator table data from microservice config
func (lt *LocatorTable) Parse(config *microservice.MicroserviceConfig) error {
	lt.entries = make(map[string]locatorEntry)

	// Parse
	for _, element := range config.UserDefines["locatorTable"].([]interface{}) {
		entryMap := element.(map[string]interface{})

		entry := locatorEntry{
			ip: entryMap["ip"].(string),
		}
		lt.entries[entryMap["service"].(string)] = entry
	}

	return nil
}

// Get the ip of the service
func (lt *LocatorTable) GetIp(service string) (string, error) {
	entry, exists := lt.entries[service]
	if !exists {
		return "", errors.New("service " + service + " doesn't exist in the locator table")
	}
	return entry.ip, nil
}
