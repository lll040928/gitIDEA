package main

import "fmt"

func main() {
	var scoremap = make(map[string]int, 8)
	scoremap["沙河娜扎"] = 100
	scoremap["沙河小王子"] = 200
	//fmt.Println(scoremap)
	////判断key是否存在
	//value, ok := scoremap["沙河小王子"]
	//fmt.Println(value, ok)
	//if ok {
	//	fmt.Println("张二狗子在scoremap中", value)
	//} else {
	//	fmt.Println("查无此人")
	//}
	//遍历map,map是无序的，键值对添加的顺序无关
	for _, v := range scoremap {
		fmt.Println(v)
	}
	//只遍历key
	for k := range scoremap {
		fmt.Println(k)
	}
}
