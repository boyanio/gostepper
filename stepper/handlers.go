package stepper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type StepsDto struct {
	Steps int `json:"steps"`
}

func urlPathParts(path string) []string {
	parts := make([]string, 0)
	for _, part := range strings.Split(path, "/") {
		if part != "" {
			parts = append(parts, part)
		}
	}
	return parts
}

func writeJsonToResponse(w http.ResponseWriter, body interface{}) {
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func CreateStepperHandler(repository *StepperReposistory) http.HandlerFunc {
	stats := StepperStats{}

	updateStats := func() {
		fmt.Println("Updating stats in 3 secs...")
		time.Sleep(3 * time.Second)
		stats = GenerateStepperStats(repository)
		fmt.Println("Stats updated")
	}

	//
	// GET /stepper
	// GET /stepper/{date}
	//
	getHandler := func(w http.ResponseWriter, r *http.Request) {
		urlPathParts := urlPathParts(r.URL.Path)

		if len(urlPathParts) == 1 {
			// stats
			writeJsonToResponse(w, &stats)
		} else if len(urlPathParts) == 2 {
			// steps for day
			day := urlPathParts[1]
			dto := StepsDto{
				Steps: repository.Get(day),
			}
			writeJsonToResponse(w, &dto)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	//
	// PUT /stepper/{date}
	//
	putHandler := func(w http.ResponseWriter, r *http.Request) {
		urlPathParts := urlPathParts(r.URL.Path)
		if len(urlPathParts) != 2 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		day := urlPathParts[1]

		var req StepsDto
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		repository.Add(day, req.Steps)

		go updateStats()

		w.WriteHeader(http.StatusOK)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getHandler(w, r)
		} else if r.Method == "PUT" {
			putHandler(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
