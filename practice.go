package main

import "fmt"

func main() {

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

	if finished {
		fmt.Println("Completed")
	} else {
		fmt.Println("Still playing")
	}
	

	fmt.Println(name)
	fmt.Println(rating)
	fmt.Println(hoursPlayed)
	fmt.Println(developer)
	fmt.Println(finished)
	fmt.Println(price)
	fmt.Println(isHighlyRated)
}