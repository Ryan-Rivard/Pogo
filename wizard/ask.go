package wizard

// Composite pattern
// Container
// Element that contains sub-containers and/or leafs

type Ask struct {
	Question   string
	Answer     string
	Components []Step
}

func (a *Ask) Execute() {
	println("processing")
	a.Components[0].Execute()
}
