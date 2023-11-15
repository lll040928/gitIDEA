package main

import "fmt"

func main() {
	var a map[string]int
	fmt.Println(a == nil)
	a = make(map[string]int, 8)
	fmt.Println(a == nil)
	//map添加键值对
	a["沙河娜扎"] = 100
	a["沙河小王子"] = 200

	fmt.Println(a)
	fmt.Printf("a=%#v\n", a)
	fmt.Printf("%v\n", a)
	fmt.Printf("type:%T\n", a)
	//声明map的同时完成初始化
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("b:%#v\n", b)
	fmt.Printf("type:%T\n", b)

	var c map[int]int
	c[100] = 200
	fmt.Println(c)
}
