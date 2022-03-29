package main

import (
	"container/list"
	"fmt"
)

func main(){
   List := list.New()
   List.PushBack(10)
   List.PushBack("sabit")
   fmt.Println(List.Back().Value)
   fmt.Println(List.Front().Value)

   List.PushBack("gopal")
   fmt.Println(List.Back().Value)
   fmt.Println(List.Front().Value)
   

}