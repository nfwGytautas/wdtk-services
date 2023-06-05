package forward

import (
	"encoding/json"
	"os"
)

// PUBLIC TYPES
// ========================================================================

// Single entry inside locator table
type LocatorEntry struct {
	ServiceName string `json:"Service"`
	IP          string `json:"IP"`
}

// Table of location for services
type LocatorTable struct {
	Entries []LocatorEntry `json:"Mapping"`
}

// PRIVATE TYPES
// ========================================================================

// PUBLIC FUNCTIONS
// ========================================================================

// Read a locator table
func LoadLocatorTable() (LocatorTable, error) {
	result := LocatorTable{}

	data, err := os.ReadFile("LocatorTable.json")
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

// Returns a mapping for a specified service, empty string if can't map
func (lt *LocatorTable) Map(service string) string {
	for _, entry := range lt.Entries {
		if entry.ServiceName == service {
			return entry.IP
		}
	}

	return ""
}

// PRIVATE FUNCTIONS
// ========================================================================
