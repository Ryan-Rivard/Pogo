package wizard

func init() {
	server_adminStep.execute = createServer_AdminAction(server_adminStep)
}

var server_adminStep = &Step{
	name: "Server Admin",
}

func createServer_AdminAction(s *Step) func() {
	return func() {
		println("my Server Admin function goes here")
	}
}
