package factory

import "testing"

func TestSimple(t *testing.T) {
	s := new(RuleConfigSource)
	r := s.load()
	t.Log(r.configFormat)
}

func TestMethod(t *testing.T) {
	s := new(RuleConfigSource1)
	r := s.load()
	t.Log(r.configFormat)
}
