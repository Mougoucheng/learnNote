<center><font size=10 color="blue"> Bridge [桥接模式] </font></center>

[GOF_23](../DesignPatterns.md)

# "单一职责"模式:

* 在软件组件的设计中，如果责任划分的不清晰，使用继承得到的结果往往是随着需求的变化，子类急剧膨胀，同时充斥着重复代码，这时候的关键是划清责任。
* 典型模式：
  * [Decorator](./gof_23/Decorator.md)
  * [Bridge](./gof_23/Bridge.md)

# 动机

* 在某些情况下我们可能会“过度地使用继承来扩展对象的功能”，由于继承为类型引入的静态特质，使得这种扩展方式缺乏灵活性；并且随着子类的增多（扩展功能的增多），各种子类的组合（扩展功能的组合）会导致更多子类的膨胀。
* 如何使“对象功能的扩展”能够根据需要来动态地实现？同时避免“扩展功能的增多”带来的子类膨胀问题？从而使得任何“功能扩展变化”所导致的影响降为最低？

[代码：Decorator_bad ](../example/go/structural/Decorator/bad/decorator_bad.go)

[代码：Decorator_good ](../example/go/structural/Decorator/good/decorator_good.go)

[代码：Decorator_test ](../example/go/structural/Decorator/decorator_test.go)

# 模式定义

动态（组合）地给一个对象增加一些额外的职责。就增加功能而言，Decorator模式比生成子类（继承）更为灵活（消除重复代码 & 减少子类个数）。

<p align="right">——《设计模式》GOF</p>

# 结构

![Decorator](D:\My\learnNote\design_patterns\gof_23\images\Decorator\structure.png)

# 要点总结

* 通过采用组合而非继承的手法，Decorator模式实现了在**`运行时`**动态扩展对象功能的能力，而且可以根据需要扩展多个功能。避免了使用继承带来的“灵活性差”和“多子类衍生问题”。
* Decorator类在接口上表现为**`is-a`** Component的继承关系，即Decorator类继承了Component类所具有的接口。但在实现上又表现为**`has-a`** Component的组合关系，即Decorator类又使用了另外一个Component类。
* Decorator模式的目的并非解决“多子类衍生的多继承”问题，Decorator模式应用的要点在于解决“主体类在多个方向上的扩展功能”——是为了`“装饰”`的含义。

# 参考

* [23个设计模式](https://www.bilibili.com/video/BV1kW411P7KS?p=9&spm_id_from=pageDriver)

