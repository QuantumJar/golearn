# Go Learn
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

