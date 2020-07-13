package main

import "fmt"

func main() {
	x := []int{0,71,29,35,64}
	fmt.Println(x)
	fmt.Println(x[2:])
	fmt.Println(x[:3])
	for i,y:= range x{
		fmt.Println(i,y)
	}
	x = append(x, 89,78,65)
	fmt.Println(x)
	y := []int{56,48}
	x = append(x, y...)
	fmt.Println(x)

	z:= make([]int, 5, 50)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
}
