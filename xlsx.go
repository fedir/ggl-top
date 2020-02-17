package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func writeXLSXData(data []GoogleResult, xlsxFilePath string) {

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	defer func() {
		err = file.Save(xlsxFilePath)
		if err != nil {
			fmt.Printf(err.Error())
		}
	}()

	sheet, err = file.AddSheet("Google Links")
	if err != nil {
		fmt.Printf(err.Error())
	}

	for _, value := range data {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = value.URL
		cell = row.AddCell()
		cell.Value = value.Title
		cell = row.AddCell()
		cell.Value = value.Description
		err = file.Save(xlsxFilePath)
		if err != nil {
			fmt.Printf(err.Error())
		}
	}

}
