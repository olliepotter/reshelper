package reshelper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ErrorMessage defines a structure for handling errors
type ErrorMessage struct {
	AppCode  string `json:"app_code"`
	Message  string `json:"message"`
	Remedies string `json:"potential_remedies"`
}

// Message defines a format to send JSON messages
type Message struct {
	Status   int    `json:"status_code"`
	AppCode  string `json:"app_code"`
	Message  string `json:"message"`
	Remedies string `json:"potential_remedies"`
}

// HardError sends 500 and logs a server error to stdout
func HardError(w http.ResponseWriter, em ErrorMessage) {

	// Send response
	w.WriteHeader(http.StatusInternalServerError)

	// Log
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s :: %d [%s] %s\n", timestamp, http.StatusInternalServerError, em.AppCode, em.Message)

}

// SendMessage sends informational response message and logs it
func SendMessage(w http.ResponseWriter, rm Message) {

	// Send response
	w.WriteHeader(rm.Status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rm)

	// Log
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s :: %d [%s] %s\n", timestamp, rm.Status, rm.AppCode, rm.Message)

}

// DecodeError - Error decoding JSON onto struct
var DecodeError = "DECODE_ERR"

// ExternalConnectionError - Error connecting to external resource
var ExternalConnectionError = "EXT_CONNECTION_ERR"

// ExternalRequestFailed - Request to external resource failed
var ExternalRequestFailed = "EXT_REQUEST_FAIL"

// MissingKey - A key required to complete a task is missing
var MissingKey = "MISSING_KEY"

// ParseError - Error when parsing
var ParseError = "PARSE_ERROR"
