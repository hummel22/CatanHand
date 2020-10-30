package main

import (
	"fmt"

	"Catan/game"
	"Catan/objects"
)

func main() {
	cards := []objects.Resource{objects.WHEAT,objects.WOOD,objects.ROCK,objects.WHEAT,objects.BRICK, objects.SHEEP, objects.BRICK, objects.WOOD,objects.ROCK,objects.ROCK}
		goods := game.Buy(cards)
	fmt.Printf("Bought [%v] goods worth [%v], \n\t%v", len(goods), game.Score(goods), goods)
}
