package wizard

// "github.com/Ryan-Rivard/Pogo/inquire"

func ProcessRoot() {

	buildTree()
	// do something here

	// options := []string{
	// 	"init",
	// 	"clone",
	// 	"branch",
	// 	"fetch",
	// 	"pull",
	// 	"status",
	// 	"add",
	// 	"commit",
	// 	"push",
	// }

	// options := []string{
	// 	"status",
	// 	"version",
	// }

	// myStep := &Step{
	// 	prev:    nil,
	// 	options: options,
	// }

	// search := inquier.AskWithOptions("What git command would you like to execute?", options)

	// println("you selected", *search)

	// wiz := createTree()

	// ans := wiz.doer()
	// wiz.

	rootStep.execute()

}

func buildTree() {
	rootStep.addNext(setup_configStep)
	rootStep.addNext(getting_creatingStep)
	rootStep.addNext(basic_snapshottingStep)
	rootStep.addNext(branching_mergingStep)
	rootStep.addNext(sharing_updatingStep)
	rootStep.addNext(inspection_comparisonStep)
	rootStep.addNext(patchingStep)
	rootStep.addNext(debuggingStep)
	rootStep.addNext(guidesStep)
	rootStep.addNext(emailStep)
	rootStep.addNext(external_systemsStep)
	rootStep.addNext(administrationStep)
	rootStep.addNext(server_adminStep)
	rootStep.addNext(plumbing_commandsStep)
	rootStep.addNext(exitStep)
}
