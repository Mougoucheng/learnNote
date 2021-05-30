# Docker 概述

## Docker为什么出现

一款产品：开发---上线   两套环境：应用环境，应用配置！

开发----运维。

问题：

- 我在我的电脑上可以运行，环境不一致。
- 版本更新，导致服务不可用。
- 对运维来说，要求高，考验大。
- 环境配置十分麻烦，每个及其都要部署环境（集群Redis、ES、Hadoop……），费时费力。

那能不能 `发布一个项目（jar + （Redis MySQL jdk ES））,项目能不能都带上环境安装打包？`

之前在服务器配置一个应用的环境：Redis MySQL ES Hadoop，配置超麻烦，不能跨平台。

window开发，最后发布到Linux

传统：程序开发Jar，其他运维来做。

现在：开发打包部署上线，一套流程做完。



Java---apk---发布（应用商店）---张三使用apk---安装即可用！

Java---Jar(环境)---打包项目带上环境（镜像）---（上传Docker仓库：商店）---下载我们发布的镜像---直接运行即可！

​	

Docker给以上问题，提出了解决方案！

![image-20210421201950536](/Users/yaoxing.chen/Library/Application Support/typora-user-images/image-20210421201950536.png)

Docker的思想就来自集装箱

隔离：Docker核心思想，打包装箱，每个箱子都是互相隔离的，通过这种机制可以把服务器利用到极致。



## Docker的历史

Docker的出现：

- 2010年，几个搞IT的年轻人，在美国成立了 `dotCloud`，做一些PaaS的云计算服务（LXC有关的容器技术），他们将自己的技术（容器化技术）命名为Docker，刚诞生的时候没有引起行业的注意。dotCloud，就活不下去。
- 2013年，Docker开源。然后越来越多的人发现了Docker的优点，就火了。
- 2014年4月9日，Docker1.0发布。

虚拟化技术：

- Docker：十分轻巧，几MB，秒级启动

- 虚拟机技术：很笨重，几个GB，分钟级启动

Docker是基于Go开发的开源项目。

官网：https://www.docker.com/

文档：https://docs.docker.com/  Docker的文档是超级详细的，多看看

仓库：https://hub.docker.com/ 

## Docker能干嘛

![image-20210421205733053](/Users/yaoxing.chen/Library/Application Support/typora-user-images/image-20210421205733053.png)

**传统部署时代：**

早期，各个组织机构在物理服务器上运行应用程序。无法为物理服务器中的应用程序定义资源边界，这会导致资源分配问题。 例如，如果在物理服务器上运行多个应用程序，则可能会出现一个应用程序占用大部分资源的情况， 结果可能导致其他应用程序的性能下降。 一种解决方案是在不同的物理服务器上运行每个应用程序，但是由于资源利用不足而无法扩展， 并且维护许多物理服务器的成本很高。

**虚拟化部署时代：**

作为解决方案，引入了虚拟化。虚拟化技术允许你在单个物理服务器的 CPU 上运行多个虚拟机（VM）。 虚拟化允许应用程序在 VM 之间隔离，并提供一定程度的安全，因为一个应用程序的信息 不能被另一应用程序随意访问。

虚拟化技术能够更好地利用物理服务器上的资源，并且因为可轻松地添加或更新应用程序 而可以实现更好的可伸缩性，降低硬件成本等等。

每个 VM 是一台完整的计算机，在虚拟化硬件之上运行所有组件，包括其自己的操作系统。

**容器部署时代：**

容器类似于 VM，但是它们具有被放宽的隔离属性，可以在应用程序之间共享操作系统（OS）。 因此，容器被认为是轻量级的。容器与 VM 类似，具有自己的文件系统、CPU、内存、进程空间等。 由于它们与基础架构分离，因此可以跨云和 OS 发行版本进行移植。

容器因具有许多优势而变得流行起来。下面列出的是容器的一些好处：

