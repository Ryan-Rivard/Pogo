package wizard

func init() {
	sharing_updatingStep.execute = createSharing_UpdatingAction(sharing_updatingStep)
}

var sharing_updatingStep = &Step{
	name: "Sharing and Updating Projects",
}

func createSharing_UpdatingAction(s *Step) func() {
	return func() {
		println("my Sharing and Updating function goes here")
	}
}
