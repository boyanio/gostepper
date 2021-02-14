package stepper

type StepperReposistory struct {
	steps map[string]int
}

func NewStepperReposistory() *StepperReposistory {
	return &StepperReposistory{
		steps: map[string]int{},
	}
}

func (r *StepperReposistory) Add(day string, steps int) {
	stepsForDay, ok := r.steps[day]
	if ok {
		r.steps[day] = stepsForDay + steps
	} else {
		r.steps[day] = steps
	}
}

func (r *StepperReposistory) Get(day string) int {
	stepsForDay := r.steps[day]
	return stepsForDay
}

func (r *StepperReposistory) Days() []string {
	days := make([]string, 0, len(r.steps))
	for day := range r.steps {
		days = append(days, day)
	}
	return days
}
