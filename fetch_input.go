package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func fetchInput(day int) (string, error) {
	filename := fmt.Sprintf("input/day%d.txt", day)

	// Check if file already exists
	if _, err := os.Stat(filename); err == nil {
		// File exists, just read it
		data, err := os.ReadFile(filename)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("error loading .env file: %w", err)
	}

	session := os.Getenv("SESSION")
	if session == "" {
		return "", fmt.Errorf("SESSION environment variable not set")
	}

	// Construct URL
	url := fmt.Sprintf("https://adventofcode.com/2025/day/%d/input", day)
	fmt.Printf("Making request to: %s\n", url)
	
	// Create HTTP client and request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return "", err
	}

	// Add session cookie
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: %d\n", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Response: %s\n", string(body))
		return "", fmt.Errorf("response: %s", string(body))
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return "", err
	}

	// Create input directory if it doesn't exist
	inputDir := "input"
	if err := os.MkdirAll(inputDir, 0755); err != nil {
		fmt.Printf("Error creating input directory: %v\n", err)
		return "", err
	}

	// Save to file
	err = os.WriteFile(filename, body, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return "", err
	}

	return string(body), nil
}