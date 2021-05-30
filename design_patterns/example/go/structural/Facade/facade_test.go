package Facade

import "testing"

func TestCarFacade_CreateCompleteCar(t *testing.T) {
	//f := CarFacade{}
	f := NewCarFacade()
	f.CreateCompleteCar()
}
