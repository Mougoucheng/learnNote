package bad

import (
	"log"
)

// library 程序库开发人员
type library struct {
}

func (th *library) Step1() {
	log.Println("library Step1")
}
func (th *library) Step3() {
	log.Println("library Step3")
}
func (th *library) Step5() {
	log.Println("library Step5")
}

// application 程序库开发人员
type application struct {
}

func (th *application) Step2() bool {
	log.Println("application Step2")
	return true
}
func (th *application) Step4() {
	log.Println("application Step4")
}

// MainRun 程序主流程
func MainRun() {
	lib := library{}
	app := application{}
	// 稳定的流程
	lib.Step1()
	if app.Step2() {
		lib.Step3()
	}
	for i := 0; i < 4; i++ {
		app.Step4()
	}
	lib.Step5()
}
