package good

import (
	"log"
)

type iApplication interface {
	Step2() bool // 变化
	Step4()      // 变化
}

// library 程序库开发人员
type library struct {
	app iApplication
}

func newLibrary(app iApplication) *library {
	return &library{
		app: app,
	}
}
func (th *library) Step1() { // 稳定
	log.Println("library Step1")
}
func (th *library) Step5() { // 稳定
	log.Println("library Step5")
}
func (th *library) Step3() { // 稳定
	log.Println("library Step3")
}

// Run 稳定 template method
func (th *library) Run() {
	th.Step1()
	if th.app.Step2() { // 支持变化 ==> 接口的多态调用
		th.Step3()
	}
	for i := 0; i < 4; i++ {
		th.app.Step4() // 支持变化 ==> 接口的多态调用
	}
	th.Step5()
}

// application 程序库开发人员
type application struct {
}

func (th *application) Step2() bool { // 变化：接口实现者实现
	log.Println("application Step2")
	return true
}
func (th *application) Step4() { // 变化：接口实现者实现
	log.Println("application Step4")
}

// MainRun 程序主流程
func MainRun() {
	app := &application{}
	lib := newLibrary(app)
	lib.Run()
}
