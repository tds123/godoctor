// <<<<< shortassign,10,2,10,13,pass
package main

import "fmt"

func f() (int,int,int) {
	return 5,8,3
}
func main() {
	_,y,_ := f()
	fmt.Println("Value of y :",y)
}
