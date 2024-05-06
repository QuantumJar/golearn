package main

import "fmt"

//声明了一个新的类型deck，他的属性和[]string (string类型的切片)一致，但是他们属于不同的类型
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

//声明了一个接受者为deck的print方法
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}
