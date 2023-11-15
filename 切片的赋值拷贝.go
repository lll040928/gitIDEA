package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a
	c := make([]int, 5, 5)
	copy(c, a)
	b[0] = 100
	c[1] = 100
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	d := []string{"北京", "上海", "广州", "深圳"}
	d = append(d[0:2], d[3:]...)
	fmt.Println(d)
	//切片的遍历
	//a := []int{1, 2, 3, 4, 5}
	//for i := 0; i < len(a); i++ {
	//	fmt.Println(i, a[i])
	//}
	//fmt.Println()
	//for _, value := range a {
	//	fmt.Println(value)
	//}
	////append()方法，扩容切片
	//var b []int //此时并未申请内存
	//for i := 0; i < 10; i++ {
	//	b = append(b, i)
	//	fmt.Printf("%d len:%d cap:%d ptr:%p\n", b, len(b), cap(b), b)
	//}
}
