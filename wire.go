// +build wireinject

package main

import "github.com/google/wire"

/**
说明:这个 build tag 确保在常规编译时忽略 wire.go 文件。
而与之相对的 wire_gen.go 中的 //+build !wireinject。 两个对立的 build tag 是为了确保在任意情况下，两个文件只有一个文件生效， 避免出现 “ContainerByWire() 方法被重新定义” 的编译错误
*/

func InitApp(orderSet) {
	panic(wire.Build(orderSet, Application))
}
