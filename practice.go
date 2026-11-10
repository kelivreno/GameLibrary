package main

import "fmt"

func practiceBasics() {

	name := "Minecraft"
	rating := 9
	hoursPlayed := 125
	developer := "Mojang"
	finished := false
	price := 39.99

	hoursPlayed += 25
	price += 5.00
	rating -= 1
	finished = true

	isHighlyRated := rating >= 8

	
	

	fmt.Println(name)
	fmt.Println(hoursPlayed)
	fmt.Println(developer)
	
	if finished {
		fmt.Println("Completed")
	} else {
		fmt.Println("Still playing")
	}
	fmt.Println("^",finished)
	fmt.Println(price)
	fmt.Println(rating)
	if rating >= 8 {
		fmt.Println("Excellent")
	} else if rating >=5 {
		fmt.Println("Average")
	} else {
		fmt.Println("Poor")
	}
	fmt.Println("^",isHighlyRated)

}