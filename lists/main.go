package main

import (
	"container/list"
	"fmt"
)

func main() {
	var strList list.List
	myProfile := struct {
		Id   int
		Name string
	}{1, "Mohamed"}

	strList.PushBack("John Doe")
	strList.PushBack(myProfile)
	strList.PushBack(211)

	for element := strList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
