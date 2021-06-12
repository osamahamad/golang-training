package main

import (

	"fmt"
)
func main() {
// Variables in go can vary 

var i int 
i = 42
i =28
fmt.Println(i)

/////////////////////
i = 22 
fmt.Println(i)



// Another way to initilize variables


var x int = 45
fmt.Println(x)


// Another way to initilize variables 
// short variable declaration
z := 25
fmt.Println(z)
fmt.Printf("Variable value %v and have type of %T \n", z, z)


xx := int64(1)
aa := "wtf"
fmt.Printf("Variable value %v and have type of %T, another one %v \n", xx, xx, aa)


}
