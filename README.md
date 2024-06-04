# Go Learn

## vscode使用快捷键

- 打开终端 ctrl+shift+`
- 切换打开的标签 crtl+tab
- 关闭当前标签crtl+w

## go 命令行指令
- go build 用于编译一个或者多个.go文件
- go run 用于编译并运行一个或者多个.go文件
- go fmt 格式化当前文件夹下的所有.go文件
- go install 编译并且安装一个包
- go get 下载依赖的包源文件
- go test 运行当前项目的测试用例

## 声明变量
``` go
//正常的声明一个变量
var card string = "this is a string"
//短声明
card := "this is a string"
```

## slice和for循环

```go
package main

import "fmt"

func main() {
	//声明一个string类型的slice
	cards := []string{"Ace of Diamonds"}

	//append func append(slice []Type, elems ...Type) []Type
	//append方法接收一个任意类型的切片和此类型的元素，并返回这个类型的切片
	cards = append(cards, "1 of Spades")
	cards = append(cards, "King of Squares")

	//遍历一个切片
	for i, v := range cards {
		fmt.Println(i, v)
	}
   /**
   0 Ace of Diamonds
    1 1 of Spades
    2 King of Squares
   */

}
```

## 面向对象和go

面向对象语言中，有class的概念，但是go不是OOP语言，那么go如何去表示这种概念呢？

```java
public class Deck {
    
   //属性
    private List<String> cards;
    
    //方法
    public void print(){
        for(int i= 0; i<cards.length();i++){
            //do something
        }
	}
}
```

在go中，因为没有class的概念，所以需要使用type来定义新的类型

```go
//deck.go中
package main

import "fmt"

//声明了一个新的类型deck，他的属性和[]string (string类型的切片)一致，但是他们属于不同的类型
type deck []string

//声明了一个deck类型的创建方法，（个人觉得这个很像构造器）
func newDeck() deck {
	cards := deck{}
	cardsSuits := []string{"Heart", "Club", "Spade", "Diamond"}
	cardsValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "J", "Q", "k"}

	for i := 0; i < len(cardsSuits); i++ {
		for j := 0; j < len(cardsValues); j++ {
			card := cardsValues[j] + " of " + cardsSuits[i]
			cards = append(cards, card)
		}
	}
	return cards
}
//声明了一个接受者为deck的print方法
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

```



## 基本文件操作

```go
// WriteFile writes data to the named file, creating it if necessary.
// If the file does not exist, WriteFile creates it with permissions perm (before umask);
// otherwise WriteFile truncates it before writing, without changing permissions.
// Since WriteFile requires multiple system calls to complete, a failure mid-operation
// can leave the file in a partially written state.
func WriteFile(name string, data []byte, perm FileMode) error {
	f, err := OpenFile(name, O_WRONLY|O_CREATE|O_TRUNC, perm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}
```

## 从文件创建一个deck



```go
func newDeckFromFile(filename string) deck {

	bytes, err := os.ReadFile(filename)
	//这里如果没错误， err== nil
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	//[]byte 转 []string
	s := string(bytes)

	var stringList []string = strings.Split(s, ",")

	return deck(stringList)
}


```

## go的交换的写法

```go
//我写的方法
func (d deck) shuffle() {
	for i := range d {
		randNumber := rand.Intn(len(d) - 1)
		d.swap(i, randNumber)
	}
}
func (d deck) swap(i, j int) {
	temp := d[i]
	d[i] = d[j]
	d[j] = temp
}
```

```go
///视频的写法
func (d deck) shuffle() {
	for i := range d {
		randNumber := rand.Intn(len(d) - 1)
        //这里不同
		d[i],d[randNumber] = d[randNumber],d[i]
	}
}
```

## 随机数

```go
func main() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	i := r.Intn(100)
	for {
		fmt.Print(i)
	}
}
```

## Test with go

```go
package main

import "testing"

//这里的*testing.T 是test handler
func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 52 {
		t.Errorf("Expect 52 cards but got %v!", len(deck))
	} else if deck[0] != "Ace of Heart" {
		t.Errorf("first card should be a Ace of Heart if not shuffled!")
	}
}


```

## 传值和传指针

- &variable :表示取得这个变量所在的地址
- *pointer :给我这个地址所指向的东西
- *Type: 表示一个指向某种类型的指针

在 Go 语言中，结构体的指针可以通过`.`操作符来修改字段的值。这是因为在 Go 语言中，当你使用结构体指针访问字段时，编译器会自动解引用指针并访问字段。这使得修改结构体字段的值变得非常方便。

```go
package main

import "fmt"

type Person struct {
	firstname   string
	lastname    string
	contactInfo ContactInfo
}

type ContactInfo struct {
	email   string
	address string
}

func (p Person) Print() {
	fmt.Printf("%+v", p)
}

//1.这里接受者参数的*Person 表示这里需要的是一个Person类型的指针
func (personPtr *Person) setFirstname(firstname string) {
	// (*personPtr).firstname = firstname 这种写法和下面的写法有什么不一样呢？
	
	personPtr.firstname = firstname
}

