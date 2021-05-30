# 设计模式

## 简介

### 课程目标：

* **理解松耦合设计思想**
* **掌握面向对象设计原则**
* **掌握重构技法改善设计**
* **掌握GOF核心设计模式**

### 什么是设计模式

"每一个模式描述了一个在我们周围不断重复发生的问题，以及该问题的解决方案的核心。这样，你就能一次又一次重复地使用该方案而不必做重复劳动"。
<p align="right">--Christopher Alexander</p>

### GOF设计模式

* 历史性著作《设计模式：可复用面向对象软件基础》一书中描述了23种经典面向对象设计模式，创立了模式在软件设计中的地位。
* 由于《设计模式》一书确定了设计模式的地位，通常所说的设计模式隐含地表示"面向对象设计模式"。但这并不意味"设计模式"就等于"面向对象设计模式"。

### 从面向对象谈起

* 底层思维：向下，如何把握机器底层从微观理解对象构造
    * 语言构造
    * 编译转换
    * 内存模型
    * 运行时机制

* 抽象思维：向上，如何将我们的周围世界抽象为程序代码
    * 面向对象
    * 组件封装
    * 设计模式
    * 架构模式

### 深入理解面向对象：(向下 / 向上)

* 向下：深入理解三大面向对象机制
    * 封装，隐藏内部实现
    * 继承，复用现有代码
    * 多态，改写对象行为

* 向上：深刻把握面向对象机制所带来的抽象意义，理解如何使用这些机制来表达现实世界，掌握什么是"好的面向对象设计"

### 软件设计固有的复杂性

建筑商从来不会去想给一栋已建好的100层高的楼房底下再修建一个小地下室——这样做花费极大而且注定要失败。然而令人惊奇的是，软件系统的用户在要求作出类似改变时却不会仔细考虑，而且他们认为这只是需要简单编程的事。
<p align="right">——Object-Oriented Analysis and Design with Applications</p>

### 软件设计复杂的根本原因：(**<font size=10 color="red">变化</font>**)

* 客户需求的变化
* 技术平台的变化
* 开发团队的变化
* 市场环境的变化
* ......

### 如何解决复杂性？(**分解 / 抽象**)

* 分解
    * 人们对复杂性有一个常见的做法：即分而治之，将大问题分解为多个小问题，将复杂问题分解为多个简单问题

* 抽象
    * 更高层次来讲，人们处理复杂性有一个通用的技术，即抽象。由于不能掌握全部的复杂对象，我们选择忽视它的非本质细节，而去处理泛化和理想化了的对象模型。

**<font size=10 color="blue">重点：掌握抽象思维</font>**

#### 分解例子：[(divide_shape.go)](./example/divide_abstract/divide_shape/divide_shape.go)

> ```go
> package divide_shape
> 
> import (
> 	"fmt"
> 	"math"
> )
> 
> type Point struct {
> 	X, Y float64
> }
> 
> type Line struct {
> 	Start, End Point
> }
> type Rect struct {
> 	Width  float64
> 	Height float64
> 	LeftUp Point
> }
> 
> // Circle 新增需求画圆，新增结构体Circle
> type Circle struct {
> 	Center Point
> 	Radius float64
> }
> 
> type RadioButton struct {
> 	Checked bool
> }
> 
> type Form struct {
> 	rdoLine   RadioButton
> 	rdoRect   RadioButton
> 	rdoCircle RadioButton // 改变
> }
> 
> type MainForm struct {
> 	Form
> 
> 	p1, p2 Point
> 	lines  []Line
> 	rects  []Rect
> 
> 	//改变
> 	circles []Circle
> }
> 
> func (th *MainForm) OnMouseDown(x, y float64) {
> 	th.p1.X = x
> 	th.p1.Y = y
> 	// 鼠标按下记录第一个点，继续其他处理...
> }
> 
> func (th *MainForm) OnMouseUp(x, y float64) {
> 	th.p2.X = x
> 	th.p2.Y = y
> 	// 鼠标抬起记录第二个点，根据选择构造对象并加入slice
> 	if th.rdoLine.Checked {
> 		line := Line{
> 			Start: th.p1,
> 			End:   th.p2,
> 		}
> 		th.lines = append(th.lines, line)
> 	} else if th.rdoRect.Checked {
> 		rect := Rect{
> 			Width:  math.Abs(th.p2.X - th.p1.X),
> 			Height: math.Abs(th.p2.Y - th.p1.Y),
> 			LeftUp: th.p1,
> 		}
> 		th.rects = append(th.rects, rect)
> 	} else if th.rdoCircle.Checked { // 改变
> 		circle := Circle{
> 			Radius: math.Abs(th.p2.X - th.p1.X),
> 			Center: th.p1,
> 		}
> 		th.circles = append(th.circles, circle)
> 	}
> 
> 	// ...
> 	th.OnPain()
> }
> 
> func (th *MainForm) OnPain() {
> 	// 针对直线
> 	for _, p := range th.lines {
> 		fmt.Printf("e.Graphics.DrawLine(Pens.Red, %f, %f, %f, %f)",
> 			p.Start.X, p.Start.Y, p.End.X, p.End.Y)
> 	}
> 
> 	// 针对矩形
> 	for _, r := range th.rects {
> 		fmt.Printf("e.Graphics.DrawRectangle(Pens.Red, %v, %f, %f)",
> 			r.LeftUp, r.Width, r.Height)
> 	}
> 
> 	// 改变：针对圆形
> 	for _, c := range th.circles {
> 		fmt.Printf("e.Graphics.DrawCircle(Pens.Red, %v, %f)",
> 			c.Center, c.Radius)
> 	}
> 	// ...
> }
> 
> ```

