package main

import (
	"fmt"
)

func main() {
	// fmt.Println("hello golang!")
	// fmt.Println("11*11=", 11*11, " 12/12=", 12/12)
	// fmt.Printf("My name is %v, gender is %v", "张三", 18)

	//变量和常量
	const lightspace = 299792 //km/s
	var distance = 56000000   //km
	fmt.Println(distance/lightspace, "seconds")
	distance = 40000000 //km
	fmt.Println(distance/lightspace, "seconds")

	//另一种使用方式
	var (
		name = "张三"
		age  = 18
	)

	var name2, age2 = "李四", 20

	fmt.Println(name, age)
	fmt.Println(name2, age2)

}
