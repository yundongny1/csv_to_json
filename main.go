package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	//if less than 3 arguments, return error msg
	if len(os.Args) < 3 {
		fmt.Println("Please run the app like this: csv_to_jsonl <input.csv> <output.jsonl>")
		return
	}

	//store the input and output file name from command line argument
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	//open the csv file
	file, err := os.Open(inputFile)
	//error message for opening file
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	//A Reader reads records from a CSV-encoded file.
	csvReader := csv.NewReader(file)

	//using the first line as headers to use for parsing csv to json.
	headers, err := csvReader.Read()
	if err != nil {
		fmt.Println("Error reading CSV headers:", err)
		return
	}

	//create the output file to put the parsed data
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	//An Encoder writes JSON values to an output stream.
	jsonEncoder := json.NewEncoder(outFile)

	//read and process each record from the CSV
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading CSV:", err)
			return
		}

		//map using headers as keys and csv rows as values
		recordMap := make(map[string]string)
		for i, value := range record {
			recordMap[headers[i]] = value
		}

		//encode as jsonl
		if err := jsonEncoder.Encode(recordMap); err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
	}
	fmt.Println("CSV to JSONL conversion is a success!")
}
