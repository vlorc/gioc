# [Gioc](https://github.com/vlorc/gioc)

[English](https://github.com/vlorc/gioc/blob/master/README.md)

[![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![codebeat badge](https://codebeat.co/badges/c41b426c-4121-4dc8-99c2-f1b60574be64)](https://codebeat.co/projects/github-com-vlorc-gioc-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/vlorc/gioc)](https://goreportcard.com/report/github.com/vlorc/gioc)
[![GoDoc](https://godoc.org/github.com/vlorc/gioc?status.svg)](https://godoc.org/github.com/vlorc/gioc)
[![Build Status](https://travis-ci.org/vlorc/gioc.svg?branch=master)](https://travis-ci.org/vlorc/gioc?branch=master)
[![Coverage Status](https://coveralls.io/repos/github/vlorc/gioc/badge.svg?branch=master)](https://coveralls.io/github/vlorc/gioc?branch=master)

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

	go get -u github.com/vlorc/gioc

## 快速开始

* 创建根模块

* Create Root Module

```golang
gioc.NewRootModule()
```

* Import Module

```golang
NewRootModule(
    Import(
        ConfigModule,
        ServerModule,
    )
)
```

* Declare Instance

```golang
NewRootModule(
    Declare(
        Instance(1), Id("id"),
        Instance("ioc"), Id("name"),
    ),
)
```

* Export Instance

```golang
NewModuleFactory(
    Export(
        Instance(1), Id("id"),
        Instance("ioc"), Id("name"),
    ),
)
```

## Examples

* Basic Module

```golang
import (
    ."github.com/vlorc/gioc"
    ."github.com/vlorc/gioc/module"
    ."github.com/vlorc/gioc/module/operation"
)

// config.go
var ConfigModule = NewModuleFactory(
    Export(
        Mapping(map[string]interface{}{
            "id": 1,
            "name": "ioc",
        }),
    ),
)

// main.go
func main() {
    NewRootModule(
        Import(ConfigModule),
        Bootstrap(func(param struct{ id int; name string }) {
            println("id: ", param.id, " name: ",param.name)
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
+ Register(注册器)
    + 作为工厂和选择器的连接
    + 提供类型、实例、方法工厂转换
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
    + 通过类型和名称寻找工厂
+ Module(模板)
    + 导入模块
    + 导出工厂
    + 声明工厂

# 路线图

有关计划特性和未来方向的详细信息请参考[路线图](https://github.com/vlorc/gioc/blob/master/ROADMAP.md)

# 关键字

**依赖注入，控制反转**

# 参考