- 敏捷应用程序的创建和部署：与使用 VM 镜像相比，提高了容器镜像创建的简便性和效率。
- 持续开发、集成和部署：通过快速简单的回滚（由于镜像不可变性），支持可靠且频繁的 容器镜像构建和部署。
- 关注开发与运维的分离：在构建/发布时而不是在部署时创建应用程序容器镜像， 从而将应用程序与基础架构分离。
- 可观察性不仅可以显示操作系统级别的信息和指标，还可以显示应用程序的运行状况和其他指标信号。
- 跨开发、测试和生产的环境一致性：在便携式计算机上与在云中相同地运行。
- 跨云和操作系统发行版本的可移植性：可在 Ubuntu、RHEL、CoreOS、本地、 Google Kubernetes Engine 和其他任何地方运行。
- 以应用程序为中心的管理：提高抽象级别，从在虚拟硬件上运行 OS 到使用逻辑资源在 OS 上运行应用程序。
- 松散耦合、分布式、弹性、解放的微服务：应用程序被分解成较小的独立部分， 并且可以动态部署和管理 - 而不是在一台大型单机上整体运行。
- 资源隔离：可预测的应用程序性能。
- 资源利用：高效率和高密度。



> DevOps（开发、运维）
>
> **应用更快速的交付和部署**
>
> - 传统：一堆帮助文档，安装程序
>
> - Docker：打包镜像发布测试，一键运行
>
> **更便捷的升级和扩缩容**
>
> - 使用Docker之后，我们部署应用就和搭积木一样
>
> - 项目打包为一个镜像，扩展  服务器A、服务器B
>
> **更简单的系统运维**
>
> - 在容器化之后，我们的开发、测试环境都是一致的
>
> **更高效的计算资源利用**
>
> - Docker是内核级别的虚拟化，可以在一个物理机上运行很多的容器实例，服务器的性能可以被压榨到极致

# Docker 安装

## Docker的组成

