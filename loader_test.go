package pluginloader

import (
	"os/exec"
	"testing"
)

func buildFoo2() {
	cmd := exec.Command("go", "build", "-o", "foo2.so", "-buildmode=plugin", "./cmd/pluginwrap/testdata/foo2")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func TestLoader(t *testing.T) {
	buildFoo()
	run1(t)
	buildFoo2()
	run2(t)
}

func run1(t *testing.T) {
	p, err := NewPluginLoader("foo.so")
	if err != nil {
		t.Fatal(err)
	}
	p.CallValue("Hello", nil)
	p.CallValue("Hello2", "a", "b", "c")
}

func run2(t *testing.T) {
	p, err := NewPluginLoader("foo2.so")
	if err != nil {
		t.Fatal(err)
	}
	p.CallValue("Hello", nil)
	p.CallValue("Hello2", "a", "b", "c")
}
