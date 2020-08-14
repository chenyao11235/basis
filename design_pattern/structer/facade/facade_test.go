package facade

import "testing"

func TestFacade(t *testing.T) {
	api := NewAPI()
	t.Log(api.Request())
}
