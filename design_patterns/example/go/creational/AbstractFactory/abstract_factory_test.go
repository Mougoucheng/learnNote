package AbstractFactory

import "testing"

func TestNewSimpleLunchFactory(t *testing.T) {
	factor := NewSimpleLunchFactory()
	food := factor.CreateFood()
	food.Cook()

	vegetable := factor.CreateVegetable()
	vegetable.Cook()
}
