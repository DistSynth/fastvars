package fastvars

import (
	"testing"
)

func TestNewFastVars(t *testing.T) {
	_, err := NewFastVars()
	if err != nil {
		t.Fatalf("NewFastVars() failed")
	}
}

func TestNewFastVarsDict(t *testing.T) {
	_, err := NewFastVarsDict(nil)
	if err == nil {
		t.Fatalf("NewFastVarsDict() failed")
	}
}

func TestFVProcess(t *testing.T) {
	fv, _ := NewFastVars()
	fv.Append(map[string]interface{}{
		"IP":   "127.0.0.1",
		"PORT": "80",
		"URL":  "http://#{IP}:#{PORT}/",
	})
	t.Log(fv.Process("#{URL}"))
    t.Log(fv.Get("URL"))
}
