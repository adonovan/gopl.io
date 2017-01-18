// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//Modified by Douglas Will 2017

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.

//Modifed by Douglas Will
//For excercise 5.10 the instructions were not real clear.  I inturpretted it to  mean
//use maps for the pre-requisite course in the prereqs map. I had to modify func topoSort
//To take a map of maps instead of a map of strings.
//Func visitAll had to be modified to iterate over a map key and values instead of []string
//As mentioned in the book, the output order is now non-deterministic

package main

import (
	"fmt"
	//"sort"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string]map[int]string{
	"algorithms":            {1: "data structures"},
	"calculus":              {1: "linear algebra"},
	"compilers":             {1: "data structures", 2: "formal languages", 3: "computer organization"},
	"data structures":       {1: "discrete math"},
	"databases":             {1: "data structures"},
	"discrete math":         {1: "intro to programming"},
	"formal languages":      {1: "discrete math"},
	"networks":              {1: "operating systems"},
	"operating systems":     {1: "data structures", 2: "computer organization"},
	"programming languages": {1: "data structures", 2: "computer organization"},
	"linear algebra":        {1: "calculus"},
}

//!-table

//!+main
func main() {
	for i, course := range topoSort(prereqs) {
		//fmt.Printf("Back to Main\n")
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[int]string) []string {
	//fmt.Printf("In Toposort\n")
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[int]string)

	visitAll = func(items map[int]string) {
		//fmt.Printf("In VisitAll\n")
		for _, item := range items {
			//fmt.Printf("Item: %v\n", item)
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	//fmt.Printf("map: %v\n\n", m)
	keys := make(map[int]string)
	keycount := 1
	for key := range m {
		keys[keycount] = key
		keycount++

	}

	//	sort.Strings(keys)
	//fmt.Printf("initail sort: %v\n\n", keys)
	visitAll(keys)
	return order
}

//!-main
