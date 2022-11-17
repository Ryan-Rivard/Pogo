package wizard

// "github.com/Ryan-Rivard/Pogo/inquire"

func Process(s string) {
	if s == "root" {
		execWizardFromStep(rootStep)
	} else if s == "fetch" {
		execWizardFromStep(fetchStep)
	}
}

func execWizardFromStep(s *Step) {
	s.prev = exitStep
	s.build()
	s.execute(s)
}
