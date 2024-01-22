package main

import (
	"testing"

	"github.com/montanaflynn/stats"
)

func TestCalculateLinearRegression(t *testing.T) {
	// Set up test data
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}
	data := createCoordinateSlice(x, y)

	slope, yIntercept := calculateLinearRegression(data)

	//Round outputs to correct decimal place
	slope, _ = stats.Round(slope, 4)
	yIntercept, _ = stats.Round(yIntercept, 4)

	//Answers comes from Python output
	if slope != 0.5000 || yIntercept != 3.0009 {
		t.Errorf("Unexpected values for slope or y-intercept")
		if slope != 0.5000 {
			t.Errorf("Expected slope: 0.5000\tActual slope: %.4f", slope)
		}
		if yIntercept != 3.0009 {
			t.Errorf("Expected y-intercept: 3.0009\tActual y-intercept: %.4f", slope)
		}
	}
}

func TestCalculateRSquaredAndAdjustedRSquared(t *testing.T) {
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}
	data := createCoordinateSlice(x, y)
	slope, yIntercept := calculateLinearRegression(data)

	yHat := calculateYHat(x, slope, yIntercept)

	sse, ssto := calculateSSEandSSTO(y, yHat)

	rSquared, adjustedRSquared := calculateRSquaredAndAdjustedRSquared(x, y, sse, ssto)

	rSquared, _ = stats.Round(rSquared, 3)
	adjustedRSquared, _ = stats.Round(adjustedRSquared, 3)

	if rSquared != 0.666 || adjustedRSquared != 0.629 {
		t.Errorf("Unexpected values for R-squared or adjusted R-squared")
		if rSquared != 0.666 {
			t.Errorf("Expected r-squared: 0.666\tActual r-squared: %.3f", rSquared)
		}
		if adjustedRSquared != 0.629 {
			t.Errorf("Expected adjusted r-squared: 3.0009\tActual adjusted r-squared: %.3f", adjustedRSquared)
		}
	}
}

func TestGetFStatistic(t *testing.T) {
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}
	data := createCoordinateSlice(x, y)
	slope, yIntercept := calculateLinearRegression(data)

	yHat := calculateYHat(x, slope, yIntercept)

	sse, ssto := calculateSSEandSSTO(y, yHat)

	fStatistic, prob := getFStatistic(x, ssto, sse)

	fStatistic, _ = stats.Round(fStatistic, 2)
	prob, _ = stats.Round(prob, 5)

	expectedFStatistic := 17.97
	expectedProb := 0.00218
	if fStatistic != expectedFStatistic || prob != expectedProb {
		t.Errorf("Unexpected values for F-statistic or p-value")
		if fStatistic != expectedFStatistic {
			t.Errorf("Expected F-statistic: %.2f\tActual F-statistic: %.2f", expectedFStatistic, fStatistic)
		}
		if prob != expectedProb {
			t.Errorf("Expected adjusted p-value: %.5f\tActual adjusted p-value: %.5f", expectedProb, prob)
		}
	}
}

func TestMain(m *testing.M) {
	// Run the tests
	m.Run()
}
