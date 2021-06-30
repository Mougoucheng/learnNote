<center><font size=10 color="blue"> Bridge [桥接模式] </font></center>

[GOF_23](../DesignPatterns.md)

# "单一职责"模式:

* 在软件组件的设计中，如果责任划分的不清晰，使用继承得到的结果往往是随着需求的变化，子类急剧膨胀，同时充斥着重复代码，这时候的关键是划清责任。
* 典型模式：
  * [Decorator](./gof_23/Decorator.md)
  * [Bridge](./gof_23/Bridge.md)

# 动机

* 由于某些类型的固有的实现逻辑，使得它们具有两个变化的维度，乃至多个维度的变化。
* 如何应对这种“多维度的变化”？如何利用面向对象技术来使得类型可以轻松地沿着两个乃至多个方向变化，而不引入额外的复杂度？

[代码：Bridge_bad ](../example/go/structural/Bridge/bad/bridge_bad.go)

[代码：Bridge_good ](../example/go/structural/Bridge/good/bridge_good.go)

[代码：Bridge_test ](../example/go/structural/Bridge/bridge_test.go)

# 模式定义

动态（组合）地给一个对象增加一些额外的职责。就增加功能而言，Decorator模式比生成子类（继承）更为灵活（消除重复代码 & 减少子类个数）。

<p align="right">——《设计模式》GOF</p>

# 结构



# 要点总结

* 通过采用组合而非继承的手法，Decorator模式实现了在**`运行时`**动态扩展对象功能的能力，而且可以根据需要扩展多个功能。避免了使用继承带来的“灵活性差”和“多子类衍生问题”。
* Decorator类在接口上表现为**`is-a`** Component的继承关系，即Decorator类继承了Component类所具有的接口。但在实现上又表现为**`has-a`** Component的组合关系，即Decorator类又使用了另外一个Component类。
* Decorator模式的目的并非解决“多子类衍生的多继承”问题，Decorator模式应用的要点在于解决“主体类在多个方向上的扩展功能”——是为了`“装饰”`的含义。

# 参考

* [23个设计模式](https://www.bilibili.com/video/BV1kW411P7KS?p=9&spm_id_from=pageDriver)

