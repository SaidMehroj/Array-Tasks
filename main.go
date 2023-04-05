package main

import "fmt"

func main() {
	fmt.Println("Выберите действие")
	fmt.Println("1 - Массивы, задача 2")
	fmt.Println("2 - Массивы, задача 14")
	fmt.Println("3 - Массивы, задача 21")
	var a int
	fmt.Scan(&a)
	switch a {
	case 1:
		Array2()
	case 2:
		Array14()
	case 3:
		Array21()
	}
}
