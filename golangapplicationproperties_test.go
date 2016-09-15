package golangapplicationproperties_test

import (
	"testing"

	"github.com/DaKine23/golangapplicationproperties"
)

func TestInit(t *testing.T) {
	var p *golangapplicationproperties.Properties
	props, err := p.Init("golangapplicationproperties_test.go")
	if err != nil {
		t.Error(err)
	}
	m := props.GetProps()

	if m["m :"] != "props.GetProps()" {
		t.Error("wrong value for key \"m :\"")
	}

	_, err = p.Init("wasIeverThere?")

	if err == nil {
		t.Error("file should not be found")
	}

}
