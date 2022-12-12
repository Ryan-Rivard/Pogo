package wizard

type Step2 struct {
	name    string
	execute func(*Step2)
	next    []*Step2
	prev    *Step2
}

func (s *Step2) build(p *Step2) {
	s.prev = p

	for _, step := range s.next {
		step.prev = s
		step.build(s)
	}
}
