package stepper

type StepperStats struct {
	Min int
	Max int
	Avg int
}

func GenerateStepperStats(repository *StepperReposistory) StepperStats {
	total := 0
	min := 0
	max := 0

	days := repository.Days()
	for _, day := range days {
		steps := repository.Get(day)

		if min == 0 || min > steps {
			min = steps
		}

		if max < steps {
			max = steps
		}

		total += steps
	}

	avg := 0
	if len(days) > 0 {
		avg = total / len(days)
	}

	return StepperStats{
		Min: min,
		Max: max,
		Avg: avg,
	}
}
