package main

import (
	"fmt"

	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/stat/distuv"

	"time"
)

func createCoordinateSlice(x, y []float64) []stats.Coordinate {
	data := make([]stats.Coordinate, len(x))
	for i := range x {
		data[i] = stats.Coordinate{X: float64(x[i]), Y: y[i]}
	}
	return data
}

func calculateLinearRegression(data []stats.Coordinate) (float64, float64) {
	r, _ := stats.LinearRegression(data)
	slope := (r[len(r)-1].Y - r[0].Y) / (r[len(r)-1].X - r[0].X)
	yIntercept := r[0].Y - slope*r[0].X
	return slope, yIntercept
}

func calculateYHat(x []float64, slope, yIntercept float64) []float64 {
	yHat := make([]float64, len(x))
	for i, val := range x {
		yHat[i] = slope*val + yIntercept
	}
	return yHat
}

func calculateSSEandSSTO(y, yHat []float64) (float64, float64) {
	var sse float64
	for i := range y {
		sse += (y[i] - yHat[i]) * (y[i] - yHat[i])
	}

	yBar, _ := stats.Mean(y)
	var ssto float64
	for i := range y {
		ssto += (y[i] - yBar) * (y[i] - yBar)
	}

	return sse, ssto
}

func calculateRSquaredAndAdjustedRSquared(x []float64, y []float64, sse, ssto float64) (float64, float64) {
	//Calculate R-Squared
	rSquared := 1 - sse/ssto

	//Calculate adjusted R-Squared
	df1 := float64(1)

	n := float64(len(x))
	adjustedRSquared := 1 - ((1 - rSquared) * (n - 1) / (n - df1 - 1))

	//Return variables
	return rSquared, adjustedRSquared
}

func getFStatistic(x []float64, ssto, sse float64) (float64, float64) {
	df1 := float64(1) // number of predictors (slope)
	df2 := len(x) - 2 // number of observations - number of predictors - 1

	// Calculate F-statistic
	fStatistic := (ssto - sse) / float64(df1) / (sse / float64(df2))

	fDist := distuv.F{
		D1: float64(df1),
		D2: float64(df2),
	}

	// Calculate Prob (F-statistic) using the CDF of the F-distribution
	prob := 1 - fDist.CDF(fStatistic)

	return fStatistic, prob
}

func getLinearRegressionSummary(x, y []float64, variable string) {
	// Create coordinate slice
	data := createCoordinateSlice(x, y)

	// Calculate linear regression
	slope, yIntercept := calculateLinearRegression(data)

	// Calculate yHat
	yHat := calculateYHat(x, slope, yIntercept)

	// Calculate SSE and SSTO
	sse, ssto := calculateSSEandSSTO(y, yHat)

	// Calculate R-squared value
	rSquared, adjustedRSquared := calculateRSquaredAndAdjustedRSquared(x, y, sse, ssto)

	//Calculate f-statistic
	fStat, prob := getFStatistic(x, ssto, sse)

	// Print summary
	fmt.Println("\t\t\tLinear Regression Results")
	fmt.Println("===============================================================================")
	fmt.Printf("Dep. Variable:               %s\t\t\tR-squared:\t\t%.3f\n", variable, rSquared)
	fmt.Printf("Date:                        %s\tAdjusted R-squared:\t%.3f\n", time.Now().Format("Mon, 02 Jan 2006"), adjustedRSquared)
	fmt.Printf("Time:                        %s\t\tF-Statistic:\t\t%.2f\n", time.Now().Format("15:04:05"), fStat)
	fmt.Printf("Number of Observations:      %d\t\t\tProb (F-Statistic):\t%.5f\n", len(x), prob)
	fmt.Printf("y-intercept/const            %.4f\t\tSlope:\t\t\t%.4f\n", yIntercept, slope)
	fmt.Println("===============================================================================")
	fmt.Println()
}

func main() {

	//Declare variables using slices
	var x1 = []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	var x2 = []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	var x3 = []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	var x4 = []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}

	var y1 = []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}
	var y2 = []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}
	var y3 = []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}
	var y4 = []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89}

	//Get linear regression details using custom functions built on montanaflynn/stats and gonum/distuv
	getLinearRegressionSummary(x1, y1, "y1")
	getLinearRegressionSummary(x2, y2, "y2")
	getLinearRegressionSummary(x3, y3, "y3")
	getLinearRegressionSummary(x4, y4, "y4")

}
