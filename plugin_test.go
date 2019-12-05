package goplugin

import(
	"reflect"
	"testing"
)

type Foo interface{
	Say()
}

type Bar struct {}

func (this *Bar) Say() {
	println("Hello world")
}

type Nobar struct {}

func TestNewManager(t *testing.T) {
	manager := NewManager("foo")
	manager.AddExtension(NewExtension("bar", reflect.TypeOf((*Foo)(nil)).Elem(), "bar", "bar"))

	err := Register(&Plugin{Extension: "foo/bar", Implement: &Bar{}})
	if err != nil {
		t.Error(err)
	}

	err = Register(&Plugin{Extension: "foo/bar", Implement: &Nobar{}})
	if err == nil {
		t.Error("Type assert failed.")
	}

	implements := manager.GetImplements("bar")
	if implements == nil {
		t.Error("Nil implements.")
	}

	implements[0].(Foo).Say()
}