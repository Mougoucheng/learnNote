查找不到本地包导入：

go mod init test

go mod tidy

go build 



建go项目

go mod init 自定义包名

go mod tidy

启用go 模块集成（go module）

e.g:

go mod init GitHub.com/vangoleo/hello

---

http.request的body只能读一次，如果request请求不是表单Form，可以不用parseForm



Chrome浏览器http链接测试插件：Talend API Tester



postman：调试工具(http请求、css、HTML、脚本等)



speedtest.cn

supervisor、

tmpwatch

---

1、了解Elastic Search

2、ssh 原理？：

​	cat ~/.ssh/config

​	cat ~/.ssh/authorized_keys

3、技术栈：grpc / grpc-gateway / yaml / gorm / jwt Redis, mongodb

redis:队列、pubsub、zset时间复杂度，各个命令的时间复杂度

mongodb：aagregate - X ， map reduce - X（其他的增删改查，索引，分片）



4、go.mod文件中记录项目中使用到的库，了解他们是干什么用的

5、git config的一些配置

[pull]

​	rebase = true

[merge]

​	ff = false

[push]

​	default = simple

[alias]

​	st = status

​	ci = commit

​	co = checkout

​	br = branch

​	sur = submodule update --recursive



6、Mac是一个unix like的系统，需要了解熟悉的相关内容

$PATH

which

.bash_rc

.bash_profile

.profile



source xxx

chmod +x

pyenv, npm, 



7、vuejs页面、npm(node package manager，前端代码管理工具)



8、alpha：给策划配置联调

   beta：QA测试

   prod：线上

---

Go grpc. Docker k8s   yaml



Redis, 

mongodb

Postgres





ssh beta-word		远程连接beta-word服务器

ssh配置：vi ~/.ssh/config



ps -ef|grep event	查看进程

scp ids.txt ec2-user@beta-word:~/yaoxing

scp 文件 用户名@服务器IP：存放路径	（复制到远程，从远程复制到本地呢？）

nohup ./eventTool --file ./bq-results-20210423-131433-mujwx0y2r0fm.csv --event start_level_27 --appid 469292470183058 > level_evt.log & 

sudo vi /etc/hosts ：修改Host文件

mac命令：networksetup 选项

networksetup -listallhardwareports：列出所有硬件端口 

command+k:弹出“连接服务器”窗口

command+shift+G：弹出“前往文件夹”窗口

pkill -9 -f zIdCardServerTest：pkill不生效时，是由于进程名超过了15个字符，使用-f选项来避免这种情况。参考：http://manpages.ubuntu.com/manpages/precise/en/man1/pkill.1.html

---

TODO:

1、提升主动性（关注进度与问题）

2、工程能力提升

3、独当一面（更多的知识面学习）

4、注意时间和进度

5、代码能力



bootstrap项目参考？

移动平均率算法？

版本标识？alpha，beta，？

日志分级？

context收集数据、git更换源站



watchdog

防抖动、防重复、多次重复警报升级

http,other探测请求，



http跨域问题：

​	w.Header().Set("Access-Control-Allow-Origin", "*")

​	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")



https://github.com/rs/cors库





TODO：

struct中零值和空问题(结构体嵌套最好使用指针，这样指针为nil时，json.Marshal可以省略nil指针部分的数据转换，节省传输带宽)

依赖注入

​			

域名、

http服前面挡一层Nginx



errcode不为0，就算失败

TODO：fakeRsp，不直接用json字符串，用结构体赋值再Marshal？？？

收集collections，1分钟1发/满128发



自动补齐下线的消息：

1、上线补齐no、si、bt、ot、ct、[di、pi]由客户端给

2、map记录已上线pi/di，再次收到collection时查询map，有记录则更新在线时间，并丢弃collection

3、时间轮询map，对过期记录构造发送下线collection

4、定时输出Server状态日志信息。



5、压力测试？

6、部署脚本优化

7、ssh无反应

---

常用工具：

​	iTerm			终端

​	oh-my-zsh		

​	sublime         	打开大文件用

​	lastpass			chrome的插件

​	alfred			效率神器（替代spotlight）

​	xcode			代码编辑器

​	sourcetree		代码版本管理工具（由于sourcetree上做操作可能会出问题，主要是用来查看，各个操作还是用命令行）

​	homebrew		包管理器（和apt-get、yum、wget之类的类似）

---

go get 流程

1、https访问url，

2、拿repo的sum和sum.golang.org上注册的sum进行check操作，防止错误/被劫持