```

## 指针陷阱（go传参的注意点）

基础类型传入的是值，struct是传值。引用类型传入的是指针

![1715182149124](C:\Users\69049\AppData\Roaming\Typora\typora-user-images\1715182149124.png)

## map in go

声明一个map，删除一个key

```go
package main

import "fmt"

func main() {
	//声明map
	//1.
	map1 := make(map[string]int)
	map1["name"] = 1
	fmt.Println(map1)
	delete(map1, "name")
	fmt.Println(map1)
	//2.
	map2 := map[string]int{"name": 1}
	fmt.Println(map2)
}

//遍历map
func printMap(m map[string]int) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

```

## interface 重要

概述：

interface类型表示了对其他类型行为的概括或者抽象。这里的概括指的是interface可以让我们写出更加灵活更加有适应性的函数，**因为这种函数的实现可以有多种不同的实现，而不是只能有一种具体的实现**。

很多面向对象语言中都有interface的概念，go也有接口的概念，但是什么让go语言的接口和其他语言截然不同的的？——go语言的接口的实现是**隐式的（implicit）**。

换言之，我们不用显示的指明某一个类型必须要**实现**（这里使用了实现这个措辞是因为在java中，其他的类需要使用implement来实现某个接口定义的方法）某一个接口。其他类型仅仅需要实现一些接口中的方法就行了。

这种设计



## 接口是一种契约 contract

目前我们接触到的类型全部是具体的类型。一个具体的类型必须精确的指明。




## RPC in go



## 多线程

### go routine 

Q: 什么是go routine？

A：是一个独立运行的函数，通过go关键字开启。它有自己独立的调用栈。开销很小，开启几千个甚至十万个go routine都是非常实用的。 他不是一个线程。可能一个go程序只有一个线程，但是却有几千个go routine。

在引入多线程之前，首先看一个例子

```go
//这个例子中,控制台打印了两行输出语句，而且他们是顺序的执行。
func main() {
	someFunc("somefunc")
	fmt.Println("print from main routine")
}

func someFunc(str string) {
	fmt.Println(str)
}
//控制台打印信息：
//[Running] go run "g:\Code\go_learn\golearn\concurrency\main.go"
//somefunc
//print from main routine
```



然后在看看另一个例子,这个例子在函数someFunc调用前加入了一个关键字go，go 这个关键字，在go语言中表示开启一个新的go routine。

Q: 结合自己对线程的了解，你觉得控制台应该打印怎么样结果呢？

```go
//go 这个关键字，在go语言中表示开启一个新的线程，这个线程在go中被称为一个go routine
func main() {
	go someFunc("somefunc")
	fmt.Println("print from main routine")
}

func someFunc(str string) {
	fmt.Println(str)
}

```

我们来看看控制台多次打印的结果

```go
[Running] go run "g:\Code\go_learn\golearn\concurrency\main.go"
print from main routine
somefunc

[Done] exited with code=0 in 0.585 seconds

[Running] go run "g:\Code\go_learn\golearn\concurrency\main.go"
somefunc
print from main routine

[Done] exited with code=0 in 0.567 seconds

[Running] go run "g:\Code\go_learn\golearn\concurrency\main.go"
print from main routine
somefunc

[Done] exited with code=0 in 0.581 seconds

[Running] go run "g:\Code\go_learn\golearn\concurrency\main.go"
print from main routine


[Done] exited with code=0 in 0.589 seconds
```

请注意：

- 在第一次结果中，首先打印的是主线程的语句，然后打印了someFunc线程
- 在第二次首先打印了someFunc,然后是主线程
- 第三次又和第一次的结果相同
- 第四次只有主线程的结果，完全没有看到someFunc线程的输出语句。

我们可以总结出当有多线程在当前线程被启动后，这些被启动的线程的执行顺序不再是顺序的。为什么会出现以上的情况呢？这个答案目前还不能完整的回答



### Channel

Channel是一种go语言用于线程间通信的工具，对于很多语言来说，线程间的通信是由共享内存来实现的。共享内存实现的通信需要必要的同步手段来实现数据的线程安全。

下面演示一个go使用channel来进行线程见通信的例子

```go

func main() {
	//创建一个string类型的channel
	strChannel := make(chan string)

	//声明一个匿名函数并且在另一个线程中运行他
	go func() {
		//数据写入channel使用了箭头函数
		strChannel <- "write something to the string channel"
	}()

	//在主线程中阻塞的读取strChannel中的数据
	data := <-strChannel

	fmt.Println(data)
}

