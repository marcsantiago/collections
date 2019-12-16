package counter

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/marcsantiago/collections"
)

func TestCounter(t *testing.T) {
	type args struct {
		data         []collections.Data
		optionalType []collections.Type
	}
	tests := []struct {
		name string
		args args
		want collections.CounterMap
	}{
		{
			name: "should count int slice without type passed in",
			args: args{
				data:         collections.IntValues{2, 2, 3, 4, 4, 4, 10, 11, 11}.Data(),
				optionalType: nil,
			},
			want: IntMap{2: 2, 3: 1, 4: 3, 10: 1, 11: 2},
		},
		{
			name: "should count int slice with type passed in",
			args: args{
				data:         collections.IntValues{2, 2, 3, 4, 4, 4, 10, 11, 11}.Data(),
				optionalType: []collections.Type{collections.IntSliceType},
			},
			want: IntMap{2: 2, 3: 1, 4: 3, 10: 1, 11: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Counter(tt.args.data, tt.args.optionalType...)
			if !cmp.Equal(tt.want, c) {
				t.Fatalf("Counter() got values are not expected diff\n%+v", cmp.Diff(tt.want, c))
			}
		})
	}
}

func TestCounter_MostCommon(t *testing.T) {
	type args struct {
		data         []collections.Data
		optionalType []collections.Type
	}
	tests := []struct {
		name string
		args args
		want []collections.Element
	}{
		{
			name: "should count int slice with type passed in and get the 3 most common",
			args: args{
				data:         collections.IntValues{2, 2, 3, 4, 4, 4, 10, 11, 11}.Data(),
				optionalType: []collections.Type{collections.IntSliceType},
			},
			want: []collections.Element{
				{Key: collections.IntValue(4), Value: collections.IntValue(3)},
				{Key: collections.IntValue(2), Value: collections.IntValue(2)},
				{Key: collections.IntValue(11), Value: collections.IntValue(2)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Counter(tt.args.data, tt.args.optionalType...)
			elements := c.MostCommon(3)
			sort.Sort(collections.ElementsByValueIntDesc(elements))
			sort.Sort(collections.ElementsByValueIntDesc(tt.want))
			if !cmp.Equal(tt.want, elements) {
				t.Fatalf("Counter() got values are not expected diff\n%+v", cmp.Diff(tt.want, elements))
			}
		})
	}
}

//BenchmarkCounterNoType-12      	 1605631	       738 ns/op	     494 B/op	       3 allocs/op
//BenchmarkCounterNoType-12      	 1641736	       717 ns/op	     494 B/op	       3 allocs/op
//BenchmarkCounterNoType-12      	 1683136	       708 ns/op	     494 B/op	       3 allocs/op
func BenchmarkCounterNoType(b *testing.B) {
	var c collections.CounterMap
	_ = c
	data := collections.IntValues([]int{2, 2, 3, 4, 4, 4, 10, 11, 11, 1, 1, 2, 7, 7, 8, 8, 9, 10, 19, 20}).Data()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c = Counter(data)
	}
}

//BenchmarkCounterWithType-12    	 1700364	       709 ns/op	     494 B/op	       3 allocs/op
//BenchmarkCounterWithType-12    	 1691558	       698 ns/op	     494 B/op	       3 allocs/op
//BenchmarkCounterWithType-12    	 1666542	       700 ns/op	     494 B/op	       3 allocs/op
func BenchmarkCounterWithType(b *testing.B) {
	var c collections.CounterMap
	_ = c
	data := collections.IntValues([]int{2, 2, 3, 4, 4, 4, 10, 11, 11, 1, 1, 2, 7, 7, 8, 8, 9, 10, 19, 20}).Data()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c = Counter(data, collections.IntSliceType)
	}
}
