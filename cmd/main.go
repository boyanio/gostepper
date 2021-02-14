package main

import (
	"fmt"
	"log"
	"net/http"

	"boyan.io/gostepper/stepper"
)

func main() {
	repository := stepper.NewStepperReposistory()
	http.HandleFunc("/stepper/", stepper.CreateStepperHandler(repository))

	fmt.Println("Listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
