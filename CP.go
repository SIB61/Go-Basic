package main

import(
	"fmt"
)
func main(){
  var (
	  a int
	  b int 
	  c int 
  )
  fmt.Scan(&a,&b,&c)
  if a-c >= (20-b)*36 {
	  fmt.Println("NO")
  }else{
	  fmt.Println("YES")
  }

}