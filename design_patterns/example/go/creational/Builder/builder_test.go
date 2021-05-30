package Builder

import (
	"log"
	"testing"
)

func TestConcreteBuilder_GetResult(t *testing.T) {
	builder := NewConcreteBuilder()
	director := NewDirector(&builder)
	director.Construct()
	result := builder.GetResult()
	log.Println(result.Built)
}
