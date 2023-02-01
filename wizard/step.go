package wizard

type step interface {
	Exec(interface{})
}
