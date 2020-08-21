package builder

import "testing"

func TestBuilder(t *testing.T) {
	c := &ResourcePoolConfig{}
	r, err := NewMySQLBuilder(c).
		setName("dbconnectionpool").
		setMaxTotal(16).
		setMaxIdle(10).
		setMinIdle(12).
		Build()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}
