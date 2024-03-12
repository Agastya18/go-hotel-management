package main

import (
	"fmt"
	"strings"
)

type Menu struct{
	itemNo    uint
	itemName  string
	itemPrice float64

}

var menu = []Menu{
	{1, "Chole Bhature", 100},
	{2, "Rajma Chawal", 80},
	{3, "Paneer Tikka", 150},
	{4, "Butter Chicken", 200},
	{5, "Chicken Biryani", 180},
	{6, "Mutton Rogan Josh", 250},
	{7, "Fish Curry", 200},
	{8, "Fish Fry", 150},
	{9, "Papad", 10},
	{10, "Raita", 20},
	{11, "Lassi", 30},
	{12, "Cold Drink", 40},
	{13, "Water Bottle", 20},

}

func printMenu(){
	fmt.Printf("%15s\n", "Menu")
	fmt.Printf("%s\n", strings.Repeat("-", 35))
	fmt.Printf("%-7s %6s    %12s\n", "S.No.","Item Name", "Price" )
	fmt.Printf("%s\n", strings.Repeat("-", 35))

	for _, element := range menu {
		fmt.Printf(" %-7d %.2f    %-4s\n", element.itemNo, element.itemPrice, element.itemName)
	}
	fmt.Printf("%s", strings.Repeat("-", 35))
	fmt.Println()
}