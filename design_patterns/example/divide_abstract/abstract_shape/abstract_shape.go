package abstract_shape

import (
	"fmt"
	"math"
)

type Shape interface {
	Draw()
}
type Point struct {
	X, Y float64
}
type Line struct {
	Start, End Point
}

// Draw Line实现自己的Draw，负责画自己
func (th *Line) Draw() {
	fmt.Printf("e.Graphics.DrawLine(Pens.Red, %f, %f, %f, %f)",
		th.Start.X, th.Start.Y, th.End.X, th.End.Y)
}

type Rect struct {
	LeftUp Point
	Width  float64
	Height float64
}

// Draw Rect实现自己的Draw，负责画自己
func (th *Rect) Draw() {
	fmt.Printf("e.Graphics.DrawRectangle(Pens.Red, %v, %f, %f)",
		th.LeftUp, th.Width, th.Height)
}

// Circle 新增需求画圆，新增结构体Circle
type Circle struct {
	Center Point
	Radius float64
}

// Draw Circle实现自己的Draw，负责画自己
func (th *Circle) Draw() {
	fmt.Printf("e.Graphics.DrawCircle(Pens.Red, %v, %f)",
		th.Center, th.Radius)
}

type RadioButton struct {
	Checked bool
}

type Form struct {
	rdoLine   RadioButton
	rdoRect   RadioButton
	rdoCircle RadioButton // 改变
}

type MainForm struct {
	Form
	p1, p2 Point
	shapes []Shape
}

func (th *MainForm) OnMouseDown(x, y float64) {
	th.p1.X = x
	th.p1.Y = y
	// 鼠标按下记录第一个点，继续其他处理...
}

func (th *MainForm) OnMouseUp(x, y float64) {
	th.p2.X = x
	th.p2.Y = y
	// 鼠标抬起记录第二个点，根据选择构造对象并加入slice
	if th.rdoLine.Checked {
		line := &Line{
			Start: th.p1,
			End:   th.p2,
		}
		th.shapes = append(th.shapes, line)
	} else if th.rdoRect.Checked {
		rect := &Rect{
			Width:  math.Abs(th.p2.X - th.p1.X),
			Height: math.Abs(th.p2.Y - th.p1.Y),
			LeftUp: th.p1,
		}
		th.shapes = append(th.shapes, rect)
	} else if th.rdoCircle.Checked { // 改变
		circle := &Circle{
			Radius: math.Abs(th.p2.X - th.p1.X),
			Center: th.p1,
		}
		th.shapes = append(th.shapes, circle)
	}
	// ...
	th.OnPain()
}

func (th *MainForm) OnPain() {
	// 针对所有形状
	for _, s := range th.shapes {
		s.Draw() // 多态调用，各负其责
	}

	// ...
}
