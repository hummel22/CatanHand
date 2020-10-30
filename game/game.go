package game

import "Catan/objects"

func Buy(cards []objects.Resource) []objects.Good {
	resources := objects.ToResources(cards)

	return Graph(resources)


}
