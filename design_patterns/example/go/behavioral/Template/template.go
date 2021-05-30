package Template

import "log"

type IWorker interface {
	GetUp()
	Work()
	Sleep()
}

type Worker struct {
	IWorker
}

func NewWorker(w IWorker) *Worker {
	return &Worker{w}
}

func (th *Worker) Daily() {
	th.GetUp()
	th.Work()
	th.Sleep()
}

type Coder struct {

}

func (th *Coder) GetUp() {
	log.Println("Coder GetUp")
}
func (th *Coder) Work() {
	log.Println("Coder Work")
}
func (th *Coder) Sleep() {
	log.Println("Coder Sleep")
}