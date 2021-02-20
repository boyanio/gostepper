package stepper

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bruth/assert"
)

func jsonRequest(t *testing.T, method string, url string, body interface{}) *http.Request {
	var bodyReader io.Reader
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}

		bodyReader = bytes.NewBuffer(buf)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req
}

func TestGetStepperHandler(t *testing.T) {
	repository := NewStepperReposistory()
	handler := http.HandlerFunc(CreateStepperHandler(repository))

	expectedSteps := 1500

	// Add steps for 2020-02-25
	req := jsonRequest(t, "PUT", "/stepper/2020-02-25", &Steps{Steps: expectedSteps})
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)

	// Check stats
	req = jsonRequest(t, "GET", "/stepper/2020-02-25", nil)
	resp = httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)

	var dto Steps
	json.NewDecoder(resp.Body).Decode(&dto)

	assert.Equal(t, expectedSteps, dto.Steps)
}
