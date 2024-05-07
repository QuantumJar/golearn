package main

import (
	"fmt"
	"os"
	"strings"
)

// 声明了一个新的类型deck，他的属性和[]string (string类型的切片)一致，但是他们属于不同的类型
type deck []string

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

// 声明了一个接受者为deck的print方法
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

// 发牌 从一个deck中发出handSize张牌，返回为发出的牌形成的deck和剩下牌组成的deck
func deal(d deck, handSize int) (deck, deck) {
	if handSize > len(d) {

	}

	return d[:handSize], d[handSize:]

}

// 把deck类型转换为[]string,然后把这个切片分割为string，使用，作为分隔符返回为一个string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// 保存deck为字符串并以文本的形式保存到本地
func (d deck) saveToFile() error {
	return os.WriteFile("myDeck", []byte(d.toString()), 0666)
}
