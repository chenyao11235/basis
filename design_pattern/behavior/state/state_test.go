package state

import "testing"

func TestState(t *testing.T) {
	work := NewWork()
	work.Hour = 9
	work.WriteProgram()
	work.Hour = 13
	work.WriteProgram()
	work.Hour = 17
	work.WriteProgram()
	work.Hour = 19
	work.WriteProgram()
	work.Finished = true
	work.WriteProgram()
}
