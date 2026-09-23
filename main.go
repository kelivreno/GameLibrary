package main

import "fmt"

type Game struct {
	Name        string
	Rating      int
	Finished    bool
	HoursPlayed int
	Developer   string
	Price       float64
}

func main() {
	minecraft := Game{
		Name:        "Minecraft",
		Rating:      9,
		Finished:    false,
		HoursPlayed: 125,
		Developer:   "Mojang",
		Price:       29.99,
	}

	terraria := Game{
		Name:        "Terraria",
		Rating:      8,
		Finished:    true,
		HoursPlayed: 80,
		Developer:   "Re-logic",
		Price:       9.99,
	}

	stardewValley := Game{
		Name:        "Stardew Valley",
		Rating:      10,
		Finished:    false,
		HoursPlayed: 50,
		Developer:   "ConcernedApe",
		Price:       14.99,
	}

	games := []Game{
		minecraft,
		terraria,
		stardewValley,
	}

	fmt.Println(games[0])

}
