package main

type pogoBasic struct {
	inquier iInquier
}

func (b *pogoBasic) process() {
	options := []string{"branch", "fetch", "pull", "push"}

	answer := b.inquier.askWithPresetAnswers("Answer", "Basic - What can pogo do for you today?", options)

	git := &gitService{}

	if *answer == "branch" {
		processBranch := &pogoBranch{
			git: git,
		}

		processBranch.process()
	}

	if *answer == "fetch" {
		processFetch := &pogoFetch{
			git: git,
		}

		processFetch.process()
	}

	if *answer == "pull" {
		processPull := &pogoPull{
			git: git,
		}

		processPull.process()
	}

	if *answer == "push" {
		processPush := &pogoPush{
			git: git,
		}

		processPush.process()
	}

}
