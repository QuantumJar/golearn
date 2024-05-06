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

}
