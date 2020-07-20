package consterr_test

import (
	"testing"

	. "github.com/Devoter/consterr"
)

func TestErrorError(t *testing.T) {
	const s string = "test error"
	const e = Error(s)

	if s != e.Error() {
		t.Errorf("Expected \"%s\", got \"%s\"\n", s, e.Error())
	}
}
