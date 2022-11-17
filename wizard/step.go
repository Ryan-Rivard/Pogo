package wizard

type Step struct {
	name    string
	execute func(*Step)
	next    []*Step
	prev    *Step
}

func (s *Step) build() {
	for _, step := range s.next {
		step.prev = s
		step.build()
	}
}
