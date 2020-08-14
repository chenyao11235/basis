package adapter

import "testing"

func TestAdapter(t *testing.T) {
	old := NewAdaptee()
	new := NewAdapter(old)
	t.Log(new.Request())
}
