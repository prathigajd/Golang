package main

import "fmt"

type employee struct{
	name string
	age int
	Experience int

}

func main() {

	// Arithmetic operators 
	a := 2
    b := 3
	fmt.Println("addition of two number a+b: ", a+b)
	fmt.Println("subraction of two number a-b: ", a-b)
	fmt.Println("Multiplication of two number a*b: ", a*b)
	fmt.Println("division of two number a/b: ", a/b)
	fmt.Println("Remainder of two number a%b: ", a%b)

	//logical operators
     
	var c int = 4 
	var d int = 10 

	fmt.Println(" And operator : ", c&d)
	fmt.Println(" Or operator : ", (c>d) || (c==d))
	fmt.Println(" not operator : ", !(c!=d))

	//relational operators 

	if a>b {
	fmt.Println ("a is greater:", a)
	} else {
		fmt.Println ("b is greater:", b)
	}
	if a<b {
		fmt.Println ("a is lesser:", a)
	} else {
		fmt.Println ("b is lesser:", b)
	}
	if a==b {
		fmt.Println ("a is equals to b", a==b)	
	} else {
		fmt.Println ("a is not equals to b", a!=b)
	}

	//for loop with slice 
	var days = []string{"Mon", "Tue,", "wed," ,"thurs", "fri", "sat", "sun"}
	for i, j := range days {
         fmt.Println("index for", i, "value is", j)
	}
    
   //structure 
   
   var e = employee{name: "Prathiga", age: 23, Experience: 3}
   fmt.Println(e)
} 