![第2 章容器架构- 007 - Docker 架构详解- osc_uh2p1ff2的个人空间- OSCHINA - 中文开源技术交流社区](https://oscimg.oschina.net/oscnet/6980be9f402ffabf9c044b8170630006086.png)

**镜像（image）**：

​		Docker镜像就相当于一个root文件系统，如同一个模板/类，可以通过这个模板/类来创建容器服务(实例)，如：tomcat镜像==》run镜像==》tomcat01容器(提供服务器)。通过镜像可以创建多个容器（最终服务/项目就运行在容器中）。

**容器（container）**：

​		镜像和容器多关系，就像是面向对象程序设计中的类和实例一样，镜像是静态的定义(只读)，容器是镜像运行时的实体。容器可以被创建、启动、停止、暂停等。

**仓库（repository）**：

​		仓库可看成一个代码控制中心，用来保持镜像。仓库分为共有仓库、私有仓库。Docker Hub（默认是国外的），阿里云……都有容器服务器(假如拉取镜像到本地比较慢时，可以配置镜像加速)。

## Docker的安装

> 环境准备

1、需要会一些Linux基础

2、CentOS 7

3、使用Xshell连接远程服务器进行操作

> 环境查看

```shell
# 系统内核是 3.10 以上的
[root@yaoxing.chen]# uname -r

# 系统版本
[root@yaoxing.chen]# cat /etc/os-release

```

> 安装

根据帮助文档，逐步安装(Install On CentOS):

```shell
#1、卸载旧的版本
yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine
#2、需要的安装包
yum install -y yum-utils

#3、设置镜像的仓库
yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo #默认是国外的
    
yum-config-manager \
    --add-repo \
    http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo #推荐使用阿里云镜像

#4、更新yum软件包索引
yum makecache fast

#4、安装docker相关的内容 docker-ce：社区版  docker-ee：企业版
yum install docker-ce docker-ce-cli containerd.io

#5、启动docker
systemctl start docker

#6、使用docker version 查看是否安装成功

```

![image-20210422143850580](/Users/yaoxing.chen/Library/Application Support/typora-user-images/image-20210422143850580.png)



```shell
#7、hello-world程序
docker run hello-world
```

![image-20210422143359322](/Users/yaoxing.chen/Library/Application Support/typora-user-images/image-20210422143359322.png)

```shell
#8、查看一下下载的这个hello-world镜像
```

![image-20210422144009578](/Users/yaoxing.chen/Library/Application Support/typora-user-images/image-20210422144009578.png)

了解：卸载docker

```shell
#1、卸载依赖
yum remove docker-ce docker-ce-cli containerd.io

#2、删除资源
rm -rf /var/lib/docker 
rm -rf /var/lib/containerd
# /var/lib/docker docker的默认工作路径
```



## 阿里云镜像加速

1、登陆阿里云找到容器服务

2、找到镜像加速地址

![image-20210422150737192](/Users/yaoxing.chen/Library/Application Support/typora-user-images/image-20210422150737192.png)

3、配置使用

```shell
sudo mkdir -p /etc/docker

sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://ax0ikzv1.mirror.aliyuncs.com"]
}
EOF

sudo systemctl daemon-reload

sudo systemctl restart docker
```

## Hello-world启动流程

![image-20210422143359322](/Users/yaoxing.chen/Library/Application Support/typora-user-images/image-20210422143359322.png)

![image-20210422162201291](/Users/yaoxing.chen/Library/Application Support/typora-user-images/image-20210422162201291.png)

## 底层原理

**Docker是怎么工作的？**

docker是一个Client-Server结构的系统，Docker的守护进程运行在主机上，通过Socket从客户端访问，DockerServer接收到DockerClient的命令，就会执行这个命令。

![image-20210422164549408](/Users/yaoxing.chen/Library/Application Support/typora-user-images/image-20210422164549408.png)

Docker为什么比VM快？

1、Docker有着比虚拟机更少的抽象层。

2、Docker利用的是宿主机的内核，vm需要Guest OS。

![image-20210421205733053](/Users/yaoxing.chen/Library/Application Support/typora-user-images/image-20210421205733053.png)

所以说，新建一个容器的时候，Docker利用宿主机的操作系统，不需要像虚拟机一样重新加载一个操作系统内核，避免引导，秒级启动。虚拟机是加载Guest OS，需要引导加载，分钟级别的。

![image-20210422170038242](/Users/yaoxing.chen/Library/Application Support/typora-user-images/image-20210422170038242.png)





# Docker 常用命令

## 帮助命令

```shell
docker version			# 显示docker的版本信息
docker info					# 显示docker的系统信息，包括镜像、容器数量
docker help					# 显示docker帮助信息，包括所有选项、管理命令、普通命令
docker 命令 --help   # 万能命令，查看某个命令的用法
```

帮助文档地址：https://docs.docker.com/engine/reference/commandline/build/

## 镜像命令

### docker images：查看镜像

```shell
# Usage:  docker images [OPTIONS] [REPOSITORY[:TAG]]
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker images		# 查看本机上的所有镜像
REPOSITORY    TAG       IMAGE ID       CREATED        SIZE
busybox       latest    388056c9a683   2 weeks ago    1.23MB
hello-world   latest    d1165f221234   6 weeks ago    13.3kB
alpine/git    latest    a939554ad0d0   2 months ago   25.1MB

# 解释：
REPOSITORY		镜像的仓库源
TAG						镜像的标签
IMAGE ID			镜像的id
CREATED				镜像的创建时间
SIZE					镜像的大小

# 可选项
	-a, --all							# 显示所有镜像
  -q, --quiet           # 只显示所有镜像的ID
```

### docker search ：搜索镜像

```shell
# Usage:  docker search [OPTIONS] TERM
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker search mysql
NAME                              DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
mysql                             MySQL is a widely used, open-source relation…   10777     [OK]       
mariadb                           MariaDB Server is a high performing open sou…   4058      [OK]       
mysql/mysql-server                Optimized MySQL Server Docker images. Create…   793                  [OK]
# 可选项
  -f, --filter filter   # 基于给定的filter条件过滤搜索结果
  -f STARS=3000 				# 搜索出来的镜像STAR大于3000
  -f=STARS=3000 				# 搜索出来的镜像STAR大于3000
  --filter STARS=3000 				# 搜索出来的镜像STAR大于3000
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker search mysql -f=STARS=3000
NAME      DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
mysql     MySQL is a widely used, open-source relation…   10777     [OK]       
mariadb   MariaDB Server is a high performing open sou…   4058      [OK]       
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker search mysql -f STARS=3000
NAME      DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
mysql     MySQL is a widely used, open-source relation…   10777     [OK]       
mariadb   MariaDB Server is a high performing open sou…   4058      [OK]       
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker search mysql --filter STARS=3000
NAME      DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
mysql     MySQL is a widely used, open-source relation…   10777     [OK]       
mariadb   MariaDB Server is a high performing open sou…   4058      [OK]  
```

### docker pull：下载镜像

```shell
# Usage:  docker pull [OPTIONS] NAME[:TAG|@DIGEST]
# 下载镜像 docker pull 镜像名[:tag]
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker pull mysql
Using default tag: latest								# 如果不写tag，默认latest
latest: Pulling from library/mysql
f7ec5a41d630: Pull complete 						# 分层下载，docker image的核心，联合文件系统
9444bb562699: Pull complete 
6a4207b96940: Pull complete 
181cefd361ce: Pull complete 
8a2090759d8a: Pull complete 
15f235e0d7ee: Pull complete 
d870539cd9db: Pull complete 
493aaa84617a: Pull complete 
bfc0e534fc78: Pull complete 
fae20d253f9d: Pull complete 
9350664305b3: Pull complete 
e47da95a5aab: Pull complete 
Digest: sha256:04ee7141256e83797ea4a84a4d31b1f1bc10111c8d1bc1879d52729ccd19e20a		# 签名(防伪)
Status: Downloaded newer image for mysql:latest
docker.io/library/mysql:latest					# mysql对应的真实地址

# 以下两个命令说等价的
docker pull mysql
docker pull docker.io/library/mysql:latest

# 指定版本下载
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker pull mysql:5.7
5.7: Pulling from library/mysql
f7ec5a41d630: Already exists 
9444bb562699: Already exists 
6a4207b96940: Already exists 
181cefd361ce: Already exists 
8a2090759d8a: Already exists 
15f235e0d7ee: Already exists 
d870539cd9db: Already exists 
cb7af63cbefa: Pull complete 
151f1721bdbf: Pull complete 
fcd19c3dd488: Pull complete 
415af2aa5ddc: Pull complete 
Digest: sha256:a655529fdfcbaf0ef28984d68a3e21778e061c886ff458b677391924f62fb457
Status: Downloaded newer image for mysql:5.7
docker.io/library/mysql:5.7
```

![image-20210422195321444](/Users/yaoxing.chen/Library/Application Support/typora-user-images/image-20210422195321444.png)

### Docker rmi：删除镜像

```shell
# Usage:  docker rmi [OPTIONS] IMAGE [IMAGE...]
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker rmi -f 镜像ID						# 删除指定的镜像
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker rmi -f 镜像ID1 镜像ID2	 # 删除多个镜像
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker rmi -f $(docker images -aq) # 删除全部的容器，其中$(docker images -aq)为参数传递，递归删除

# 可选项
 -f, --force      # 强制删除镜像
```

## 容器命令

**说明：有了镜像才能创建容器，且需要将镜像从远程仓库下载到本地才能创建**

下载一个Cent OS镜像来测试学习：

```shell
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker pull centos
Using default tag: latest
latest: Pulling from library/centos
7a0437f04f83: Pull complete 
Digest: sha256:5528e8b1b1719d34604c87e11dcd1c0a20bedf46e83b5632cdeac91b8c04efc1
Status: Downloaded newer image for centos:latest
docker.io/library/centos:latest
```

### docker run：创建并启动容器

```shell
# Usage:  docker run [OPTIONS] IMAGE [COMMAND] [ARG...]

# 参数说明
--name="name"			为新建的容器制定名字
-d								后台方式运行
-it								使用交互方式运行，进入容器查看内容
-p								(小写)指定容器的端口  -p 8080:8080
			-p ip:宿主机端口:容器端口
			-p 宿主机端口:容器端口(常用)
			-p 容器端口
			容器端口
-P								(大写)随机指定端口

# 测试：启动并进入容器
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker run -it centos /bin/bash
[root@6eeaf1c16314 /]# ls					# 查看容器内的centos，基础版本，很多命令都是不完善的
bin  dev  etc  home  lib  lib64  lost+found  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var

# 测试：退出容器到宿主机，容器停止 exit
[root@6eeaf1c16314 /]# exit
exit
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ ls
Accounts		Dictionaries		LaunchAgents		Saved Application State
Application Scripts	Family			Logs			Screen Savers

# 测试：退出容器到宿主机，容器不停止 ctrl+P+Q
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker run -it centos /bin/bash
[root@08397955f2c5 /]#			# 同时按下ctrl+P+Q
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker ps
CONTAINER ID   IMAGE     COMMAND       CREATED          STATUS          PORTS     NAMES
08397955f2c5   centos    "/bin/bash"   15 seconds ago   Up 15 seconds             affectionate_franklin
```

### docker ps：查看容器

```shell
# Usage:  docker ps [OPTIONS]
					(空选项)列出当前正在运行的容器
-a				列出当前正在运行的容器+历史运行过的容器
-n=num		列出最近运行过的num个容器
-q				只显示容器的ID

yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker ps -a
CONTAINER ID   IMAGE          COMMAND       CREATED          STATUS                       PORTS     NAMES
6eeaf1c16314   centos         "/bin/bash"   11 minutes ago   Exited (127) 8 minutes ago             cranky_edison
75dd9d53c179   d1165f221234   "/hello"      6 hours ago      Exited (0) 6 hours ago                 beautiful_lalande
b47d12c249d7   busybox        "sh"          7 hours ago      Exited (127) 7 hours ago               adoring_galileo
cf6e373fd068   busybox        "sh"          30 hours ago     Exited (0) 30 hours ago                compassionate_curie
18080301b933   busybox        "/bin/bash"   30 hours ago     Created                                affectionate_nobel
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker ps -n=2
CONTAINER ID   IMAGE          COMMAND       CREATED          STATUS                       PORTS     NAMES
6eeaf1c16314   centos         "/bin/bash"   11 minutes ago   Exited (127) 8 minutes ago             cranky_edison
75dd9d53c179   d1165f221234   "/hello"      6 hours ago      Exited (0) 6 hours ago                 beautiful_lalande
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker ps -aq
6eeaf1c16314
75dd9d53c179
b47d12c249d7
cf6e373fd068
18080301b933
yaoxingchendeMacBook-Pro:Library yaoxing.chen$ docker ps -aq -n=2
6eeaf1c16314
75dd9d53c179
```

### docker rm：删除容器

```shell
# Usage:  docker rm [OPTIONS] CONTAINER [CONTAINER...]
docker rm 容器ID							# 删除指定容器，不能删除正在运行的容器，如果要强制删除则docker rm -f 容器ID
docker rm -f $(docker ps -aq)			# 删除所有的容器
docker ps -aq |xargs docker rm		# 删除所有的容器(使用管道传参)
```

### docker start：启动容器

```shell
# Usage:  docker start [OPTIONS] CONTAINER [CONTAINER...]
docker
```



### docker restart：重启容器

### docker stop：停止容器

### docker kill：停止容器(杀死)

# Docker 镜像

# Docker 数据卷

# DockerFile

# Docker网络原理

# Docker IDEA整合

# Docker Compose

按照docker-compose.yaml文件，启动一组容器，组成一个完整应用对外提供服务（如web+mysql，组成博客）

# Docker Swarm

容器编排

# CI/CD Jenkins

```go
import log
func main(){
  log.Println("hello world")
}
```

```shell
#
yaoxingchendeMacBook-Pro:~ yaoxing.chen$ docker ps -a
CONTAINER ID   IMAGE     COMMAND       CREATED          STATUS                      PORTS     NAMES
cf6e373fd068   busybox   "sh"          14 minutes ago   Exited (0) 14 minutes ago             compassionate_curie
18080301b933   busybox   "/bin/bash"   16 minutes ago   Created                               affectionate_nobel
```

