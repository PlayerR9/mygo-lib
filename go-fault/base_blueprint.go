package faults

// baseBlueprint is a base implementation of the Blueprint interface.
type baseBlueprint struct {
	// name is the name of the blueprint.
	name string
}

// String implements Blueprint.
func (bb baseBlueprint) String() string {
	return bb.name
}

// Init implements Blueprint.
func (bb baseBlueprint) Init(msg string) Fault {
	return &baseFault{
		blueprint: &bb,
		msg:       msg,
	}
}

// New creates a new Blueprint with the given name.
//
// If the name is empty, it defaults to "Generic Fault".
//
// Parameters:
//   - name: The name of the blueprint.
//
// Returns:
//   - Blueprint: The new blueprint. Never returns nil.
func New(name string) Blueprint {
	if name == "" {
		name = "Generic Fault"
	}

	return &baseBlueprint{
		name: name,
	}
}
