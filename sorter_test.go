package sorter_test

import (
	"testing"

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
	want []interface{}
}{
	{
		rule: rule{
			Name: "Number",
			Rule: func(i1, i2 interface{}) bool {
				return i1.(V).Number < i2.(V).Number
			},
		},
		want: []interface{}{V{Name: "BBB", Number: 10}, V{Name: "AAA", Number: 10}, V{Name: "CCC", Number: 16}, V{Name: "AAA", Number: 20}, V{Name: "DDD", Number: 20}},
	},
	{
		rule: rule{
			Name: "Name",
			Rule: func(i1, i2 interface{}) bool {
				return i1.(V).Name < i2.(V).Name
			},
		},
		want: []interface{}{V{Name: "AAA", Number: 20}, V{Name: "AAA", Number: 10}, V{Name: "BBB", Number: 10}, V{Name: "CCC", Number: 16}, V{Name: "DDD", Number: 20}},
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
		want: []interface{}{V{Name: "AAA", Number: 10}, V{Name: "BBB", Number: 10}, V{Name: "CCC", Number: 16}, V{Name: "AAA", Number: 20}, V{Name: "DDD", Number: 20}},
	},
}

var values = []interface{}{V{"BBB", 10}, V{"AAA", 20}, V{"CCC", 16}, V{"DDD", 20}, V{"AAA", 10}}

func TestSorter(t *testing.T) {
	for _, tst := range tests {

		// Test our sorter interface
		clone := cloneSlice(values)

		sorter.Sort(clone).By(tst.rule.Rule)

		if !equal(clone, tst.want) {
			t.Errorf("[FAILED] testing rule %s\n\tgot:    %+v\n\twanted: %+v", tst.rule.Name, clone, tst.want)
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

func cloneSlice(old []interface{}) []interface{} {
	return append([]interface{}{}, old...)
}
