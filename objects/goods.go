package objects

type Good uint64

const (
	NONE_GOOD Good = iota
	CITY
	TOWN
	ROAD
	CARD
	MAX_GOOD
)

const (
	CITY_VALUE uint64 = 4
	ROAD_VALUE uint64 = 2
	CARD_VALUE uint64 = 3
	TOWN_VALUE uint64= 4
)

var Values map[Good]uint64 = map[Good]uint64{
	CITY: CITY_VALUE,
	ROAD: ROAD_VALUE,
	CARD: CARD_VALUE,
	TOWN: TOWN_VALUE,
}


var (
	CITY_COST map[Resource]uint64 = map[Resource]uint64{
		ROCK: 3,
		WHEAT: 2,
	}
	TOWN_COST map[Resource]uint64 = map[Resource]uint64{
		BRICK: 1,
		WHEAT: 1,
		SHEEP: 1,
		WOOD: 1,
	}
	ROAD_COST map[Resource]uint64 = map[Resource]uint64{
		BRICK: 1,
		WOOD: 1,
	}
	CARD_COST map[Resource]uint64 = map[Resource]uint64{
		ROCK: 1,
		WHEAT: 1,
		SHEEP: 1,
	}

	Costs map[Good]map[Resource]uint64 = map[Good]map[Resource]uint64{
		CITY: CITY_COST,
		ROAD: ROAD_COST,
		TOWN: TOWN_COST,
		CARD: CARD_COST,
	}

	CostsKey map[Good]Cost = make(map[Good]Cost, len(Costs))
)

func init() {
	// Hard Code to bitmap
	for good, costDef := range Costs {
		CostsKey[good] = toCost(costDef)
	}
}

func toCost(costDef map[Resource]uint64) Cost {
	var cost Cost
	for k, v := range costDef {
		cost[k-1] = v
	}
	return cost
}

