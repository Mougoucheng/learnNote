package Mediator

import "log"

type Mediator interface {
	Communicate(who string)
}

type WildStallion interface {
	SetMediator(mediator Mediator)
}

type Bill struct {
	m Mediator
}

func (th *Bill) SetMediator(mediator Mediator) {
	th.m = mediator
}
func (th *Bill) Respond() {
	log.Println("Bill what ?")
	th.m.Communicate("Bill")
}

type Ted struct {
	m Mediator
}

func (th *Ted) SetMediator(mediator Mediator) {
	th.m = mediator
}

func (th *Ted) Talk() {
	log.Println("Ted:Bill ?")
	th.m.Communicate("Ted")
}

func (th *Ted) Respond() {
	log.Println("Ted:how much?")
}

type ConcreteMediator struct {
	Bill
	Ted
}

func NewConcreteMediator() *ConcreteMediator {
	mediator := &ConcreteMediator{}
	mediator.Bill.SetMediator(mediator)
	mediator.Ted.SetMediator(mediator)
	return mediator
}

func (th *ConcreteMediator) Communicate(who string) {
	switch who {
	case "Ted":
		th.Bill.Respond()
	case "Bill":
		th.Ted.Respond()
	default:
		log.Println("no body")
	}
}
