package main

import (
	"fmt"

	hdrhistogram "github.com/HdrHistogram/hdrhistogram-go"
)

func main() {
	hdr := hdrhistogram.New(1, 30000000, 4)
	input := []int64{
		1, 2, 3, 4, 5, 6, 7, 8,
		9, 10,
	}

	for _, sample := range input {
		hdr.RecordValue(sample)
	}

	fmt.Printf("Min : %d\n", hdr.Min())
	fmt.Printf("Max : %d\n", hdr.Max())
	fmt.Printf("Avg : %f\n", hdr.Mean())
	fmt.Printf("Percentile 50: %d\n", hdr.ValueAtQuantile(50.0))
	fmt.Printf("Percentile 99: %d\n", hdr.ValueAtQuantile(99.0))
}
