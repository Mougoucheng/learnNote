package good

import "log"

type context struct {
	ctx string
}

type iTaxStrategy interface {
	Calculate(context context) float64
}

type CNTax struct {
}

func NewCNTax() *CNTax {
	return &CNTax{}
}
func (th *CNTax) Calculate(ctx context) float64 {
	log.Printf("Do something by CN_Tax:%s", ctx.ctx)
	return 1.0
}

type USTax struct {
}

func NewUSTax() *USTax {
	return &USTax{}
}
func (th *USTax) Calculate(ctx context) float64 {
	log.Printf("Do something by USTax:%s", ctx.ctx)
	return 2.0
}

type DETax struct {
}

func NewDETax() *DETax {
	return &DETax{}
}
func (th *DETax) Calculate(ctx context) float64 {
	log.Printf("Do something by DETax:%s", ctx.ctx)
	return 3.0
}

// FRTax 新增法国（扩展）
type FRTax struct {
}

func NewFRTax() *FRTax {
	return &FRTax{}
}

func (th *FRTax) Calculate(ctx context) float64 {
	log.Printf("Do something by FRTax:%s", ctx.ctx)
	return 3.0
}

type SalesOrder struct {
	taxStrategy iTaxStrategy
}

//func NewSalesOrder(stgFactory *StrategyFactory) *SalesOrder {
//	return &SalesOrder{
//		taxStrategy: stgFactory.NewStrategy(),
//	}
//}

func NewSalesOrder(tax iTaxStrategy) *SalesOrder {
	return &SalesOrder{
		taxStrategy: tax,
	}
}

func (th *SalesOrder) CalculateTax() float64 {
	// ...
	ctx := context{
		ctx: "context of Calculate",
	}
	return th.taxStrategy.Calculate(ctx)
}
