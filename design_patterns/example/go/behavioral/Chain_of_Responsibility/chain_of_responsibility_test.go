package Chain_of_Responsibility

import (
	"log"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	w := NewHandler(3, "老万", nil) // 没有下一个人
	z := NewHandler(2, "老张", w)   // 老张后面是老万
	l := NewHandler(1, "老李", z)   // 老李后面是老张

	log.Println(l.Handle(1)) // 老李1号可以直接处理自己的责任
	log.Println(l.Handle(2)) // 老李1号处理不了2号，往后传给老张2号处理
	log.Println(z.Handle(2)) // 老张2号可以直接处理自己的责任
	log.Println(z.Handle(3)) // 老张2号处理不了3号，往后传给老万3号处理
	log.Println(w.Handle(3)) // 老万3号可以直接处理自己的责任
	log.Println(w.Handle(1)) // 老万3号，只能处理3号或者3号以上往后传，1号责任应该已经在前面处理了，不该这么调用(或者修改handle允许处理)
}
