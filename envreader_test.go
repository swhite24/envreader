package envreader

import (
	"os"
	"testing"
)

var (
	keys   = []string{"ENVREADER1", "ENVREADER2"}
	values = []string{"foo1", "foo2"}
)

func TestRead(t *testing.T) {
	setup()
	vals := Read(keys...)
	if len(vals) != len(keys) {
		t.Errorf("Incorrect number of fields")
	}

	for i, key := range keys {
		val := values[i]
		if vals[key] == "" {
			t.Errorf("Missing key: %s\n", key)
		}
		if vals[key] != val {
			t.Errorf("Incorrect key for %s: Got %s, Want %s", key, vals[key], val)
		}
	}
}

func setup() {
	for i := range keys {
		os.Setenv(keys[i], values[i])
	}
}
