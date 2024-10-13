package common

// Cleaner is an interface that can be used to clean up resources.
type Cleaner interface {
	// Cleanup cleans up resources. Does nothing if
	// the receiver is nil or already cleaned up.
	Cleanup()
}
