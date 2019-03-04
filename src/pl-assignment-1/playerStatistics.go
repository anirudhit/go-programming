package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Player struct {
	FirstName        string
	LastName         string
	PlateAppearances int
	AtBats           int
	Singles          int
	Doubles          int
	Triples          int
	HomeRuns         int
	Walks            int
	HitByPitch       int
}

type PlayerStats struct {
	FirstName string
	LastName  string
	Average   float64
	Slugging  float64
	OnBase    float64
}

func main() {

	var welcomeMessage = "Welcome to the player statisticscalculator test program.  I am going to \nread playersfrom an input data file." +
		"You will tell me the names of \nyour input and output files.  I will store all of the playersin a list,\ncompute each player's averages" +
		"and then write the resulting team report to \nyour output file.\n\n"
	fmt.Println(welcomeMessage)

	// Taking the file input name
	var FileName string
	fmt.Print("Enter the name of your input file:	")
	fmt.Scan(&FileName)

	// read data from CSV file
	playerFile, err := os.Open(FileName)

	if err != nil {
		fmt.Println(err)
	}

	defer playerFile.Close()
	reader := csv.NewReader(playerFile)
	reader.Comma = ' ' // Use tab-delimited instead of comma <---- here!
	reader.FieldsPerRecord = -1

	playerData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Println("playerData: ", len(playerData))

	// var onePlayer Player
	// var allPlayers []Player

	// for _, eachPlayer := range playerData {
	// 	onePlayer.FirstName = eachPlayer[0]
	// 	onePlayer.LastName = eachPlayer[1]
	// 	onePlayer.PlateAppearances, _ = strconv.Atoi(eachPlayer[2])
	// 	onePlayer.AtBats, _ = strconv.Atoi(eachPlayer[3])
	// 	onePlayer.Singles, _ = strconv.Atoi(eachPlayer[4])
	// 	onePlayer.Doubles, _ = strconv.Atoi(eachPlayer[5])
	// 	onePlayer.Triples, _ = strconv.Atoi(eachPlayer[6])
	// 	onePlayer.HomeRuns, _ = strconv.Atoi(eachPlayer[7])
	// 	onePlayer.Walks, _ = strconv.Atoi(eachPlayer[8])
	// 	onePlayer.HitByPitch, _ = strconv.Atoi(eachPlayer[9])

	// 	allPlayers = append(allPlayers, onePlayer)
	// }

	var onePlayerStats PlayerStats
	var allPlayerStats []PlayerStats

	fmt.Println("BASEBALL TEAM REPORT --- ", len(playerData), " PLAYERS FOUND IN FILE")
	fmt.Println("OVERALL BATTING AVERAGE is 0.290")

	fmt.Println("\n \t PLAYER NAME \t\t:\t AVERAGE\tSLUGGING\tONBASE%")
	fmt.Println("-------------------------------------------------------------------------------")

	for _, eachPlayer := range playerData {
		onePlayerStats.FirstName = eachPlayer[0]
		onePlayerStats.LastName = eachPlayer[1]

		PlateAppearances, _ := strconv.ParseFloat(eachPlayer[2], 64)
		AtBats, _ := strconv.ParseFloat(eachPlayer[3], 64)
		Singles, _ := strconv.ParseFloat(eachPlayer[4], 64)
		Doubles, _ := strconv.ParseFloat(eachPlayer[5], 64)
		Triples, _ := strconv.ParseFloat(eachPlayer[6], 64)
		HomeRuns, _ := strconv.ParseFloat(eachPlayer[7], 64)
		Walks, _ := strconv.ParseFloat(eachPlayer[8], 64)
		HitByPitch, _ := strconv.ParseFloat(eachPlayer[9], 64)

		onePlayerStats.Average = (Singles + Doubles + Triples + HomeRuns) / (AtBats)
		onePlayerStats.Slugging = ((Singles) + (2 * Doubles) + (3 * Triples) + (4 * HomeRuns)) / (AtBats)
		onePlayerStats.OnBase = (Singles + Doubles + Triples + HomeRuns + Walks + HitByPitch) / (PlateAppearances)

		allPlayerStats = append(allPlayerStats, onePlayerStats)

		fmt.Print("\t", onePlayerStats.LastName, ", ")
		fmt.Print(onePlayerStats.FirstName, "\t\t:")
		fmt.Print("\t", toFixed(onePlayerStats.Average, 3), "\t")
		fmt.Print("\t", toFixed(onePlayerStats.Slugging, 3), "\t")
		fmt.Print("\t", toFixed(onePlayerStats.OnBase, 3), "\n")
	}

	// jsondata, err := json.Marshal(allPlayers) // convert to JSON
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println("1.", string(jsondata))
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
