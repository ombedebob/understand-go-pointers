package main

import "fmt"

func main() {
	msg := "Hello Gopher!"
	var sent_msg *string = &msg
	*sent_msg = "Are you getting pointers yet??"

	fmt.Printf("Here is the pointer sent_msg: %p\n", sent_msg)
	fmt.Printf("Here is the string sent_msg: %s\n", *sent_msg)
	fmt.Printf("Here is the string msg: %s\n", msg)

	//by changing the value of *sent_msg,msg also changed
}
