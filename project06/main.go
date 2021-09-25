package main

import "fmt"

func main() {
	//	空のinterfaceはC言語でいうvoid*型
	var a interface{}
	var i int = 5
	s := "Hello world"
	a = i 
	fmt.Println(a)
	a = s
	fmt.Println(s)
}