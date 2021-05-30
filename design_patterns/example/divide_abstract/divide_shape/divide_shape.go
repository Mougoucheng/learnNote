package divide_shape

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Line struct {
	Start, End Point
}
type Rect struct {
	Width  float64
	Height float64
	LeftUp Point
}

// Circle 新增需求画圆，新增结构体Circle
type Circle struct {
	Center Point
	Radius float64
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
	lines  []Line
	rects  []Rect

	//改变
	circles []Circle
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
		line := Line{
			Start: th.p1,
			End:   th.p2,
		}
		th.lines = append(th.lines, line)
	} else if th.rdoRect.Checked {
		rect := Rect{
			Width:  math.Abs(th.p2.X - th.p1.X),
			Height: math.Abs(th.p2.Y - th.p1.Y),
			LeftUp: th.p1,
		}
		th.rects = append(th.rects, rect)
	} else if th.rdoCircle.Checked { // 改变
		circle := Circle{
			Radius: math.Abs(th.p2.X - th.p1.X),
			Center: th.p1,
		}
		th.circles = append(th.circles, circle)
	}

	// ...
	th.OnPain()
}

func (th *MainForm) OnPain() {
	// 针对直线
	for _, p := range th.lines {
		fmt.Printf("e.Graphics.DrawLine(Pens.Red, %f, %f, %f, %f)",
			p.Start.X, p.Start.Y, p.End.X, p.End.Y)
	}

	// 针对矩形
	for _, r := range th.rects {
		fmt.Printf("e.Graphics.DrawRectangle(Pens.Red, %v, %f, %f)",
			r.LeftUp, r.Width, r.Height)
	}

	// 改变：针对圆形
	for _, c := range th.circles {
		fmt.Printf("e.Graphics.DrawCircle(Pens.Red, %v, %f)",
			c.Center, c.Radius)
	}
	// ...
}
