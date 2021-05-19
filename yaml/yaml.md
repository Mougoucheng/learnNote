# 一、简介

![img](http://www.ruanyifeng.com/blogimg/asset/2016/bg2016070403.png)

YAML 语言（发音 /ˈjæməl/ ）的设计目标，就是方便人类读写。它实质上是一种通用的数据串行化格式。

## 语法规则：

* **`大小写敏感`**
* **`使用缩进表示层级关系`**
* **`缩进时不允许使用Tab键，只允许使用空格`**
* **`缩进的空格数目不重要，只要形同层级的元素左侧对齐即可`**

## 注释

* **`#`**：表示注释，从这个字符一直到行尾，都会被解析器忽略

---

# 二、支持的数据结构

## 对象(映射/哈希/字典)

键值对的集合，又称为映射（mapping）/ 哈希（hashes）/ 字典（dictionary）

> 对象的一组键值对，使用冒号结构表示：
>
> ```yaml
> animal: pets
> ```
>
> 转为JSON如下：
>
> ```json
> {
>   "animal": "pets"
> }
> ```
>
> Yaml也允许另一种写法，将所有键值对写成一个行内对象：
>
> ```yaml
> hash: { name: Steve, foo: bar }
> ```
>
> 转为JSON如下：
>
> ```json
> {
>   "hash": {
>     "name": "Steve",
>     "foo": "bar"
>   }
> }
> ```
>
> 

## 数组

一组按次序排列的值，又称为序列（sequence）/ 列表（list）

> 一组连词线开头的行，构成一个数组：
>
> ```yaml
> - Cat
> - Dog
> - Goldfish
> ```
>
> 转为JSON如下：
>
> ```json
> [
>   "Cat",
>   "Dog",
>   "Goldfish"
> ]
> ```
>
> 数据结构的子成员是一个数组，则可以在该项下面缩进一个空格：
>
> ```yaml
> -
>  - Cat
>  - Dog
>  - Goldfish
> ```
>
> 转为JSON如下：
>
> ```json
> [
>   [
>     "Cat",
>     "Dog",
>     "Goldfish"
>   ]
> ]
> ```
>
> 数组也可以采用行内表示法：
>
> ```yaml
> animal: [Cat, Dog]
> ```
>
> 转为JSON如下：
>
> ```json
> {
>   "animal": [
>     "Cat",
>     "Dog"
>   ]
> }
> ```
>
> 

## 纯量（scalars）

最基本的、不可再分的值。以下数据类型都属于JavaScript的纯量:

> * 字符串
> * 布尔值
> * 整数
> * 浮点数
> * Null
> * 时间
> * 日期

### 数值

> 数值直接以字面量的形式表示：
>
> ```yaml
> number: 12.30
> ```
>
> 转为JSON如下：
>
> ```json
> {
>   "number": 12.3
> }
> ```



### 布尔值

> 布尔值用 **`true`** 和 **`false`** 表示：
>
> ```yaml
> isSet: true
> ```
>
> 转为JSON如下：
>
> ```json
> {
>   "isSet": true
> }
> ```

### Null

> **`null`** 用 **`~`** 表示：
>
> ```yaml
> parent: ~
> ```
>
> 转为JSON如下：
>
> ```json
> {
>   "parent": null
> }
> ```

### 时间

> 时间采用 ISO8601 格式，即时间和日期之间使用T连接，最后使用+代表时区
>
> ```yaml
> datetime: 2001-12-14t21:59:43.10-05:00
> ```
>
> 转为JSON如下：
>
> ```json
> {
> 	"datetime": "2001-12-15T02:59:43.100Z"
> }
> ```

### 日期

> 日期必须使用ISO 8601格式，即yyyy-MM-dd
>
> ```yaml
> date: 2018-02-17
> ```
>
> 转为JSON如下：
>
> ```json
> {
> 	"date": "2018-02-17T00:00:00.000Z"
> }
> ```

### 强制类型转换

> YAML允许使用两个感叹号，强制转换数据类型：
>
> ```yaml
> e: !!str 123
> f: !!str true
> ```
>
> 转为JSON如下：
>
> ```json
> {
> 	"e": "123",
> 	"f": "true"
> }
> ```

---

# 三、复合结构

对象和数组可以结合使用，形成复合结构。

```yaml
languages:
 - Ruby
 - Perl
 - Python
websites:
 YAML: yaml.org
 Ruby: ruby-lang.org
 Python: python.org
 Perl: use.perl.org
```

转为JSON如下。

```json
{
  "languages": [
    "Ruby",
    "Perl",
    "Python"
  ],
  "websites": {
    "YAML": "yaml.org",
    "Ruby": "ruby-lang.org",
    "Python": "python.org",
    "Perl": "use.perl.org"
  }
}
```

---

# 四、字符串

字符串是最常见，也是最复杂的一种数据结构。

## 1、默认不使用引号

> 字符串默认不使用引号表示：
>
> ```yaml
> str: 这是一行字符串
> ```
>
> 转为JSON如下：
>
> ```json
> {
> 	"str": "这是一行字符串"
> }
> ```

## 2、空格或特殊字符，需放在引号中

> 如果字符串之中包含空格或特殊字符，需要放在引号之中：
>
> ```yaml
> str: '内容： 字符串'
> ```
>
> 转为JSON如下：
>
> ```json
> {
> 	"str": "内容： 字符串"
> }
> ```

## 3、单引号会转义特殊字符，双引号不会

> 单引号和双引号都可以使用，单引号会对特殊字符转义，而双引号不会：
>
> ```yaml
> s1: '内容\n字符串'
> s2: "内容\n字符串"
> ```
>
> 转为JSON如下：
>
> ```json
> {
> 	"s1": "内容\\n字符串",
> 	"s2": "内容\n字符串"
> }
> ```

## 4、单引号中有单引号，需转义

> 单引号中如果还有单引号，必须连续使用两个单引号转义：
>
> ```yaml
> str: 'labor''s day'
> ```
>
> 转为JSON如下：
>
> ```json
> {
> 	"str": "labor's day"
> }
> ```

## 5、字符串分多行书写，第二行起需缩进

> 字符串可以写成多行，从第二行开始，必须有至少一个空格缩进。换行符会被转为空格：
>
> ```yaml
> str: 这是一段
>  多行
>  字符串
> ```
>
> 转为JSON如下：
>
> ```json
> {
> 	"str": "这是一段 多行 字符串"
> }
> ```

## 6、字符串分多行书写，换行符保留与折叠

> 多行字符串可以使用 `|` 保留换行符，也可以使用 `>` 折叠换行：
>
> ```yaml
> this: |
>   Foo
>   Bar
> that: >
>   Foo
>   Bar
> ```
>
> 转为JSON如下：
>
> ```json
> {
> 	"this": "Foo\nBar\n",
> 	"that": "Foo Bar\n"
> }
> ```

## 7、字符串分多行书写，文字块末尾换行符保留与删除

> `+` 表示保留文字块末尾的换行，`-` 表示删除字符串末尾的换行：
>
> ```yaml
> s1: |
>   Foo
> 
> s2: |+
>   Foo
> 
> 
> s3: |-
>   Foo
> 
> ```
>
> 转为JSON如下：
>
> ```json
> {
> 	"s1": "Foo\n",
> 	"s2": "Foo\n\n\n",
> 	"s3": "Foo"
> }
> ```

## 8、字符串中插入HTML标记

> 字符串之中可以插入HTML标记
>
> ```yaml
> message: |
> 
>   <p style="color: red">
>     段落
>   </p>
> ```
>
> 转为JSON如下：
>
> ```json
> {
> 	"message": "\n<p style=\"color: red\">\n  段落\n</p>\n"
> }
> ```

---

# 五、引用

**锚点 `&` 和别名 `*` ，可以用来引用**

## 理解：[`&`：定义别名，`*`：解引用]

## 映射引用

> ```yaml
> defaults: &defaults   #此处的引用名字可以不用defaults而改成其他，但是为了不混淆，还是一致比较好
>   adapter: postgres
>   host:    localhost
>   
> development:
>   database: myapp_development
>   <<: *defaults
>   
> test:
>   database: myapp_test
>   <<: *defaults
> ```
>
> 转为JSON如下：
>
> ```json
> {
> 	"defaults": {
> 		"adapter": "postgres",
> 		"host": "localhost"
> 	},
> 	"development": {
> 		"database": "myapp_development",
> 		"adapter": "postgres",
> 		"host": "localhost"
> 	},
> 	"test": {
> 		"database": "myapp_test",
> 		"adapter": "postgres",
> 		"host": "localhost"
> 	}
> }
> ```
>
> - `&` ：用来建立锚点（`defaults`）
> - `<<`： 表示合并到当前数据
> - `*`： 用来引用锚点 

## 数组引用

> ```yaml
> default: &default
>     - Mark McGwire
>     - Sammy Sosa
> hr: *default
> ```
>
> 转为JSON如下：
>
> ```json
> {
> 	"default": [
> 		"Mark McGwire",
> 		"Sammy Sosa"
> 	],
> 	"hr": [
> 		"Mark McGwire",
> 		"Sammy Sosa"
> 	]
> }
> ```
>
> - `&` ：用来建立锚点（`defaults`）
> - `*`： 用来引用锚点 

## 普通引用

> ```yaml
> - &showell Steve
> - Clark
> - Brian
> - Oren
> - *showell
> ```
>
> 转为JSON如下：
>
> ```json
> [
> 	"Steve",
> 	"Clark",
> 	"Brian",
> 	"Oren",
> 	"Steve"
> ]
> ```
>
> - `&` ：用来建立锚点（`defaults`）
> - `*`： 用来引用锚点 

# 六、函数和正则表达式的转换

这是 [JS-YAML](https://github.com/nodeca/js-yaml) 库特有的功能，可以把函数和正则表达式转为字符串。

> ```yaml
> # example.yml
> fn: function () { return 1 }
> reg: /test/
> ```

解析上面的 yml 文件的代码如下。

> ```javascript
> var yaml = require('js-yaml');
> var fs   = require('fs');
> 
> try {
>   var doc = yaml.load(
>     fs.readFileSync('./example.yml', 'utf8')
>   );
>   console.log(doc);
> } catch (e) {
>   console.log(e);
> }
> ```

从 JavaScript 对象还原到 yaml 文件的代码如下。

> ```javascript
> var yaml = require('js-yaml');
> var fs   = require('fs');
> 
> var obj = {
>   fn: function () { return 1 },
>   reg: /test/
> };
> 
> try {
>   fs.writeFileSync(
>     './example.yml',
>     yaml.dump(obj),
>     'utf8'
>   );
> } catch (e) {
>   console.log(e);
> }
> ```

# 参考：

- [YAML 语言教程-阮一峰的网络日志](https://www.ruanyifeng.com/blog/2016/07/yaml.html)
- [YAML 1.2规格 TODO](https://yaml.org/spec/1.2/spec.html)
- [YAML from Wikipedia TODO](https://en.wikipedia.org/wiki/YAML)
- [YAML for Ruby TODO](http://yaml.org/YAML_for_ruby.html)

