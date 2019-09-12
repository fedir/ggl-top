package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func writeCSVData(data []GoogleResult, csvFilePath string) {
	var csvData [][]string
	for _, gr := range data {
		row := []string{
			fmt.Sprintf("%d", gr.Rank),
			gr.URL,
			gr.Title,
			gr.Description,
		}
		csvData = append(csvData, row)

	}
	writeCSV(csvFilePath, csvData)
}

func writeCSV(csvFilePath string, ghDataCSV [][]string) {
	file, err := os.Create(csvFilePath)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range ghDataCSV {
		err := writer.Write(value)
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
}
