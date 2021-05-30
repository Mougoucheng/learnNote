package Facade

import "log"

type CarModel struct {
}

func NewCarModel() *CarModel {
	return &CarModel{}
}

func (th *CarModel) SetModel() {
	log.Println("CarModel - SetModel")
}

type CarEngine struct {
}

func NewCarEngine() *CarEngine {
	return &CarEngine{}
}

func (th *CarEngine) SetEngine() {
	log.Println("CarEngine - SetEngine")
}

type CarBody struct {
}

func NewCarBody() *CarBody {
	return &CarBody{}
}

func (th *CarBody) SetBody() {
	log.Println("CarBody - SetCarBody")
}

type CarFacade struct {
	model  CarModel
	engine CarEngine
	body   CarBody
}

func NewCarFacade() *CarFacade {
	return &CarFacade{
		model:  CarModel{},
		engine: CarEngine{},
		body:   CarBody{},
	}
}

func (th *CarFacade) CreateCompleteCar() {
	th.model.SetModel()
	th.engine.SetEngine()
	th.body.SetBody()
}