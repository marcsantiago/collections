package counter

import (
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
