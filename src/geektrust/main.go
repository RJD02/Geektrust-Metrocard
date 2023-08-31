package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type AppConfig struct {
	Stations map[string]*Station
	Users    map[string]*MetroCard
	cost     map[citizenType]int
}

func (app *AppConfig) setup() {
	app.Stations = make(map[string]*Station)
	app.Stations[string(AIRPORT)] = &Station{AmountCollected: 0, DiscountGiven: 0, CitizenCount: make(map[citizenType]int)}
	app.Stations[CENTRAL] = &Station{AmountCollected: 0, DiscountGiven: 0, CitizenCount: make(map[citizenType]int)}
	app.Users = make(map[string]*MetroCard)
	app.cost = make(map[citizenType]int)
	app.cost[SENIOR_CITIZEN] = 100
	app.cost[ADULT] = 200
	app.cost[KID] = 50
}

var app AppConfig

func evaluate(argList []string) {
	switch argList[0] {
	case string(BALANCE):
		balanceHandler(argList[1], argList[2])
	case CHECK_IN:
		checkInHandler(argList[1], citizenType(argList[2]), locationTypes(argList[3]))
	case PRINT_SUMMARY:
		printSummary()
	}
}

func main() {
	app = AppConfig{}
	app.setup()

	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")

		return
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening the input file")

		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		/*
			args := scanner.Text()
			argList := strings.Fields(args)

			Add your code here to process the input commands
		*/
		args := scanner.Text()
		argList := strings.Fields(args)
		evaluate(argList)
	}
}
