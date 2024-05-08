package main

func main() {
	p := Person{
		firstname: "qingjie",
		lastname:  "zhang",
		contactInfo: ContactInfo{
			email:   "690499491@qq.com",
			address: "chengdu"}}

	p.Print()

	p.setFirstname("a")
	// personPrt := p

	// p.setFirstname("sss")
	p.Print()
	// p := &Person{
	// 	firstname: "Alice",
	// 	lastname:  "jim",
	// }

	// // 使用指针修改字段的值
	// p.firstname = "ha"

	// // 打印修改后的值
	// fmt.Println("Name:", p.firstname)
	// fmt.Println("Age:", p.lastname)

}
