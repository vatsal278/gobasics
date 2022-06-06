package main

import "log"


func main() {
	counter_variable := 1
	
	func increment(x int) {
		for i := 0; i < 10; i++ {
			x = x+1
			log.Print(x)
		}
	}
	for i := 0; i < 5; i++ {
		go increment(counter_variable)
	}
	log.Print(counter_variable)
}
