package main

import "fmt"

func main() {
	var count int = 10

//for loop
	for i := 0; i < count; i++ {
		fmt.Printf("this is line #%d\n", i)	
	}

//while loop
	var cont bool = true

	count = 0;
	for cont {
		count++
		if count == 10 {
			cont = false
		}
		fmt.Printf("This is iteration %d\n", count)
	}
}