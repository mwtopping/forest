package main

import (
	"fmt"
	"math"
	"slices"
)

type Node struct {
	key   string
	val   float64
	Left  *Node
	Right *Node
	// add in parent node?
}

type Tree struct {
	max_depth int
	Root      *Node
}

func (node *Node) display() {
	fmt.Println(node.key, node.val)
	fmt.Println(node.Left, node.Right)
	fmt.Println("===")
}

func (tree *Tree) display() {

	tree.Root.display()

}

func NewTree(filename string, max_depth int) Tree {
	var tree Tree
	tree.max_depth = max_depth

	// read in the data and handle that shit here
	csv_data := read_csv(filename)

	fmt.Println("Read in csv file with columns: ")
	for c := range csv_data {
		fmt.Println(c)
	}

	columns := []string{"chol", "age", "trestbps"}
	tree.Root = branch_node(columns, csv_data)

	//	//	Ncolumns := len(csv_data)
	//	split_col, split_val, _ := determine_best_split(columns, csv_data)
	//
	//	fmt.Printf("Best split on column:%v at value:%.2f\n",
	//		split_col, split_val)

	return tree
}

func branch_node(columns []string, csv_data map[string][]float64) *Node {

	// check if we're done
	Nitems := len(csv_data["target"])

	fmt.Printf("Found only %v items\n", Nitems)
	if Nitems < 200 {
		return &Node{val: 0.0}
	}

	split_col, split_val, _ := determine_best_split(columns, csv_data)

	//if we need to continue
	left, right := split_data_at_value(csv_data, split_col, split_val)
	return &Node{
		Left:  branch_node(columns, left),
		Right: branch_node(columns, right),
		key:   split_col,
		val:   split_val,
	}
}

func determine_best_split(columns []string, csv_data map[string][]float64) (col string, bound, gini float64) {
	var best_column string
	var best_boundary float64 = 0.0
	var best_gini = 1.0

	for _, c := range columns {
		sorted_values := make([]float64, len(csv_data[c]))
		copy(sorted_values, csv_data[c])
		slices.Sort(sorted_values)

		unique_values := slices.Compact(sorted_values)
		bounds := make([]float64, len(unique_values)-1)

		for i := range bounds {
			avg := 0.5 * (unique_values[i] + unique_values[i+1])
			bounds[i] = avg
		}
		// remove duplicate values from bounds

		// for given boundary
		for _, bound := range bounds {
			left, right := split_slice_targets(csv_data[c], csv_data["target"], bound)
			Nleft := float64(len(left))
			Nright := float64(len(right))
			if (Nleft == 0) || (Nright == 0) {
				break
			}
			gini_left := 1 - math.Pow((sum(left)/Nleft), 2) -
				math.Pow(((Nleft-sum(left))/Nleft), 2)

			gini_right := 1 - math.Pow((sum(right)/Nright), 2) -
				math.Pow(((Nright-sum(right))/Nright), 2)

			tot_gini := (Nleft*gini_left + Nright*gini_right) / (Nleft + Nright)

			if tot_gini < best_gini {
				best_column = c
				best_boundary = bound
				best_gini = tot_gini
			}

		}

	}
	return best_column, best_boundary, best_gini
}

func sum(arr []float64) float64 {
	total := 0.0
	for _, v := range arr {
		total += v
	}
	return total
}

func split_data_at_value(data map[string][]float64, col string, value float64) (a, b map[string][]float64) {

	left_data := make(map[string][]float64)
	right_data := make(map[string][]float64)

	for c := range data {
		left_data[c] = make([]float64, 0)
		right_data[c] = make([]float64, 0)
	}

	// IMPLEMENT
	for i := range data["target"] {
		if data[col][i] < value {
			// put it in the left
			for c := range data {
				left_data[c] = append(left_data[c], data[c][i])
			}
		} else {
			// put it in the right
			for c := range data {
				right_data[c] = append(right_data[c], data[c][i])
			}
		}
	}

	return left_data, right_data

}

func split_slice_targets(arr, targs []float64, bound float64) (a, b []float64) {

	var less []float64
	var more []float64

	for i, v := range arr {
		if v <= bound {
			less = append(less, targs[i])
		} else {
			more = append(more, targs[i])
		}
	}
	return less, more
}
