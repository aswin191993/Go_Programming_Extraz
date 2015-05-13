package main

import "fmt"
import "./subname"

func main(){
	val:= subname.Filename()
	fmt.Println("okey")
	fmt.Println(val)

}
