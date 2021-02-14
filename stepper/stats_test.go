package stepper

import (
	"testing"

	"github.com/bruth/assert"
)

func TestGenerateStepperStats(t *testing.T) {
	repository := NewStepperReposistory()
	repository.Add("2020-02-24", 1000)
	repository.Add("2020-02-25", 2000)

	stats := GenerateStepperStats(repository)

	assert.Equal(t, 1000, stats.Min)
	assert.Equal(t, 2000, stats.Max)
	assert.Equal(t, 1500, stats.Avg)
}
