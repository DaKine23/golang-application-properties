package golangapplicationproperties_test

import (
	"testing"

	"github.com/DaKine23/golangapplicationproperties"
)

func TestInit(t *testing.T) {

	props, err := golangapplicationproperties.New("golangapplicationproperties_test.go")
	if err != nil {
		t.Error(err)
	}
	m := props.PropertyMap

	if m["m :"] != "props.PropertyMap" {
		t.Error("wrong value for key \"m :\"")
	}

	_, err = golangapplicationproperties.New("wasIeverThere?")

	if err == nil {
		t.Error("file should not be found")
	}

}
