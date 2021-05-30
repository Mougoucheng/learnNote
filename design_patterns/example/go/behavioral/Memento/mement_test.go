package Memento

import (
	"log"
	"testing"
)

func TestNumber_ReinstateMemento(t *testing.T) {
	n := NewNumber(10)
	n.Double()
	n.Double()
	memento := n.CreateMemento()
	log.Println(n.Value())
	n.Half()
	log.Println(n.Value())
	n.ReinstateMemento(memento)
	log.Println(n.Value())
}
