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

func TestGlobalThis_Copy(t *testing.T) {
	globalThis := map[string]any{
		"input": 100,
		"session": map[string]any{
			"userId": 100,
		},
		"ctx": 100,
	}

	wrappedParam := make(GlobalThis)
	wrappedParam.SetInput(100)
	wrappedParam.Copy(globalThis)

	wrappedParam["input"] = 200
	sess := wrappedParam["session"].(map[string]any)
	sess["userId"] = 200
	wrappedParam["ctx"] = 200

	t.Log(globalThis)
	t.Log(wrappedParam)
}
