package serverbox

import "testing"

func TestInit(t *testing.T) {
	err := Initialize(true, "./internal/sample.conf")
	if err != nil {
		t.Error(err)
	}
}
