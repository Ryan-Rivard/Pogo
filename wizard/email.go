package wizard

func init() {
	emailStep.execute = createEmailAction(emailStep)
}

var emailStep = &Step{
	name: "Email",
}

func createEmailAction(s *Step) func() {
	return func() {
		println("my Email function goes here")
	}
}
