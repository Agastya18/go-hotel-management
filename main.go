package main
import "fmt"



func main(){

	var customerName string
	fmt.Println("What is your first name?")
	fmt.Scan(&customerName)

	greet(customerName)
	orderItems()
}