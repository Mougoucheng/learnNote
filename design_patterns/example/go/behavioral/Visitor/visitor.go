package Visitor

import "log"

// IVisitor 访问者
type IVisitor interface {
	Visit()
}

type Daddy struct {

}

func (th *Daddy) Visit() {
	log.Println("Daddy come to see you, and say: hey son!")
}

type Mammy struct {

}

func (th *Mammy) Visit() {
	log.Println("Mammy come to see you, and say: hi honey!")
}

type Sister struct {

}

func (th *Sister) Visit() {
	log.Println("Sister come to see you, and say: hi big brother!")
}

type Brother struct {

}

func (th *Brother) Visit() {
	log.Println("Brother come to see you, and say: hi bro!")
}

type Stranger struct {

}

func (th *Stranger) Visit() {
	log.Println("Stranger come to you, and say: f**k you!")
}


// IAcceptor 受访者
type IAcceptor interface {
	Accept(visitor IVisitor)
}

type Me struct {

}

func (th *Me) Accept(v IVisitor) {
	v.Visit()
}