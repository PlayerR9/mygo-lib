package fault

// baseFault is a base implementation of the Fault interface.
type baseFault struct {
	// blueprint is the blueprint of the Fault.
	blueprint Blueprint

	// msg is the error message.
	msg string
}

// Error implements Fault.
func (f baseFault) Error() string {
	return f.msg
}

// Init implements Fault.
func (f baseFault) Init(msg string) Fault {
	return &baseFault{
		blueprint: f.blueprint,
		msg:       msg,
	}
}
