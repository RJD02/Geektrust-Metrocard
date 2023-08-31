package main

import (
	"fmt"
	"strconv"
)

type commandTypes string

const (
	BALANCE       commandTypes = "BALANCE"
	CHECK_IN                   = "CHECK_IN"
	PRINT_SUMMARY              = "PRINT_SUMMARY"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func balanceHandler(id string, amount string) {
	intAmount, err := strconv.Atoi(amount)
	if err != nil {
		panic("Ill formatted text")
	}
	app.Users[id] = &MetroCard{Amount: intAmount, CurrLocation: UNDEFINED_STRING, CitizentType: UNDEFINED_STRING, count: 1}
}

func checkInHandler(id string, citizen citizenType, fromLocation locationTypes) {
	if _, ok := app.Users[id]; !ok {
		panic("No user with that id present")
	}
	app.Stations[string(fromLocation)].CitizenCount[citizen] += 1
	reqAmount := app.cost[citizen]
	currBalance := app.Users[id].Amount
	if app.Users[id].count%2 == 0 {
		app.Stations[string(fromLocation)].DiscountGiven += reqAmount / 2
        reqAmount /= 2
	} else {
		app.Users[id].count += 1
	}
	if reqAmount <= currBalance {
		app.Stations[string(fromLocation)].AmountCollected += reqAmount
		app.Users[id].Amount -= reqAmount
	} else {
		app.Stations[string(fromLocation)].AmountCollected += int((reqAmount - currBalance) * 2 / 100)
		app.Stations[string(fromLocation)].AmountCollected += reqAmount
		app.Users[id].Amount = 0
	}
	if fromLocation == CENTRAL {
		app.Users[id].CurrLocation = AIRPORT
	} else if fromLocation == AIRPORT {
		app.Users[id].CurrLocation = CENTRAL
	}
	fmt.Println(citizen, app.cost[citizen])
	fmt.Println("Users:")
	for k, v := range app.Users {
		fmt.Println(k, v)
	}
	fmt.Println("Stations:")
	for k, v := range app.Stations {
		fmt.Println(k, v)
	}
}

func printSummaryUtil(stationName locationTypes) {
	fmt.Println("TOTLA_COLLECTION", stationName, app.Stations[string(stationName)].AmountCollected, app.Stations[string(stationName)].DiscountGiven)
	fmt.Println("PASSENGER_TYPE_SUMMARY")
	for k, v := range app.Stations[string(stationName)].CitizenCount {
		fmt.Println(k, v)
	}
}

func printSummary() {
	printSummaryUtil(CENTRAL)
	printSummaryUtil(AIRPORT)
}
