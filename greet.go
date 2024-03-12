package main

import "fmt"

func greet(customerName string) {
	fmt.Printf("%52s %s\n", "Namaste ", customerName)
	fmt.Printf("%72s\n", "_/\\_ Welcome to Haridwar Haat Bhojanalya! _/\\_ ")
	fmt.Println()
}
