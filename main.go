package main

import "fmt"

type User interface{
	toStoring(s string)
}

type Hoge struct{
	user string
}

func (h Hoge)toStoring(s string){
	fmt.Println(s)
}

func main() {
	var hoge Hoge = Hoge{user: "man"}
	hoge.toStoring("test")

	var user User 
	user = hoge
	user.toStoring("test2")
}