package main

import "fmt"

func main() {

	for x := 33; x < 123; x++{
		fmt.Printf("%v\t%#U\n",x,x)
	}
}
