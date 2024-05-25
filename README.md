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

