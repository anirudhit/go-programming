package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	Name string
	Age  int
	Job  string
}

func main() {
	// read data from CSV file

	csvFile, err := os.Open("./tabdata.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.Comma = ' ' // Use tab-delimited instead of comma <---- here!

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord Employee
	var allRecords []Employee

	for _, each := range csvData {
		fmt.Println(each)
		fmt.Println(len(each))
		for i := 0; i < len(each); i++ {
			switch i {
			case 0:
				oneRecord.Name = each[i]
				break
			case 1:
				oneRecord.Age, _ = strconv.Atoi(each[i]) // need to cast integer to string
				break
			case 2:
				oneRecord.Job = each[i]
				break
			default:
			}
		}
		allRecords = append(allRecords, oneRecord)
	}

	jsondata, err := json.Marshal(allRecords) // convert to JSON

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sanity check
	// NOTE : You can stream the JSON data to http service as well instead of saving to file
	fmt.Println("1.", string(jsondata))
}
