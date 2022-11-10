package wizard

type Step struct {
	name    string
	execute func()

	// nextStep map[string]*Step
	next []*Step
	// prev    *Step
	// options []string
	// gitCmd []string
}

func (s *Step) addNext(next *Step) {
	s.next = append(s.next, next)
}

// func (s *Step) goBack() {
// 	if s.prev != nil {
// 		s.prev.doer()
// 	}
// }

// func (s *Step) run() {
// 	s.doer()

// }

// ask something
// - somtimes need input
// - produce output

// do something
// run git cmd and ask something
// run git cmd and exit
// go back
// exit

// func (s *Step) processStep() {
// 	// if len(s.gitCmd) > 0 {
// 	// 	executeGitCommand(s.gitCmd...)
// 	// }

// 	// if s.nextStep != nil {
// 	// 	s.nextStep[""].processStep()
// 	// }
// }
