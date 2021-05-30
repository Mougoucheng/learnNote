package Builder

type Builder interface {
	Build()
}

type Director struct {
	builder Builder
}

func NewBuilder(b Builder) Director { // TODO??可以不要？重复？
	return Director{
		builder: b,
	}
}
func (th *Director) Construct() {
	th.builder.Build()
}


func NewDirector(b Builder) Director {
	return Director{
		builder: b,
	}
}

type ConcreteBuilder struct {
	built bool
}

func NewConcreteBuilder() ConcreteBuilder {
	return ConcreteBuilder{false}
}

func (th *ConcreteBuilder) Build()  {
	th.built = true
}

type Product struct {
	Built bool
}

func (th *ConcreteBuilder) GetResult() Product {
	return Product{th.built}
}