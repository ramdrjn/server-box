package serverbox

import "testing"

func TestInit(t *testing.T) {
	err := Initialize(true, "./pkgs/common/sample.conf")
	if err != nil {
		t.Error(err)
	}
}
