package main

import "fmt"

func main() {
	var i1 = 5
	//the operand "&" when placed in font of a variable,
	//shows us the memory address of that variable
	fmt.Printf("An integer: %d, and its location in memory is: %p\n", i1, &i1)
	//This address can be stored in a special data type called a pointer
	//in this case it is a pointer to an int, here i1: this is denoted by *int.
	//If we call that pointer intP, we can declare it as
	var intP *int
	//Then the following is true:
	intP = &i1
	//intP points to i1
	fmt.Println(intP)
	//if we print,we get same memory address.

}
