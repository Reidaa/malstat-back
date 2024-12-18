package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	timeout                           = 10
	unsuccessfulHTTPResponseThreshold = 300
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

// Info writes logs in the color blue with "INFO: " as prefix.
var Info = log.New(os.Stdout, "\u001b[34mINFO: \u001B[0m", log.LstdFlags)

// Warning writes logs in the color yellow with "WARNING: " as prefix.
var Warning = log.New(os.Stdout, "\u001b[33mWARNING: \u001B[0m", log.LstdFlags|log.Lshortfile)

// Error writes logs in the color red with "ERROR: " as prefix.
var Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)

// Debug writes logs in the color cyan with "DEBUG: " as prefix.
var Debug = log.New(os.Stdout, "\u001b[36mDEBUG: \u001B[0m", log.LstdFlags|log.Lshortfile)

var NetClient = &http.Client{
	Timeout: time.Second * timeout,
}

type UnsuccessfulRequestError struct {
	StatusCode int
	Url        string
}

func (e *UnsuccessfulRequestError) Error() string {
	return fmt.Sprintf("http response status code from %s is not successful: %d", e.Url, e.StatusCode)
}

func HttpGet(url string) ([]byte, error) {
	var client = &http.Client{
		Timeout: time.Second * timeout,
	}

	Debug.Printf("GET %s", url)

	// response, err := http.Get(url)
	response, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request to %s failed: %w", url, err)
	}

	defer response.Body.Close()

	if response.StatusCode >= unsuccessfulHTTPResponseThreshold {
		return nil, &UnsuccessfulRequestError{
			StatusCode: response.StatusCode,
			Url:        url,
		}
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body from %s: %w", url, err)
	}

	return responseData, nil
}
