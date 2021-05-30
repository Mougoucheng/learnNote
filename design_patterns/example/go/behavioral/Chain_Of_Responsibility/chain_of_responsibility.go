package Chain_Of_Responsibility

import "fmt"

type IHandler interface {
	Handle(handlerId int) string
}

type Handler struct {
	id   int
	name string
	next IHandler
}

func NewHandler(id int, name string, next IHandler) *Handler {
	return &Handler{
		id:   id,
		name: name,
		next: next,
	}
}

func (th *Handler) Handle(id int) string {
	// 是自己的职责，处理
	if th.id == id {
		return fmt.Sprintf("%d:%s handled %d", th.id, th.name, id)
	}
	if th.next == nil {
		return fmt.Sprintf("当前责任人[%d:%s]处理不了[%d]号任务，且已经没有下一个责任人了。", th.id, th.name, id)
	}
	// 否则传给职责链上的下一个人处理
	return th.next.Handle(id)
}
