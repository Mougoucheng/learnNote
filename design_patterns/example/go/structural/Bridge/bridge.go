package Bridge

import "log"

type Draw interface {
	DrawCircle(x, y, radius int)
}
type RedCircle struct {
}

func (th *RedCircle) DrawCircle(x, y, radius int) {
	log.Println("RedCircle x,y,radius", x, y, radius)
}

type YellowCircle struct {
}

func (th *YellowCircle) DrawCircle(x, y, radius int) {
	log.Println("YellowCircle x,y,radius", x, y, radius)
}

type Shape struct {
	draw Draw
}

func (th *Shape) SetDraw(d Draw) {
	th.draw = d
}

type Circle struct {
	sp     Shape
	x      int
	y      int
	radius int
}

func (th *Circle) Constructor(x, y, radius int, d Draw) {
	th.x = x
	th.y = y
	th.radius = radius
	th.sp.SetDraw(d)
}

func (th *Circle) Draw() {
	th.sp.draw.DrawCircle(th.x, th.y, th.radius)
}