#### 抽象例子：[(abstract_shape.go)](./example/divide_abstract/abstract_shape/abstract_shape.go)

> ``` go
>package abstract_shape
> 
> import (
> 	"fmt"
> 	"math"
> )
> 
> type Shape interface {
> 	Draw()
> }
> type Point struct {
> 	X, Y float64
> }
> type Line struct {
> 	Start, End Point
> }
> 
> // Draw Line实现自己的Draw，负责画自己
> func (th *Line) Draw() {
> 	fmt.Printf("e.Graphics.DrawLine(Pens.Red, %f, %f, %f, %f)",
> 		th.Start.X, th.Start.Y, th.End.X, th.End.Y)
> }
> 
> type Rect struct {
> 	LeftUp Point
> 	Width  float64
> 	Height float64
> }
> 
> // Draw Rect实现自己的Draw，负责画自己
> func (th *Rect) Draw() {
> 	fmt.Printf("e.Graphics.DrawRectangle(Pens.Red, %v, %f, %f)",
> 		th.LeftUp, th.Width, th.Height)
> }
> 
> // Circle 新增需求画圆，新增结构体Circle
> type Circle struct {
> 	Center Point
> 	Radius float64
> }
> 
> // Draw Circle实现自己的Draw，负责画自己
> func (th *Circle) Draw() {
> 	fmt.Printf("e.Graphics.DrawCircle(Pens.Red, %v, %f)",
> 		th.Center, th.Radius)
> }
> 
> type RadioButton struct {
> 	Checked bool
> }
> 
> type Form struct {
> 	rdoLine   RadioButton
> 	rdoRect   RadioButton
> 	rdoCircle RadioButton // 改变
> }
> 
> type MainForm struct {
> 	Form
> 	p1, p2 Point
> 	shapes []Shape
> }
> 
> func (th *MainForm) OnMouseDown(x, y float64) {
> 	th.p1.X = x
> 	th.p1.Y = y
> 	// 鼠标按下记录第一个点，继续其他处理...
> }
> 
> func (th *MainForm) OnMouseUp(x, y float64) {
> 	th.p2.X = x
> 	th.p2.Y = y
> 	// 鼠标抬起记录第二个点，根据选择构造对象并加入slice
> 	if th.rdoLine.Checked {
> 		line := &Line{
> 			Start: th.p1,
> 			End:   th.p2,
> 		}
> 		th.shapes = append(th.shapes, line)
> 	} else if th.rdoRect.Checked {
> 		rect := &Rect{
> 			Width:  math.Abs(th.p2.X - th.p1.X),
> 			Height: math.Abs(th.p2.Y - th.p1.Y),
> 			LeftUp: th.p1,
> 		}
> 		th.shapes = append(th.shapes, rect)
> 	} else if th.rdoCircle.Checked { // 改变
> 		circle := &Circle{
> 			Radius: math.Abs(th.p2.X - th.p1.X),
> 			Center: th.p1,
> 		}
> 		th.shapes = append(th.shapes, circle)
> 	}
> 	// ...
> 	th.OnPain()
> }
> 
> func (th *MainForm) OnPain() {
> 	// 针对所有形状
> 	for _, s := range th.shapes {
> 		s.Draw() // 多态调用，各负其责
> 	}
> 
> 	// ...
> }
> 
> ```

### 软件设计的目标：(**<font size=10 color="red">复用</font>**)

什么是好的软件设计？软件设计的金科玉律：复用

---

## TODO: to be continue~~~



