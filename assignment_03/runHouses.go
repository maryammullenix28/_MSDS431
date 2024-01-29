package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Specify the input and output file names
	inputFileName := "/Users/maryammullenix/Documents/GitHub/_MSDS431/assignment_03/housesInput.csv"
	outputFileName := "housesOutputGo.txt"

	// Read the CSV file
	data, err := readCSV(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	// Compute summary statistics for each column
	summary := make(map[string]SummaryStats)
	columnNames := []string{"value", "income", "age", "rooms", "bedrooms", "pop", "hh"}

	for _, columnName := range columnNames {
		values := extractColumn(data, columnName)
		stats := computeSummaryStats(values)
		summary[columnName] = stats
	}

	// Write the summary statistics to a text file
	err = writeSummaryToFile(outputFileName, columnNames, summary)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Summary statistics have been written to", outputFileName)
}

// SummaryStats represents the summary statistics for a column
type SummaryStats struct {
	Min    float64
	FirstQ float64
	Median float64
	Mean   float64
	ThirdQ float64
	Max    float64
}

// readCSV reads a CSV file and returns a 2D slice of strings
func readCSV(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// extractColumn extracts a specific column from a 2D slice of strings
func extractColumn(data [][]string, columnName string) []float64 {
	var columnValues []float64
	index := indexOfColumn(data[0], columnName)

	for i := 1; i < len(data); i++ {
		value, err := strconv.ParseFloat(data[i][index], 64)
		if err != nil {
			log.Fatal(err)
		}
		columnValues = append(columnValues, value)
	}

	return columnValues
}

// indexOfColumn finds the index of a column in the header
func indexOfColumn(header []string, columnName string) int {
	for i, column := range header {
		if column == columnName {
			return i
		}
	}
	return -1
}

// computeSummaryStats computes summary statistics for a given set of values
func computeSummaryStats(values []float64) SummaryStats {
	sort.Float64s(values)

	n := len(values)
	min := values[0]
	max := values[n-1]
	mean := computeMean(values)
	firstQ := computePercentile(values, 0.25)
	median := computePercentile(values, 0.5)
	thirdQ := computePercentile(values, 0.75)

	return SummaryStats{
		Min:    min,
		FirstQ: firstQ,
		Median: median,
		Mean:   mean,
		ThirdQ: thirdQ,
		Max:    max,
	}
}

// computeMean computes the mean of a set of values
func computeMean(values []float64) float64 {
	sum := 0.0
	for _, value := range values {
		sum += value
	}
	return sum / float64(len(values))
}

// computePercentile computes the percentile of a set of values
func computePercentile(values []float64, percentile float64) float64 {
	index := int(percentile * float64(len(values)-1))
	return values[index]
}

// writeSummaryToFile writes summary statistics to a text file
func writeSummaryToFile(fileName string, columnNames []string, summary map[string]SummaryStats) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, columnName := range columnNames {
		stats := summary[columnName]
		fmt.Fprintf(file, "%-12s%-12s%-12s%-12s%-12s%-12s%-12s\n", columnName, "Min.", "1st Qu.", "Median", "Mean", "3rd Qu.", "Max.")
		fmt.Fprintf(file, "%-12s%-12.2f%-12.2f%-12.2f%-12.2f%-12.2f%-12.2f\n", "", stats.Min, stats.FirstQ, stats.Median, stats.Mean, stats.ThirdQ, stats.Max)
	}

	return nil
}
