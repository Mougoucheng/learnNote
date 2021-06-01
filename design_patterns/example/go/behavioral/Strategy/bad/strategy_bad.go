package bad

import "log"

type context struct {
	ctx string
}
type taxBase int

const (
	CN_Tax taxBase = iota // 中国
	US_Tax                // 美国
	DE_Tax                // 德国
	FR_Tax                // 新增法国(更改)
)

type SalesOrder struct {
	tax taxBase
}

func NewSalesOrder(tax taxBase) *SalesOrder {
	return &SalesOrder{
		tax: tax,
	}
}
func (th *SalesOrder) CalculateTax() float64 {
	// ...
	ctx := context{
		ctx: "context of Calculate",
	}
	result := 0.0
	if th.tax == CN_Tax {
		log.Printf("Do something by CN_Tax:%s", ctx.ctx)
		result = 1.0
	} else if th.tax == US_Tax {
		log.Printf("Do something by US_Tax:%s", ctx.ctx)
		result = 2.0
	} else if th.tax == DE_Tax {
		log.Printf("Do something by DE_Tax:%s", ctx.ctx)
		result = 3.0
	} else if th.tax == FR_Tax { // 更改
		log.Printf("Do something by FR_Tax:%s", ctx.ctx)
		result = 4.0
	}
	// ...
	return result
}
