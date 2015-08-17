package sorter

import "sort"

// Rule is a generic comparator. Client must provide its own specific implementation per each type.
type Rule func(i1, i2 interface{}) bool

// Sorter
type Sorter interface {
	By(Rule)
}

// Sort takes a slice of interface{} values and returns a Sorter object
func Sort(vals []interface{}) Sorter {
	return &genericSorter{vals, nil}
}

// Private implementation details

// genericSorter is a type that implements Sorter interface.
type genericSorter struct {
	vals []interface{}
	rule Rule
}

// By performs actual sort based on the rule provided. It implements Sorter interface.
func (gs *genericSorter) By(r Rule) {
	gs.rule = r
	sort.Sort(gs)
}

// Implement sort.Interface interface

func (gs *genericSorter) Len() int {
	return len(gs.vals)
}

func (gs *genericSorter) Swap(i, j int) {
	gs.vals[i], gs.vals[j] = gs.vals[j], gs.vals[i]
}

func (gs *genericSorter) Less(i, j int) bool {
	return gs.rule(gs.vals[i], gs.vals[j])
}
