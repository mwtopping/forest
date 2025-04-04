package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func read_csv(filename string) map[string][]float64 {
	fmt.Println("Reading in csv data from ", filename)

	output := make(map[string][]float64)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file", filename)
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	r := csv.NewReader(f)

	// read out the first line
	column_names, err := r.Read()
	name_indices := make(map[int]string)
	for i, c := range column_names {
		output[c] = make([]float64, 0)
		name_indices[i] = c
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for i, val := range record {
			column_name := name_indices[i]
			fval, err := strconv.ParseFloat(val, 64)
			if err != nil {
				fmt.Println("Error converting entry to float")
				fmt.Println(err)
				os.Exit(1)
			}
			output[column_name] = append(output[column_name], fval)
		}
	}

	return output
}
