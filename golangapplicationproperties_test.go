package golangapplicationproperties

import (
	"fmt"
	"testing"
	"time"
)

func TestConstructor(t *testing.T) {

	var notInitialized Properties

	if notInitialized.IsInitialized() {
		t.Error("Properties was not initialized yet but IsInitialized returns true")
	}

	props, err := NewProperties("golangapplicationproperties_test.go") //#some comment
	if err != nil {
		t.Error(err)
	}

	if !props.IsInitialized() {
		t.Error("Properties was initialized but IsInitialized returns false")
	}

	m := props.PropertyMap

	if m["props, err :"] != "NewProperties(\"golangapplicationproperties_test.go\") //" {
		t.Error("wrong value for key \"m :\"")
	}

	fmt.Println("Filepath was :", props.FilePath)

	var zeroValue time.Time
	if props.InitTime == zeroValue {
		t.Error("Init time was not set")
	}

	if props.FilePath != "golangapplicationproperties_test.go" {
		t.Error("FilePath was not set")
	}

	_, err = NewProperties("wasIeverThere?")

	if err == nil {
		t.Error("file should not be found")
	}

}

func BenchmarkInit(b *testing.B) {

	props, _ := NewProperties("golangapplicationproperties_test.go")
	props.IsInitialized()
}
