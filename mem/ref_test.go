package mem

import "testing"

type MockMyType struct {
}

func NewMockMyType() *MockMyType {
	return &MockMyType{}
}

func (mt *MockMyType) free() {
	if mt == nil {
		return
	}

	// Assume there is the cleanup procedure here.
}

func MockRes(mt *MockMyType) *Ref {
	NewRef(mt, mt.free)
}

func TestRef(t *testing.T) {
	mt := NewMockMyType()
}
