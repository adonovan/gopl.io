// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.

//Modifed by Douglas Will
//For excercise 5.11 the instructions were to detect cycles,
//IOW - courses that a requriements of each other
//It was necessary to create the cycle in the data. "linear algebra": {"calculus"} was added
//The section markd as ---Cycle Test--- was added to check for cycles
//
package main

import (
	"fmt"
	"sort"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"linear algebra": {"calculus"},
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {

			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {

		//----Cycle Test---------
		//fmt.Printf("Key: %v\n", key)
		//fmt.Printf("\tReq: %v\n", m[key])

		//Iterate thorugh the range of requiremts for the course
		for cycle := range m[key] {
			//fmt.Printf("\tReq: %v\n", m[key][cycle])

			//Check to see any of the requirements are listed courses
			_, ok := m[m[key][cycle]]
			if ok {
				//fmt.Printf("\tThe Req '%v' is a listed class.\n", m[key][cycle])
				//fmt.Printf("\tReqs of listed class: %v\n", m[m[key][cycle]])

				//iterate through th list of requriements for the matching course
				for cycle2 := range m[m[key][cycle]] {
					//fmt.Printf("\t\tReq: %v\n", m[m[key][cycle]][cycle2])

					//fmt.Printf("\n Cycle test '%v' - '%v'\n\n", key, m[m[key][cycle]][cycle2])

					//if the current req is the same as the orginal course, then ther is cycle
					if m[m[key][cycle]][cycle2] == key {
						fmt.Printf("\nA Cycle has been detected!!!\n")
						fmt.Printf("\t'%v' is a requirement of '%v' and \n", m[m[key][cycle]][cycle2], m[key][cycle])
						fmt.Printf("\t'%v' is a requirement of '%v'\n", m[key][cycle], m[m[key][cycle]][cycle2])
					}
				}
			}

		}
		//-------End Cycle test-------

		keys = append(keys, key)
	}

	sort.Strings(keys)
	//fmt.Printf("initial sort: %v\n", keys)
	visitAll(keys)
	return order
}

func contains(strSlice []string, searchStr string) bool {
	for _, value := range strSlice {
		if value == searchStr {
			return true
		}
	}
	return false
}

//!-main
