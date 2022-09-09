package client

import (
	"ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	series.GetFibonacciSeries5(2)
}
