package wizard

type step interface {
	Exec([]string)
}
