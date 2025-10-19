package script

import (
	"testing"
)

func TestGlobalThis(t *testing.T) {
	m := make(map[string]any)

	GlobalThis(m).SetOutput("a")

	a := GlobalThis(m).Output()

	if a != "a" {
		t.Error("GlobalThis error")
	}
}
