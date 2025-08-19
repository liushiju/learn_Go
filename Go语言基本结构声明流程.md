# 常见的Go工程的基本结构声明流程

1、包引用声明

2、标准库或第三方开源库引用

3、常量声明

4、变量声明

5、结构体及其所属方法声明

6、默认初始化方法声明

7、主函数声明

8、普通函数声明

## 对应代码

```go
// 包引用
package main

// 库引用
import "fmt"

// 常量声明
const s1 = "脑子进煎鱼了"

// 变量声明
var s2 string = "煎鱼"

// 结构体及其所属方法声明
type T struct {}
func (t T) Func1() {}

// 初始化方法声明
func init() {}

// 主函数声明
func main() {
fmt.Println(s2)
}

// 普通函数声明
func Func2() {}
```

## 注意

在Go语言中存在<font color=red>"导出和非导出"</font>的概念，类比PHP中的public和private。

Go语言中是以<font color=red>大小写</font>作为标识区分，如：大写字母开头，则为可到出(对外公开)，若为小写字母，则为非导出(对外不公开)属性

```go
// 可导出，允许外部包引用
type T struct {}

// 非导出，仅允许当前包引用
type t struct {}
```

## 基本程序结构

### 函数调用

#### 1、普通函数

定义格式：

```go
func [函数名](入参)(出参)
```

例：

```sh
func callFuncA(x, y string) (s string, err error) {
return x + y, nil
}

func main() {
callFuncA("炸", "煎鱼")
}
```

说明：

示例中，声明了一个函数名为callFuncA的方法。从调用场景上讲，它只允许在包内调用，因此首字母小写。

其具有两个入参，x，y。类型为string

两个出参变量，s，err。类型为string，error

另外在函数体内返回值时，可以采用快捷方式

```go
func callFuncA(x,y string)(s string, err error){
    s = x + y
    return
}
```

在出参时声明的变量名，是可以应用到自身函数的。因此若直接return则会隐式返回已经声明的出参变量。

而在函数定义时，其入参还支持可变参数的语法：

```go
func callFuncA(x ... string)(s string, err error){
    s = strings.Join(x,",")
    return
}

func main(){
    fmt.Println(callFuncA("炸","煎鱼"))
}
```

说明：

在入参变量上声明为`x ... string`，则表示变量x是string类型的可变变量，能够在入参时传入多个string参数。

可变变量所传入的格式为切片(slice)类型

```go
[0:炸 1:煎鱼]
```

<font color=red>一般对可变变量的常见后续操作是循环遍历处理，或是进行拼接操作</font>

#### 2、匿名函数

Go语言也默认支持匿名函数的声明，声明方式几乎与普通函数几乎一样

```go
func main(){
    s := func(x,y string)(s string, err error){
        return x + y, nil
    }
    s("炸","煎鱼")
}
```

<font color=red>匿名函数可以在任意地方声明，且不需要定义函数名，如果在函数体后马上跟 () 则表明后立即执行</font>

```go
func main(){
    s,\_ := func(x,y string)(s string,err error){
        return x+y, nil
    }("炸","煎鱼")
}
```

而在所有函数类使用中，有一点非常重要，那就是函数变量的作用域

```go
func main(){
    x,y := "炸","煎鱼"
    s,\_ := func()(s string,err error){
        return x+y, nil
    }()
    fmt.Println(s)
}
```

该匿名函数没有形参，函数内部没有定义相应的变量，此时其读取的是全局的x,y变量的值，输出结果是"炸煎鱼"

```go
func main(){
    x,y := "炸","煎鱼"
    \_,\_ = func(x,y string)(s string,err error){
        x = "吃"
        return x+y,nil
    }(x,y)
    fmt.Println(x,y)
}
```

该匿名函数有形参，但是函数内部又重新赋了值x。那么最终外部所输出的变量 x 的值是什么呢？输出结果是 “炸 煎鱼”。

为什么明明在函数内已经对变量 x 重新赋值，却依然没有改变全局变量 x 的值呢？

其本质原因是作用域不同，函数内部所修改的变量 x 是函数内的局部变量。而外部的是全局的变量，所归属的作用域不同。

#### 3、结构方法

在结合结构体(struct)的方式下，可以声明归属于该结构体下的方法

```go
type T struct{}

func NewT() \*T {
return &T{}
}

func (t \*T) callFuncA(x, y string) (s string, err error) {
return x + y, nil
}

func main() {
NewT().callFuncA("炸", "煎鱼")
}
```

具体的函数的使用方法与普通函数一样，无其他区别。

#### 4、内置函数

Go语言本身有支持一些内置函数，这些内置函数的调用不需要引用第三方标准库。内置函数的作用是用于配合 Go 语言的常规使用，数量非常少，如下：

- 用于获取某些类型的长度和容量：len、cap。

- 用于创建并分配某些类型的内存：new、make。
- 用于错误处理机制（异常恐慌、异常捕获）：panic、recover。
- 用于复制和新增切片（slice）：copy、append。
- 用于简单输出信息：print、println。
- 用于处理复数：complex、real、imag。

### Go语言高级特性

#### 1、复合数据类型

- **map**

Map(字典)的标准格式：

```go
map[KeyType]ValueType
```

对应的KeyType是键类型，ValueType是值类型

声明方式：

1、可利用var关键字声明

```go
var m map[string]int
```

2、借助内置函数make声明

```go
m = make(map[string]int)
```

两者区别：

1）直接用var声明，只是声明了变量m，但没有初始化内存，如果直接写入会导致出现nil map继而导致恐慌(panic)事

2）而使用内置函数make，会声明并初始化一个哈希映射数据结构，并返回一个指向它的映射值。这是正确的初始化map的方法

在 map 的使用上，可以直接填写某一个key，设置它的值：

```go
m["jianyu"]=66
```

写入后，如果在程序中需要读取 map 里的值，可以直接根据 Key 读取：

```go
// 直接读取
i := m["jianyu"]

// 第二个参数可以判断是否存在，false 表示不存在
i, ok := m["jianyu"]
```

如果是对 map 有删除的诉求，可以借助内置函数 delete，指定某一个 Key：

```go
delete(m, "jianyu")
```

关于 map 有一点是需求特别注意的，原生的 go map 并不支持并发写入。若一个不小心并发写入了，就会导致抛出系统系别的致命错误：

```go
fatal error: concurrent map read and map write
```

这个报错会导致进程退出，这种时候建议调整写法、加上读写锁，或是使用标准库 sync.Mutex 来解决。

- **array**

有一个潜伏在深处，但明处相对少用的数据类型，那就是数组（array），声明的格式为：

```go
var a [10]int
```

代表着变量 a 是类型为 int、长度为 10 的数组，数组的长度是其类型的一部分，在声明后不可变，是受限制的。

在短声明上，我们还可以这么去使用：

```go
primes := [6]int{2, 3, 5, 7, 11, 13}
```

读取时，可以直接通过索引访问数组：

```go
fmt.Println(primes[0])
```

就会读取到变量 primes 的第 1 个参数（索引从 0 开始）。