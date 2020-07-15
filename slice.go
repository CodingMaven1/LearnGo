package main

import "fmt"

func main() {

	//slices
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
	z = []int{1,5,7,9,8}
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	//mapping

	m := map[string]int{
		"Naman": 20,
		"Ayush": 21,
	}

	fmt.Println(m)

	if value, ok := m["Naman"]; ok {
		fmt.Println("The user exists and his age is", value)
	}

	m["Manas"] = 19

	for l,v :=range m {
		fmt.Println(l,v)
	}

	delete(m, "Manas")

	fmt.Println(m)
}
