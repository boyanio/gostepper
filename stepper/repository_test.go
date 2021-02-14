package stepper

import (
	"testing"

	"github.com/bruth/assert"
)

func TestRepository_ReturnsZeroStepsWhenNoStepsAdded(t *testing.T) {
	repository := NewStepperReposistory()
	steps := repository.Get("2020-02-25")
	assert.Equal(t, 0, steps)
}

func TestRepository_ReturnsStepsForDay(t *testing.T) {
	expectedSteps := 1500
	repository := NewStepperReposistory()
	repository.Add("2020-02-25", expectedSteps)
	steps := repository.Get("2020-02-25")
	assert.Equal(t, expectedSteps, steps)
}

func TestRepository_SumsStepsForDay(t *testing.T) {
	steps1 := 1000
	steps2 := 500
	expectedSteps := steps1 + steps2
	repository := NewStepperReposistory()
	repository.Add("2020-02-25", steps1)
	repository.Add("2020-02-25", steps2)
	steps := repository.Get("2020-02-25")
	assert.Equal(t, expectedSteps, steps)
}
