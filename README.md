#### 目录机构
```
├─app
│  ├─Common
│  ├─Helpers
│  ├─Http
│  │  ├─Controllers
│  │  ├─Middlewares
│  │  ├─Models
│  │  └─Request
│  ├─Libs
│  └─Repositories
├─bootstrap
│  └─driver
├─config
├─database
├─routes
│  └─RouterGroup
├─service
├─storage
├─tmp
└─vendor
```

####  如时使用指针类型

1. 如果方法需要修改接受者，接受者必须是指针类型。

2. 如果接受者是一个包含了 sync.Mutex 或者类似同步字段的结构体，接受者必须是指针，这样可以避免拷贝。

3. 如果接受者是一个大的结构体或者数组，那么指针类型接受者更有效率。

4. 如果接受者是一个结构体，数组或者 slice，它们中任意一个元素是指针类型而且可能被修改，建议使用指针类型接受者，这样会增加程序的可读性
#### 如时使用值类型（对象类型）
1. 如果接受者是一个 map，func 或者 chan，使用值类型(因为它们本身就是引用类型)。

2. 如果接受者是一个 slice，且方法内的变动，对参数没有影响

3. 如果接受者是一个小的数组或者原生的值类型结构体类型(比如 time.Time 类型)，而且没有可修改的字段和指针，又或者接受者是一个简单地基本类型像是 int 和 string，使用值类型就好了。

4. 切片and数组

#### 工具
```
go get github.com/pilu/fresh
```
#### 生成带图标的exe程序
```
go generate
go build -o younameapp.exe -ldflags="-linkmode internal"
```