//控制台输出 
//[Running] go run "g:\Code\go_learn\golearn\concurrency\channel\main.go"
//write something to the string channel
```



以上的程序实现了一个简单的多线程使用channel来通信的例子。这里的channel作为一个主线程和分线程通信的容器，把分线程的数据写入了channel，然后从主线程中获取该数据。可以看到控制台理科打印了输出结果。

这里做一个简单的修改

```go
func main() {
	//创建一个string类型的channel
	strChannel := make(chan string)

	//声明一个匿名函数并且在另一个线程中运行他
	go func() {
		//数据写入channel使用了箭头函数
		strChannel <- "write something to the string channel"
		fmt.Println("我已经发送了数据，我可以退出了吗")
	}()

	//在主线程中读取strChannel中的数据，
	time.Sleep(time.Second * 5)
	data := <-strChannel

	fmt.Println(data)
}
//控制台在5秒后打印了如下信息:
[Running] go run "g:\Code\go_learn\golearn\concurrency\channel\main.go"
我已经发送了数据，我可以退出了吗
write something to the string channel

[Done] exited with code=0 in 5.878 seconds
```

这是因为channel是阻塞的。当我们在主线程获取这个channel的数据的时候，主线程会一直处于阻塞状态。

### Select 语句

select选择器是用于对监听多个channel上发生的事件的一种工具。下面给定一个例子，我们需要对部署在设备上的温度传感器进行监控，如果温度超过了阈值这个传感器就向自己的channel发送一个消息。在主线程中，我们要捕获这个消息从而做出对应的处理。

```go
package main

import "fmt"

//定义了一个警告接口
type Alert interface {
	Alert() string
}

//定义了一个CPU结构体模拟传感器
type CPU struct {
}

//实现方法
func (c CPU) Alert() string {
	return "CPU 温度过高"
}

//定义了一个风扇结构体模拟传感器
type Fan struct {
}

//实现方法
func (f Fan) Alert() string {
	return "风扇 温度过高"
}

func main() {

	fanChannel := make(chan string)
	CPUChannel := make(chan string)

	fan := Fan{}
	cpu := CPU{}

	go func() {
		//CPU发出了温度过高的告警
		CPUChannel <- cpu.Alert()
	}()

	go func() {
		//风扇发出了温度过高的告警
		fanChannel <- fan.Alert()
	}()

	select {
	case fanAlert := <-fanChannel:
		fmt.Printf("%v,请运维人员进行处理。", fanAlert)
	case cpuAlert := <-CPUChannel:
		fmt.Printf("%v,请运维人员进行处理。", cpuAlert)
	}
}

//控制台在多次运行后的输出结果如下
风扇 温度过高,请运维人员进行处理。
CPU 温度过高,请运维人员进行处理。
```

以上是一个select语句对多个channel上发生的事件进行监听的演示。本质上这是一个事件驱动的编程模型。当某一个channel中发生了事件的时候，会有相对应的方法去处理它。但是这里有一处并不满足这个事件驱动编程模型的点就是，这个select语句只会执行一次，但是我们的CPU和风扇同时发生了告警。我们应该如何去处理这个缺陷呢？



## go并发模型

下面介绍一些并发的编程模型，这些编程模型并不是语言层面的，但是在语言层面他们的实现会有难易之分。

### 1. for - select-loop

对于我们讲的第一个select案例，实际上仅仅体现了select选择器的基本使用方式，甚至在实际的代码中，这种只使用一次的方式是错误的，违背事件驱动编程模型的。因为我们使用select选择器时，本应该持续的关注某一些通道上发生的事件。

下面我们对之前的代码进行修改。

```go
package main

import "fmt"

//定义了一个警告接口
type Alert interface {
	Alert() string
}

//定义了一个CPU结构体模拟传感器
type CPU struct {
}

//实现方法
func (c CPU) Alert() string {
	return "CPU 温度过高"
}

//定义了一个风扇结构体模拟传感器
type Fan struct {
}

//实现方法
func (f Fan) Alert() string {
	return "风扇 温度过高"
}

func main() {

	fanChannel := make(chan string)
	CPUChannel := make(chan string)

	fan := Fan{}
	cpu := CPU{}

	go func() {
		//CPU发出了温度过高的告警，使用了while循环
		for {

			CPUChannel <- cpu.Alert()
		}
	}()

	go func() {
		//风扇发出了温度过高的告警，使用了while循环
		for {

			fanChannel <- fan.Alert()
		}
	}()
    
	for {

		select {
		case fanAlert := <-fanChannel:
			fmt.Printf("%v,请运维人员进行处理。", fanAlert)
		case cpuAlert := <-CPUChannel:
			fmt.Printf("%v,请运维人员进行处理。", cpuAlert)
		}
	}

}

//这里控制台会打印每一个接收到的channel事件，并且不会自动的停止。
```



### 2. done channel

Done Channel里面的done是一种发送给这个channel的一个信号用来控制（取消）输入channel的处理。一旦从done channel中读取到一个信号或者done channel被管理，输入channel的处理就被取消

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)

	go boring(done)

	//主线程3秒后给done channel 发信号
	time.Sleep(time.Second * 3)

	done <- false

	time.Sleep(time.Hour * 3)

}

func boring(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("收到done信号")
			return
		default:
			fmt.Println("working")
		}
	}
}

```



### 3. pipeline



### 4.generator

