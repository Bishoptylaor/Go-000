## 一、工程项目结构

### 1. 标准go项目模式(Standard Go Project Layout)

1. Go目录
- /cmd
项目主要的应用程序。

对于每个应用程序来说这个目录的名字应该和项目可执行文件的名字匹配（例如，/cmd/myapp）。

不要在这个目录中放太多的代码。如果目录中的代码可以被其他项目导入并使用，那么应该把他们放在/pkg目录。如果目录中的代码不可重用，或者不希望被他人使用，应该将代码放在/internal目录。显示的表明意图比较好！

通常来说，项目都应该拥有一个小的main函数，并在main函数中导入或者调用/internal和/pkg目录中的代码。

- /internal
私有的应用程序代码库。这些是不希望被其他人导入的代码。请注意：这种模式是Go编译器强制执行的。详细内容情况Go 1.4的release notes。再次注意，在项目的目录树中的任意位置都可以有internal目录，而不仅仅是在顶级目录中。

可以在内部代码包中添加一些额外的结构，来分隔共享和非共享的内部代码。这不是必选项（尤其是在小项目中），但是有一个直观的包用途是很棒的。应用程序实际的代码可以放在/internal/app目录（如，internal/app/myapp），而应用程序的共享代码放在/internal/pkg目录（如，internal/pkg/myprivlib）中。

- /pkg
外部应用程序可以使用的库代码（如，/pkg/mypubliclib）。其他项目将会导入这些库来保证项目可以正常运行，所以在将代码放在这里前，一定要三四而行。请注意，internal目录是一个更好的选择来确保项目私有代码不会被其他人导入，因为这是Go强制执行的。使用/pkg目录来明确表示代码可以被其他人安全的导入仍然是一个好方式。
一般按照功能来分类，用于项目内跨多个应用的公共共享代码。感觉可以类比marketutil
- /vendor
应用程序的依赖关系（通过手动或者使用喜欢的依赖管理工具，如新增的内置Go Modules特性）。执行go mod vendor命令将会在项目中创建/vendor目录，注意，如果使用的不是Go 1.14版本，在执行go build进行编译时，需要添加-mod=vendor命令行选项，因为它不是默认选项。

构建库文件时，不要提交应用程序依赖项。

请注意，从1.13开始，Go也启动了模块代理特性（使用https：//proxy.golang.org作为默认的模块代理服务器）。点击这里阅读有关它的更多信息，来了解它是否符合所需要求和约束。如果Go Module满足需要，那么就不需要vendor目录。

2. 服务端应用程序的目录
- /api

3. Web应用程序的目录
- /web

4. 通用应用程序的目录
- /configs
配置文件模板或默认配置。

将confd或者consul-template文件放在这里。

- /init
系统初始化（systemd、upstart、sysv）和进程管理（runit、supervisord）配置。
- /scripts
用于执行各种构建，安装，分析等操作的脚本。
- /build
打包和持续集成。

将云（AMI），容器（Docker），操作系统（deb，rpm，pkg）软件包配置和脚本放在/build/package目录中。

将CI（travis、circle、drone）配置文件和就脚本放在build/ci目录中。请注意，有一些CI工具（如，travis CI）对于配置文件的位置有严格的要求。尝试将配置文件放在/build/ci目录，然后链接到CI工具想要的位置。

- /deployments
IaaS，PaaS，系统和容器编排部署配置和模板（docker-compose，kubernetes/helm，mesos，terraform，bosh）。请注意，在某些存储库中（尤其是使用kubernetes部署的应用程序），该目录的名字是/deploy
- /test
外部测试应用程序和测试数据。随时根据需要构建/test目录。对于较大的项目，有一个数据子目录更好一些。例如，如果需要Go忽略目录中的内容，则可以使用/test/data或/test/testdata这样的目录名字。请注意，Go还将忽略以“.”或“_”开头的目录或文件，因此可以更具灵活性的来命名测试数据目录。

5. 其他

### 2. Kit Project Layout

### 3. Service Application Project
1. 一般分为4类：interface，service，job，admin，task
- interface: 对外的 BFF 服务，接受来自用户的请求，比如暴露了 HTTP/gRPC 接口。 可以理解为我们正在用的pub project
- service: 对内的微服务，仅接受来自内部其他服务或者网关的请求，比如暴露了gRPC 接口只对内服务。
- admin：区别于 service，更多是面向运营测的服务，通常数据权限更高，隔离带来更好的代码级别安全。
- job: 流式任务处理的服务，上游一般依赖 message broker。
- task: 定时任务，类似 cronjob，部署到 task 托管平台中。










## 二、api设计
## 三、配置管理
## 四、包管理
## 五、测试
## 六、References