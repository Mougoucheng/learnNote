package Bridge

import (
	"chenyx/design_patterns/example/go/structural/Bridge/bad"
	"log"
	"testing"
)

func TestBridge(t *testing.T) {
	// bad 编译时装配
	messager := bad.MobileMessagerPerfect{
		MobileBase: bad.MobileMessagerBase{},
	}
	log.Println(messager)
}
