package main

import "fmt"

type Game struct {
	Name string
	Rating int
	Finished bool
	HoursPlayed int
	Developer string
	Price float64
}

func main() {
	minecraft := Game {
		Name: "Minecraft",
		Rating: 9,
		Finished: false,
		HoursPlayed: 125,
		Developer: "Mojang",
		Price: 29.99,

	}

	fmt.Println(minecraft.Rating)
	fmt.Println(minecraft.Finished)

	minecraft.Rating = 10
	minecraft.Finished = true

	fmt.Println(minecraft.Rating)
	fmt.Println(minecraft.Finished)




	// fmt.Println(gameName)
	// fmt.Println(rating)
	// fmt.Println(finished)
	// fmt.Println(hoursPlayed)
	// fmt.Println(developer)
	// fmt.Println(price)
	
}