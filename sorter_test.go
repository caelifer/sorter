package sorter_test

import (
	"math/rand"
	"testing"
	"sort"

	"github.com/caelifer/sorter"
)

type V struct {
	Name   string
	Number int
}

type rule struct {
	Name string
	Rule sorter.Rule
}

var tests = []struct {
	rule rule
	want []V
}{
	{
		rule: rule{
			Name: "Number",
			Rule: func(i1, i2 interface{}) bool {
				return i1.(V).Number < i2.(V).Number
			},
		},
		want: []V{{Name: "BBB", Number: 10}, {Name: "AAA", Number: 10}, {Name: "CCC", Number: 16}, {Name: "AAA", Number: 20}, {Name: "DDD", Number: 20}},
	},
	{
		rule: rule{
			Name: "Name",
			Rule: func(i1, i2 interface{}) bool {
				return i1.(V).Name < i2.(V).Name
			},
		},
		want: []V{{Name: "AAA", Number: 20}, {Name: "AAA", Number: 10}, {Name: "BBB", Number: 10}, {Name: "CCC", Number: 16}, {Name: "DDD", Number: 20}},
	},
	{
		rule: rule{
			Name: "NumberThenName",
			Rule: func(i1, i2 interface{}) bool {
				n1, n2 := i1.(V), i2.(V)

				switch {
				case n1.Number < n2.Number:
					return true
				case n2.Number < n1.Number:
					return false
				default:
					return n1.Name < n2.Name
				}
			},
		},
		want: []V{{Name: "AAA", Number: 10}, {Name: "BBB", Number: 10}, {Name: "CCC", Number: 16}, {Name: "AAA", Number: 20}, {Name: "DDD", Number: 20}},
	},
}

var values = []V{{"BBB", 10}, {"AAA", 20}, {"CCC", 16}, {"DDD", 20}, {"AAA", 10}}

func TestSorter(t *testing.T) {
	for _, tst := range tests {

		// Clone and convert our values to []interface{}
		clone := cloneSlice(values)
		want := cloneSlice(tst.want)

		// Test our sorter interface
		sorter.Sort(clone).By(tst.rule.Rule)

		if !equal(clone, want) {
			t.Errorf("[FAILED] testing rule %s\n\tgot:    %+v\n\twanted: %+v", tst.rule.Name, clone, want)
		}
	}
}

func equal(v1, v2 []interface{}) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i := 0; i < len(v1); i++ {
		i1, ok1 := v1[i].(V)
		i2, ok2 := v2[i].(V)

		if !(ok1 && ok2) || i1 != i2 {
			return false
		}
	}

	return true
}

func cloneSlice(old []V) []interface{} {
	res := make([]interface{}, len(old))
	for i, v := range old {
		res[i] = v
	}
	return res
}

var testInts []int

func init() {
	rand.Seed(0) // deterministic rand generator
	testInts = genRandInts()
}

func genRandInts() []int {
	const size = 1000
	ints := make([]int, size)

	for i := 0; i < size; i++ {
		ints[i] = rand.Intn(size)
	}

	return ints
}

func cloneInts(old []int) []int {
	res := make([]int, len(old))
	for i, v := range old {
		res[i] = v
	}
	return res
}

func cloneGens(old []int) []interface{} {
	res := make([]interface{}, len(old))
	for i, v := range old {
		res[i] = v
	}
	return res
}

func BenchmarkStdSort(b *testing.B) {
	is := cloneInts(testInts)
	for i := 0; i < b.N; i++ {
		sort.Ints(is)
	}
}

func BenchmarkSorter_(b *testing.B) {
	is := cloneGens(testInts)
	for i := 0; i < b.N; i++ {
		sorter.Sort(is).By(func(i1, i2 interface{}) bool {
			return i1.(int) < i2.(int)
		})
	}
}
