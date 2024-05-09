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

	printMap(map2)

}

//遍历map
func printMap(m map[string]int) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}
