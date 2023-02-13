package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	// Open the .xlsx file
	file, err := excelize.OpenFile("icu_prediction.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get all the rows in the first sheet and preprocess the data
	rows := file.GetRows("Sheet1")
	var data = preprocess(rows)
	fmt.Println(data[0][0])
	//
	//
	//

}

func preprocess(rows [][]string) [][]float64 {
	// prepare variable for transformed dataset
	var floatData [][]float64
	var columnCounter [231]float64
	var columnSum [231]float64

	// First loop converts all features to numerical values and also the sum of all columns
	for _, row := range rows[1:] {

		// Prepare age_percentile column to be converted to numerical
		value := strings.Replace(row[2], "Above ", "", -1)
		index := strings.Index(value, "th")
		value = value[:index]
		row[2] = value

		// Prepare window column to be converted to numerical
		if row[229] == "ABOVE_12" {
			row[229] = "12-more"
		}
		re := regexp.MustCompile(`(.+?)-`)
		match := re.FindStringSubmatch(row[229])
		if len(match) == 0 {
			fmt.Println("No match found")
		} else {
			row[229] = match[1]
		}

		// Convert all values in the dataset into numericals and ignore empty cells
		var floatRow []float64
		var missingCounter = 0
		for index, value := range row {
			num, err := strconv.ParseFloat(value, 64)
			if err != nil {
				missingCounter++
				floatRow = append(floatRow, math.NaN())
			} else {
				columnSum[index] += num
				columnCounter[index]++
				floatRow = append(floatRow, num)
			}
		}
		floatRow = append(floatRow, float64(missingCounter))
		floatData = append(floatData, floatRow)
	}

	// Second loop finds mean value of all columns
	for i, _ := range columnSum {
		columnSum[i] = columnSum[i] / columnCounter[i]
	}

	// Third loop imputes all the missing values with the mean value found in loop 2
	for i, row := range floatData {
		for j, value := range row {
			if math.IsNaN(value) {
				floatData[i][j] = columnSum[j]
			}
		}
	}

	return floatData
}
