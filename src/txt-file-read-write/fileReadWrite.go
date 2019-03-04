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

	//reader.Comma = '\t' // Use tab-delimited instead of comma <---- here!

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord Employee
	var allRecords []Employee

	for _, each := range csvData {
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

		// oneRecord.Name = each[0]
		// oneRecord.Age, _ = strconv.Atoi(each[1]) // need to cast integer to string
		// oneRecord.Job = each[2]
		allRecords = append(allRecords, oneRecord)
	}

	jsondata, err := json.Marshal(allRecords) // convert to JSON

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sanity check
	// NOTE : You can stream the JSON data to http service as well instead of saving to file
	fmt.Println(string(jsondata))

	// now write to JSON file

	jsonFile, err := os.Create("./data.json")

	if err != nil {
		fmt.Println(err)
	}

	var oneRecordWrite Employee
	var allRecordsWrite []Employee

	for _, each := range csvData {
		for i := 0; i < len(each); i++ {
			switch i {
			case 0:
				oneRecordWrite.Name = each[i]
				break
			case 1:
				oneRecordWrite.Age, _ = strconv.Atoi(each[i]) // need to cast integer to string
				break
			case 2:
				oneRecordWrite.Job = each[i]
				break
			default:
			}
		}
		allRecordsWrite = append(allRecordsWrite, oneRecordWrite)
	}

	jsondataWrite, errWrite := json.Marshal(allRecordsWrite) // convert to JSON

	if errWrite != nil {
		fmt.Println(errWrite)
		os.Exit(1)
	}

	// sanity check
	// NOTE : You can stream the JSON data to http service as well instead of saving to file
	fmt.Println(string(jsondataWrite))

	// now write to JSON file

	jsonFileWrite, errFileWrite := os.Create("./data.json")

	if errFileWrite != nil {
		fmt.Println(errFileWrite)
	}
	defer jsonFile.Close()

	jsonFileWrite.Write(jsondataWrite)
	jsonFileWrite.Close()
}
