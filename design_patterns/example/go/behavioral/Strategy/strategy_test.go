package Strategy

import (
	"chenyx/design_patterns/example/go/behavioral/Strategy/bad"
	"chenyx/design_patterns/example/go/behavioral/Strategy/good"
	"testing"
)

func TestStrategy(t *testing.T) {
	badSo := bad.NewSalesOrder(bad.DE_Tax)
	badSo.CalculateTax()

	stg := good.NewDETax()
	goodSo := good.NewSalesOrder(stg)
	goodSo.CalculateTax()
}