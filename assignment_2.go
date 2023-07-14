package main

import "fmt"

type rectangle struct {
	length  float64
	breadth float64
}

//method 
func (r rectangle) area() float64{
		return r.length*r.breadth
	}

//interface 
var rec = rectangle{length:7, breadth:5}

func main(){
	var s1 = rectangle {7,5}
	fmt.Println(rec.area())
	fmt.Println(s1.area())
}


