package main

type Station struct {
	AmountCollected int
	DiscountGiven   int
	CitizenCount    map[citizenType]int
}

func createStation(id string) {
	newStation := &Station{AmountCollected: 0, DiscountGiven: 0}
	app.Stations[id] = newStation
}
