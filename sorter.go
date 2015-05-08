package sorter

import (
	"sort"
)

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

func (gc *genericSorter) By(r Rule) {
	gc.rule = r
	sort.Sort(gc)
}

// Implement sort.Interface interface

func (gc *genericSorter) Len() int {
	return len(gc.vals)
}

func (gc *genericSorter) Swap(i, j int) {
	gc.vals[i], gc.vals[j] = gc.vals[j], gc.vals[i]
}

func (gc *genericSorter) Less(i, j int) bool {
	return gc.rule(gc.vals[i], gc.vals[j])
}