详见：http://note.youdao.com/s/93xSfc52



go get配置私有库zenLog:

1、~/.gitconfig配置新加

[url "ssh://git@github.com/Zentertain/"]

​    insteadOf = https://github.com/Zentertain/

用于解决go get流程中的第一步，即使用https访问https://github.com/Zentertain/时，会替换成使用ssh访问ssh://git@github.com/Zentertain/，就解决了GitHub上私有仓库不能访问的问题



2、设置环境变量：go env -w GOPRIVATE="github.com/Zentertain"

指定github.com/Zentertain下的仓库，都不走代理和sum校验。设置GOPRIVATE后后，GONOPROXY、GONOSUMDB也会自动填上对应值

GONOPROXY="github.com/Zentertain"

GONOSUMDB="github.com/Zentertain"



如何查看repo发布使用的端口？

---

1、安装protoc-gen-go：brew install protocols-gen-go

​	proto文件生成对应的go代码文件：protoc —go_out=$OUT_DIR $protoFile

2、安装protoc-gen-go-grpc：brew install protocols-gen-go-grpc

​	proto文件生成对应的grpc go代码文件：protoc —go-grpc_out=$OUT_DIR $protoFile







go get 下载安装需要使用的包

1、go.mod文件路径下执行：go get github.com/gogf/gf

2、go mod download ，此时gogf包会安装在GOPATH下的





TODO:不用在管testGovURL了，回头有空单独弄一个testCase，专门用来有需要时过test

代码优化总结：

1、好代码基本可以自描述，某些踩坑的地方/不易懂的地方才需要做注释

2、注意闭包参数的变化，使用到外部的变量最好是作为参数传入进去

3、尽量不使用init()函数，因为包导入的时候会调用，而我们不知道什么时候/哪里会导入导致init()调用

4、结构体尽量拆出来一个model.go，只在包内部使用的struct，首字母记得小写

5、只在某个地方用的全局变量，最好定义在离使用它最近的地方

6、只从一个channel循环收取数据时，for内部没必要用select，直接data := <-channel 让其阻塞收取

7、多个errCode直接拆出一个errors.go，错误码和对应的错误信息从其内获取

8、只在内部使用的方法，首字母记得小写

9、处理所有error返回，哪怕用断言panic(

e.g: func Ast2(notCare interface{}, err error) {

​	 	if err != nil {

​			panic(err)

​		}

​	}

)，

只有特殊情况下不做处理(defer r.Body.Close()等，关闭失败不影响)

10、更多的注意细节，得知道它是干什么用的，然后再自行实现

11、每个函数是一个完整的功能流程，代码语义表述上应该清晰明了。如果代码较长，应该把更细节的操作/代码拆出另一个单独的函数，通过函数名和在高一层级的函数调用的，达到函数功能表述清晰的目的

---

分层、可维护、private repo多分支（正式、开发、α、β…）、说明文档、部署脚本、保证持续集成的CI/CD测试用例



工程化流程：

1、基础代码版本开发

2、调试运行，建立repo版本管理

3、说明文档

4、部署脚本

5、test case：（如果够详细，甚至可以不用说明文档）

6、代码层级结构划分（有哪些？

​	代码重构40种方法

​	设计模式（动手写练习）



​	bootstrap				(干嘛用的？)

​	project-layout			(非标准)

​	api test				(goconvey测试框架)

​	日志					(log.fmt？

​	数据库

​	监控

---

https://goproxy.io/zh/一个为Go模块而生的全球代理

在cmd中设置

如果您使用的 Go 版本是 1.13 及以上 （推荐）
    go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.cn,direct


    如果将go env -w GO111MODULE=on设为on, 则会使用1.13的mod包特性,下载的包不会在src目录下


原始默认值：
GO111MODULE=""
GOPROXY="https://proxy.golang.org,direct"


【注意】：
[Go]开启go module情况下go get后文件下载存放目录


在开启了go module情况下 , 也就是执行了 

go env -w GO111MODULE=on


再执行go get xxxxx

文件会存放在$GOPATH/pkg/mod 下


并且目录名里的大写字母会转成小写字母 , 前面加一个!


例如:

go get github.com/JamesLMilner/pip-go

文件存放在:

/Users/yaoxing.chen/go/pkg/mod/github.com/!james!l!milner

---

go mod引用笔记：https://note.youdao.com/ynoteshare1/index.html?id=e108668f369cb5c670b34e65e2f24609&type=note