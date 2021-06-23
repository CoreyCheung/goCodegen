# Golang 代码生成工具项目

**此项目专注于从Mysql数据库生成Java（有包装的mybatis，免配置文件方式）、Golang（xorm或gorm）、Python（SQLAlchemy）的相应数据实体类代码**

Golang开发环境安装和环境变量配置见 [此文档](http://doc.wanpinghui.com/pages/viewpage.action?pageId=3506299)

与Java、Python、NodeJS、PHP等均不同，只是开发环境需要安装用于代码编译，生产环境直接运行项目编译好的可执行文件即可，无需安装任何类似JRE、Python、v8、PHP的虚拟机或运行时环境！

执行以下操作前请理解并确认：
1. GOPATH 环境变量已配置，例如配置到用户目录下的go目录
2. GOPATH 目录是用于存放golang项目和其相关依赖的目录，所有golang项目代码都应该位于GOPATH目录中的src子目录下的包括代码托管地址、组织名、项目名在内的多级子目录下

## 获取项目
   
### 获取

> go get -v -insecure -d git.wanpinghui.com/WPH/go_codegen

参数：-v 显示详情，-insecure 非https版本仓库路径，-d 仅获取，不编译安装到GOPATH

可能会卡住或者报错，不用管，只要项目本身源代码文件获取下来了即可，文件保存位置即在上述GOPATH目录中的子目录里

**完成上述操作之后，以后日常开发都可以正常使用git命令来获取和提交此项目代码**

### 建立软连接

建立软连接的目的仅仅是为了方便找而已！

使用GoLand等开发工具时，仍然打开GOPATH里面的项目目录，而不是打开软连接位置的！

**Windows** 

Power Shell 或 Windows CMD (需要用管理员身份运行)：

**注意，Windows Power Shell打开后首先输入cmd进入Windows CMD，下同不再累述**

> mklink /D %USERPROFILE%\Documents\project\wph\go_codegen %GOPATH%\src\git.wanpinghui.com\WPH\go_codegen

**MAC OS X**

> sudo ln -s $GOPATH/src/git.wanpinghui.com/WPH/go_codegen/ ~/Documents/project/wph/go_codegen

## 相关依赖

### Golang官方包

Golang运行时 >= 1.10.3

其他常用golang官方包，因为golang.org被墙，所以用以下步骤获取：

**Windows**

```cmd
git clone https://github.com/golang/net.git %GOPATH%\src\golang.org\x\net
git clone https://github.com/golang/text.git %GOPATH%\src\golang.org\x\text
git clone https://github.com/golang/tools.git %GOPATH%\src\golang.org\x\tools
git clone https://github.com/golang/sys.git %GOPATH%\src\golang.org\x\sys
git clone https://github.com/golang/crypto.git %GOPATH%\src\golang.org\x\crypto
cd %GOPATH%\src
go install golang.org\x\text
```

**MAC OS X**

```bash
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net \
&& git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text \
&& git clone https://github.com/golang/tools.git $GOPATH/src/golang.org/x/tools \
&& git clone https://github.com/golang/sys.git $GOPATH/src/golang.org/x/sys \
&& git clone https://github.com/golang/crypto.git $GOPATH/src/golang.org/x/crypto \
&& cd $GOPATH/src/ \
&& go install golang.org/x/text
```

### 第三方包

如果有未下载的依赖包，进入项目目录，执行以下命令即可

> cd codegen

> go get -v

本项目主要依赖包如下

[Mustache文本模板](https://github.com/cbroglie/mustache)

[配置文件读取viper](https://github.com/spf13/viper)

[数据库ORM框架gorm](https://github.com/jinzhu/gorm)

### 内部依赖

[依赖go_common项目的公共库](http://git.wanpinghui.com/WPH/go_common)

默认的golang依赖库是全局的，如果希望各个项目各自管理自己的依赖，可以使用 [godep](https://github.com/tools/godep)

## 根据数据库生成实体

配置文件在codegen/codegen.toml，详细配置项含义见配置文件中注释

用IDEA或者GoLand之类开发工具打开项目，在内嵌终端Terminal里面跑

> go run codegen/codegen.go

或者打开codegen/codegen.go文件之后，直接点绿色小三角运行
