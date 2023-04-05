package main

import (
	"fmt"
	"math"
)

func swap(x, y int) (int, int) {
	return y, x
}

func Array2() {
	fmt.Println("Enter the number N: ")
	var n int
	var arr []float64
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		arr = append(arr, math.Pow(float64(2), float64(i+1)))
	}
	fmt.Println(arr)
}

func Array14() {
	fmt.Println("Enter the number N: ")
	var n int
	var cin int
	var arr []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cin)
		arr = append(arr, cin)
	}

	for i := 1; i < n; i += 2 {
		fmt.Printf("%v ", arr[i])
	}
	for i := 0; i < n; i += 2 {
		fmt.Printf("%v ", arr[i])
	}

}

func Array21() {
	var n, k, l int
	var arr []float64
	var res, sum, d, cin float64
	fmt.Println("Enter the number N: ")
	fmt.Scan(&n)
	fmt.Println("Enter you array")
	for i := 0; i < n; i++ {
		fmt.Scan(&cin)
		arr = append(arr, cin)
	}
	fmt.Println("Enter the number K: ")
	fmt.Scan(&k)
	fmt.Println("Enter the number L: ")
	fmt.Scan(&l)
	for i := k - 1; i < l; i++ {
		sum += arr[i]
		d++
	}
	res = sum / d
	fmt.Println(res)
}
