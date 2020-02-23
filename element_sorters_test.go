package collections

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestElementsSorters(t *testing.T) {
	tests := []struct {
		name   string
		sortBy sort.Interface
		want   sort.Interface
	}{
		{
			name: "should sort by string keys ascending",
			sortBy: ElementsByKeyStringAsc{
				{Key: StringValue("z"), Value: StringValue("zz")},
				{Key: StringValue("f"), Value: StringValue("ff")},
				{Key: StringValue("a"), Value: StringValue("aa")},
				{Key: StringValue("b"), Value: StringValue("bb")},
			},
			want: ElementsByKeyStringAsc{
				{Key: StringValue("a"), Value: StringValue("aa")},
				{Key: StringValue("b"), Value: StringValue("bb")},
				{Key: StringValue("f"), Value: StringValue("ff")},
				{Key: StringValue("z"), Value: StringValue("zz")},
			},
		},
		{
			name: "should sort by string keys descending",
			sortBy: ElementsByKeyStringDesc{
				{Key: StringValue("f"), Value: StringValue("ff")},
				{Key: StringValue("a"), Value: StringValue("aa")},
				{Key: StringValue("z"), Value: StringValue("zz")},
				{Key: StringValue("b"), Value: StringValue("bb")},
			},
			want: ElementsByKeyStringDesc{
				{Key: StringValue("z"), Value: StringValue("zz")},
				{Key: StringValue("f"), Value: StringValue("ff")},
				{Key: StringValue("b"), Value: StringValue("bb")},
				{Key: StringValue("a"), Value: StringValue("aa")},
			},
		},
		{
			name: "should sort by int keys ascending",
			sortBy: ElementsByKeyIntAsc{
				{Key: IntValue(10), Value: StringValue("zz")},
				{Key: IntValue(2), Value: StringValue("ff")},
				{Key: IntValue(4), Value: StringValue("aa")},
				{Key: IntValue(50), Value: StringValue("bb")},
			},
			want: ElementsByKeyIntAsc{
				{Key: IntValue(2), Value: StringValue("ff")},
				{Key: IntValue(4), Value: StringValue("aa")},
				{Key: IntValue(10), Value: StringValue("zz")},
				{Key: IntValue(50), Value: StringValue("bb")},
			},
		},
		{
			name: "should sort by int keys descending",
			sortBy: ElementsByKeyIntDesc{
				{Key: IntValue(10), Value: StringValue("zz")},
				{Key: IntValue(2), Value: StringValue("ff")},
				{Key: IntValue(4), Value: StringValue("aa")},
				{Key: IntValue(50), Value: StringValue("bb")},
			},
			want: ElementsByKeyIntDesc{
				{Key: IntValue(50), Value: StringValue("bb")},
				{Key: IntValue(10), Value: StringValue("zz")},
				{Key: IntValue(4), Value: StringValue("aa")},
				{Key: IntValue(2), Value: StringValue("ff")},
			},
		},
		{
			name: "should sort by string values ascending",
			sortBy: ElementsByValueStringAsc{
				{Key: StringValue("z"), Value: StringValue("x")},
				{Key: StringValue("f"), Value: StringValue("p")},
				{Key: StringValue("a"), Value: StringValue("k")},
				{Key: StringValue("b"), Value: StringValue("l")},
			},
			want: ElementsByValueStringAsc{
				{Key: StringValue("a"), Value: StringValue("k")},
				{Key: StringValue("b"), Value: StringValue("l")},
				{Key: StringValue("f"), Value: StringValue("p")},
				{Key: StringValue("z"), Value: StringValue("x")},
			},
		},
		{
			name: "should sort by string values descending",
			sortBy: ElementsByValueStringDesc{
				{Key: StringValue("z"), Value: StringValue("x")},
				{Key: StringValue("f"), Value: StringValue("p")},
				{Key: StringValue("a"), Value: StringValue("k")},
				{Key: StringValue("b"), Value: StringValue("l")},
			},
			want: ElementsByValueStringDesc{
				{Key: StringValue("z"), Value: StringValue("x")},
				{Key: StringValue("f"), Value: StringValue("p")},
				{Key: StringValue("b"), Value: StringValue("l")},
				{Key: StringValue("a"), Value: StringValue("k")},
			},
		},
		{
			name: "should sort by int value ascending",
			sortBy: ElementsByValueIntAsc{
				{Key: IntValue(10), Value: IntValue(29)},
				{Key: IntValue(2), Value: IntValue(37)},
				{Key: IntValue(4), Value: IntValue(1)},
				{Key: IntValue(50), Value: IntValue(8)},
			},
			want: ElementsByValueIntAsc{
				{Key: IntValue(4), Value: IntValue(1)},
				{Key: IntValue(50), Value: IntValue(8)},
				{Key: IntValue(10), Value: IntValue(29)},
				{Key: IntValue(2), Value: IntValue(37)},
			},
		},
		{
			name: "should sort by int value descending",
			sortBy: ElementsByValueIntDesc{
				{Key: IntValue(10), Value: IntValue(29)},
				{Key: IntValue(2), Value: IntValue(37)},
				{Key: IntValue(4), Value: IntValue(1)},
				{Key: IntValue(50), Value: IntValue(8)},
			},
			want: ElementsByValueIntDesc{
				{Key: IntValue(2), Value: IntValue(37)},
				{Key: IntValue(10), Value: IntValue(29)},
				{Key: IntValue(50), Value: IntValue(8)},
				{Key: IntValue(4), Value: IntValue(1)},
			},
		},
		{
			name: "should sort by key int value ascending even though the data type is numeric strings, data type conversions are handled by the Data implementations",
			sortBy: ElementsByKeyIntAsc{
				{Key: StringValue("10"), Value: IntValue(29)},
				{Key: StringValue("2"), Value: IntValue(37)},
				{Key: StringValue("4"), Value: IntValue(1)},
				{Key: StringValue("50"), Value: IntValue(8)},
			},
			want: ElementsByKeyIntAsc{
				{Key: StringValue("2"), Value: IntValue(37)},
				{Key: StringValue("4"), Value: IntValue(1)},
				{Key: StringValue("10"), Value: IntValue(29)},
				{Key: StringValue("50"), Value: IntValue(8)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(tt.sortBy)
			if !cmp.Equal(tt.sortBy, tt.want) {
				t.Fatalf("TestElementsSorters() should sort by %T, got %+v, want %+v", tt.sortBy, tt.sortBy, tt.want)
			}
		})
	}
}
