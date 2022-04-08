package main

import (
	"github.com/go-learning-project/interface/cat"
	"testing"
)

func TestAnimal(t *testing.T) {
	wantedColor := "yellow"
	oneCat := cat.NewCat(wantedColor, 10.0, 10.0)
	if oneCat.Color() != wantedColor {
		t.Errorf("test failed. result is: %s", oneCat.Color())
	}
}
