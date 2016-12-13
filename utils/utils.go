package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// SwitchCaseError creates and returns a switch case error.
func SwitchCaseError(name string, value interface{}) error {
	return fmt.Errorf("Unrecognized %s: %v", name, value)
}

// Assert panics with message if the condition is false.
func Assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

// AssertIsAPIKey asserts that is was supplied a correctly formatted API key
func AssertIsAPIKey(apiKey string) {
	Assert(isValidAirtableID(apiKey, "key"), "invalid API Key encountered")
}

// AssertIsBaseID asserts that is was supplied a correctly formatted base ID
func AssertIsBaseID(baseID string) {
	Assert(isValidAirtableID(baseID, "app"), "invalid base ID encountered")
}

// AssertIsRecordID asserts that is was supplied a correctly formatted record ID
func AssertIsRecordID(recordID string) {
	Assert(isValidAirtableID(recordID, "rec"), "invalid record ID encountered")
}

func isValidAirtableID(id, expectedPrefix string) bool {
	regex := regexp.MustCompile("[a-zA-Z0-9]{17}")
	return len(id) == 17 && strings.HasPrefix(id, expectedPrefix) && regex.Match([]byte(id))
}
