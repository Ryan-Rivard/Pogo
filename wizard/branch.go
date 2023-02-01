package wizard

type branch struct {
	refname            string
	relativeCommitDate string
}

type branches struct {
	branches []branch
}
