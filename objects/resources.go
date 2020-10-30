package objects

import "fmt"

type Resource uint64

const (
	NONE_RESOURCE Resource = iota
	WOOD
	BRICK
	WHEAT
	ROCK
	SHEEP
	MAX_RESOURCE
)

type Resources [MAX_RESOURCE-1]uint64

type Cost Resources


func ToResources(cards []Resource) Resources {
	var resources Resources
	for _, r := range cards {
		if r > NONE_RESOURCE && r < MAX_RESOURCE{
			resources[r-1]++
		}
	}
	return  resources
}

func (r Resources)Buy(cost Cost) (Resources,bool) {
	var remaining Resources
	fmt.Printf("Wallet: %v\n", r)
	fmt.Printf("Price: %v\n", cost)

	for good, price := range cost {
		wallet := r[good]
		if price > wallet {
			return r, false
		}
		remaining[good] = wallet - price
	}
	return remaining, true
}