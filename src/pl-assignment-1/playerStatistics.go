package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type PlayerStats struct {
	FirstName string
	LastName  string
	Average   float64
	Slugging  float64
	OnBase    float64
}

type ByLastName []PlayerStats

func (a ByLastName) Len() int           { return len(a) }
func (a ByLastName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLastName) Less(i, j int) bool { return a[i].LastName < a[j].LastName }

func main() {

	var welcomeMessage = "Welcome to the player statistics calculator test program. " +
		"I am going to \nread playersfrom an input data file." +
		"You will tell me the names of \nyour input and output files. " +
		"I will store all of the playersin a list,\ncompute each player's averages " +
		"and then write the resulting team report to \nyour output file.\n\n"
	fmt.Println(welcomeMessage)

	// Taking the file input name
	var FileName string
	var TotalPlayers int
	var BattingAverage float64
	var OverallBattingAverage float64
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
	TotalPlayers = len(playerData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var onePlayerStats PlayerStats
	var allPlayerStats []PlayerStats

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

		BattingAverage = BattingAverage + onePlayerStats.Average
		allPlayerStats = append(allPlayerStats, onePlayerStats)
	}
	OverallBattingAverage = BattingAverage / float64(TotalPlayers)
	sort.Sort(ByLastName(allPlayerStats))

	fmt.Println("BASEBALL TEAM REPORT --- ", TotalPlayers, " PLAYERS FOUND IN FILE")
	fmt.Print("OVERALL BATTING AVERAGE is ")
	fmt.Printf("%.3f\n", OverallBattingAverage)

	fmt.Println("\n \t PLAYER NAME \t\t:\t AVERAGE\tSLUGGING\tONBASE%")
	fmt.Println("-------------------------------------------------------------------------------")

	for _, eachPlayer := range allPlayerStats {
		fmt.Print("\t", eachPlayer.LastName, ", ")
		fmt.Print(eachPlayer.FirstName, "\t\t:")
		fmt.Printf("\t %.3f \t", eachPlayer.Average)
		fmt.Printf("\t %.3f \t", eachPlayer.Slugging)
		fmt.Printf("\t %.3f \n", eachPlayer.OnBase)
	}
}
