package emailVerifier

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"strings"
	"time"
)

type VerifierStruct struct {
	Status bool `json:"status"`
}

// Verify checks if an email address is valid
func Verify(email, accessToken string) bool {
	// Input validation
	if !isValidInput(email, accessToken) {
		return false
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := fmt.Sprintf("https://verifyright.co/verify/%s?token=%s", 
		strings.TrimSpace(email), 
		strings.TrimSpace(accessToken))

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return false
	}

	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return false
	}

	var data VerifierStruct
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return false
	}

	return data.Status
}

// isValidInput validates the email and access token
func isValidInput(email, accessToken string) bool {
	if email == "" || accessToken == "" {
		return false
	}

	// Basic email format validation
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	return true
}
