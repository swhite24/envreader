package envreader

import (
	"testing"
)

var keys = []string{ "PATH", "USER" }

func TestRead (t *testing.T) {
	vals := Read(keys...)
	if len(vals) != len(keys) {
		t.Errorf("Incorrect number of fields")
	}

	for _, key := range keys {
		if vals[key] == "" {
			t.Errorf("Missing key: %s\n", key)
		}
	}
}