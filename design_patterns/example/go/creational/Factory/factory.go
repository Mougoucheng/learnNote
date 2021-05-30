package Factory

import "log"

type Restaurant interface {
	GetFood()
}

type DongLaiShun struct {
}

func (th *DongLaiShun) GetFood() {
	log.Println("东来顺的饭菜已经做好了，请就餐")
}

type QingFeng struct {

}

func (th *QingFeng) GetFood() {
	log.Println("庆丰包子铺的包子已经做好了，请就餐")
}

func NewRestaurant(name string) Restaurant {
	switch name {
	case "d":
		return &DongLaiShun{}
	case "q":
		return &QingFeng{}
	default:
		return nil
	}
}