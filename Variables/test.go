
					// Summary : 
					//  - Variable declaration: 

					// 	var foo int
					// 	var foo int = 42
					// 	foo := 42 

					// - Can't redeclare variables, but can shadow them 
					// - All variables must be used .
					// - Visibilty: 
					//     lower case first letter for package scope.
					// 	upper case first letter to export .
					// 	no private scope 

package main

import (

	"fmt"
)


// We can declare/initilize variables outside the main function ..


var (

actor = "john mark"
age = 32


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
// short variable declaration assignment :=
z := 25
fmt.Println(z) 
fmt.Printf("Variable value %v and have type of %T \n", z, z)


xx := int64(1)
aa := "wtf"
fmt.Printf("Variable value %v and have type of %T, another one %v \n", xx, xx, aa)
fmt.Printf("Variable value %v and have type of %T, another one %v \n", actor , actor, age)

// You can't declare the variable twice in the same scope .
aa = "no fak"
fmt.Println(aa)

// You will get comple time error if you variable declared and not used in go 


// convert from one variabe type to another with convertion function


var newvariable11 int = 42 
fmt.Printf("%v , %T",newvariable11,newvariable11)

var j float32 = float32(i)
fmt.Printf("%v , %T\n",j,j)
fmt.Printf("%v ,%T, %T\n",i,i,float32(i))


}



