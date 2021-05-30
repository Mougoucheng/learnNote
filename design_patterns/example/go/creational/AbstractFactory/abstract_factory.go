package AbstractFactory

import "log"

type Lunch interface {
	Cook()
}

type Rise struct {

}

func (th *Rise) Cook() {
	log.Println("这是大米")
}

type Tomato struct {

}

func (th *Tomato) Cook() {
	log.Println("这是西红柿")
}

type Potato struct {

}

func (th *Potato) Cook() {
	log.Println("这是马铃薯")
}

type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

type SimpleLunchFactory struct {

}

func NewSimpleLunchFactory() LunchFactory {
	return &SimpleLunchFactory{}
}

func (th *SimpleLunchFactory) CreateFood() Lunch {
	return &Rise{}
}

func (th *SimpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}