package main

import "fmt"

var a int

func main (){
	a= 41
	fmt.Printf("%d\t%b\t%#x\n", a,a,a)
	b := a<<1
	fmt.Printf("%d\t%b\t%#x\n", b,b,b)
}