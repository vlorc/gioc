
# [Gioc](https://github.com/vlorc/gioc)

[English](https://github.com/vlorc/gioc/blob/master/README.md)

[![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![codebeat badge](https://codebeat.co/badges/c41b426c-4121-4dc8-99c2-f1b60574be64)](https://codebeat.co/projects/github-com-vlorc-gioc-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/vlorc/gioc)](https://goreportcard.com/report/github.com/vlorc/gioc)
[![GoDoc](https://godoc.org/github.com/vlorc/gioc?status.svg)](https://godoc.org/github.com/vlorc/gioc)
[![Build Status](https://travis-ci.org/vlorc/gioc.svg?branch=dev)](https://travis-ci.org/vlorc/gioc?branch=dev)
[![Coverage Status](https://coveralls.io/repos/github/vlorc/gioc/badge.svg?branch=dev)](https://coveralls.io/github/vlorc/gioc?branch=dev)

gioc是一个轻量级的Ioc框架，它提供注册表和工厂、依赖解决方案

## 特性

* 依赖解析
* 依赖注入
* 单例、瞬态
* 自定义tag
* 调用器
* [惰性加载](https://github.com/vlorc/gioc/blob/master/examples/lazy/main.go)
* [结构体扩展](https://github.com/vlorc/gioc/blob/master/examples/depend/main.go)
* [模块](https://github.com/vlorc/gioc/blob/master/examples/module/main.go)

## 安装
	go get github.com/vlorc/gioc

## 快速开始

* 创建根容器
```golang
container := gioc.NewRootContainer()
```

* 注册实例
```golang
err := container.AsRegister().RegisterInstance(1,"age")
```

* 获取实例
```golang
instance,err := container.AsProvider().Resolve((*int)(nil), "age"))
```

## 例子

* 基本工厂
```golang
import (
    "fmt"
    "github.com/vlorc/gioc"
    "github.com/vlorc/gioc/factory"
    "github.com/vlorc/gioc/types"
)

func main() {
    container := gioc.NewRootContainer()
    age := 17

    // register an int type value factory,this is similar to RegisterInstance
    container.AsRegister().RegisterFactory(factory.NewValueFactory(age),(*int)(nil),"age")
    // create a custom func factory
    inc := factory.NewFuncFactory(func(types.Provider) (interface{}, error) {
        age++
        return age, nil
    })

    // register an int type
    container.AsRegister().RegisterFactory(inc,&age,"inc")
    // convert custom factory into singleton mode factory
    container.AsRegister().RegisterFactory(factory.NewSingleFactory(inc),&age,"once")
    // get an instance type int and name age
    fmt.Println(container.Resolve((*int)(nil), "age"))
    // same as above,this value add 1 every times
    fmt.Println(container.Resolve((*int)(nil), "inc"))
    // same as above,but only once
    fmt.Println(container.Resolve((*int)(nil), "once"))
}
```

* 基本模块
```golang
import (
    "fmt"
    . "github.com/vlorc/gioc"
    . "github.com/vlorc/gioc/module/operation"
)

func main() {
    NewRootModule(
        Import(),//import module
        Declare(
            Instance(1), Id("id"),//declare instance
        ),
        Bootstrap(func(param struct{ id int64 }) {
            fmt.Println("id:", param.id)
        }),
    )
}
```

## 许可证

这个项目是在Apache许可证下。查看完整的许可证文本的许可证文件。

## 接口

+ Provider(提供商)
	+ 提供工厂发现
	+ 提供实例填充
+ Factory(工厂)
	+ 负责生成实例
	+ 基本工厂有价值工厂，方法工厂，代理工厂，单例工厂，类型工厂
+ Mapper(映射器)
	+ 通过ID获取工厂
+ Binder(绑定器)
	+ 通过ID绑定工厂
	+ 可以转换为只读映射器
+ Register(注册器)
	+ 作为工厂和选择器的连接
	+ 提供类型、实例、方法工厂转换
	+ 提供绑定器、映射器、自定义工厂的注册
+ Dependency(依赖)
	+ 是目标类型依赖性分析结果的集合
	+ 通过实例转换为注射器
+ Injector(注射器)
	+ 根据依赖填充实例
+ Builder(构造器)
	+ 也是一个工厂
	+ 使用Factory来获取实例和注入器来解决依赖关系
+ Container(容器)
	+ 提供Register和Provider，并且父容器组成遍历
	+ 转换为只读提供程序
	+ 转换为密封容器
+ Selector(选择器)
	+ 使用类型和id索引工厂
	+ 自动创建绑定器和映射器
	+ 索引模式隔离
+ Module(模板)
    + 导入模块
    + 导出工厂
    + 声明工厂


# 路线图
有关计划特性和未来方向的详细信息请参考[路线图](https://github.com/vlorc/gioc/blob/master/ROADMAP.md)

# 关键字

**依赖注入，控制反转**

# 参考
