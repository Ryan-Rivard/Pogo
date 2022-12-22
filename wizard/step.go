package wizard

// Composite Pattern
// Component Interface
// Describes operations that are common to both simple and complex elements of the tree

type Step interface {
	GetId() string
	Execute()
}
