package pluginloader

import (
	"testing"
)

func TestLoader(t *testing.T) {
	buildFoo()
	p, err := NewPluginLoader("foo.so")
	if err != nil {
		t.Fatal(err)
	}
	p.CallValue("Hello", nil)
	p.CallValue("Hello2", "a", "b", "c")
}
