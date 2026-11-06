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

	fmt.Println(name)
	fmt.Println(rating)
	fmt.Println(hoursPlayed)
	fmt.Println(developer)
	fmt.Println(finished)
	fmt.Println(price)
}