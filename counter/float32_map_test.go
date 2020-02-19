package counter

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/marcsantiago/collections"
)

func TestFloatMap32_Delete(t *testing.T) {
	type args struct {
		key collections.Data
	}
	tests := []struct {
		name string
		i    FloatMap32
		args args
		want FloatMap32
	}{
		{
			name: "should delete data from the map by key",
			i: FloatMap32{
				1: 1,
				2: 10,
			},
			args: args{key: collections.FloatValue32(1)},
			want: FloatMap32{
				2: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.i.Delete(tt.args.key)
			if !cmp.Equal(tt.i, tt.want) {
				t.Fatalf("Delete(): got %s want %s", tt.i, tt.want)
			}
		})
	}
}

func TestFloatMap32_Get(t *testing.T) {
	type args struct {
		key collections.Data
	}
	tests := []struct {
		name  string
		i     FloatMap32
		args  args
		want  collections.Data
		want1 bool
	}{
		{
			name: " should get a value without issue",
			i: FloatMap32{
				1: 1,
				2: 10,
			},
			args:  args{key: collections.FloatValue32(1)},
			want:  collections.IntValue(1),
			want1: true,
		},
		{
			name: " should not find a value",
			i: FloatMap32{
				1: 1,
				2: 10,
			},
			args:  args{key: collections.FloatValue32(10)},
			want:  collections.IntValue(0),
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.i.Get(tt.args.key)
			if !cmp.Equal(got.Float32(), tt.want.Float32()) {
				t.Fatalf("Get() value got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Fatalf("Get() found got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFloatMap32_Items(t *testing.T) {
	tests := []struct {
		name string
		i    FloatMap32
		want []collections.Element
	}{
		{
			name: "should return items",
			i: FloatMap32{
				1: 1,
				2: 10,
				3: 1000,
				4: 10000,
			},
			want: []collections.Element{
				collections.Element{Key: collections.FloatValue32(1), Value: collections.IntValue(1)},
				collections.Element{Key: collections.FloatValue32(2), Value: collections.IntValue(10)},
				collections.Element{Key: collections.FloatValue32(3), Value: collections.IntValue(1000)},
				collections.Element{Key: collections.FloatValue32(4), Value: collections.IntValue(10000)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.i.Items()
			sort.Sort(collections.ElementsByValueIntAsc(got))
			if !cmp.Equal(got, tt.want) {
				t.Fatalf("Items() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatMap32_Iterate(t *testing.T) {
	tests := []struct {
		name string
		i    FloatMap32
		want []collections.Element
	}{
		{
			name: " should get a value without issue",
			i: FloatMap32{
				1: 1,
				2: 10,
				3: 1000,
				4: 10000,
			},
			want: []collections.Element{
				collections.Element{Key: collections.FloatValue32(1), Value: collections.IntValue(1)},
				collections.Element{Key: collections.FloatValue32(2), Value: collections.IntValue(10)},
				collections.Element{Key: collections.FloatValue32(3), Value: collections.IntValue(1000)},
				collections.Element{Key: collections.FloatValue32(4), Value: collections.IntValue(10000)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []collections.Element
			for elem := range tt.i.Iterate() {
				got = append(got, elem)
			}

			sort.Sort(collections.ElementsByValueIntAsc(got))
			if !cmp.Equal(got, tt.want) {
				t.Fatalf("Iterate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatMap32_MostCommon(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		i    FloatMap32
		args args
		want []collections.Element
	}{
		{
			name: "should return the most common 3 words",
			i: FloatMap32{
				1: 1,
				2: 10,
				3: 1000,
				4: 10000,
			},
			want: []collections.Element{
				collections.Element{Key: collections.FloatValue32(4), Value: collections.IntValue(10000)},
				collections.Element{Key: collections.FloatValue32(3), Value: collections.IntValue(1000)},
				collections.Element{Key: collections.FloatValue32(2), Value: collections.IntValue(10)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.MostCommon(tt.args.n); !cmp.Equal(got, tt.want) {
				t.Fatalf("MostCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatMap32_Set(t *testing.T) {
	type args struct {
		key   collections.Data
		value collections.Data
	}
	tests := []struct {
		name string
		i    FloatMap32
		args args
		want int
	}{
		{
			name: "should override the counter value with the set value",
			i: FloatMap32{
				1: 1,
			},
			args: args{
				key:   collections.FloatValue32(1),
				value: collections.IntValue(100),
			},
			want: 100,
		},
		{
			name: "should not set the counter value",
			i:    make(FloatMap32),
			args: args{
				key:   collections.FloatValue32(1),
				value: collections.IntValue(1000),
			},
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.i.Set(tt.args.key, tt.args.value)
			got, _ := tt.i.Get(tt.args.key)
			if !cmp.Equal(got.Int(), tt.want) {
				t.Fatalf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatMap32_Subtract(t *testing.T) {
	type args struct {
		key collections.Data
	}
	tests := []struct {
		name string
		i    FloatMap32
		args args
		want int
	}{
		{
			name: "should decrement from the value",
			i: FloatMap32{
				1: 1,
				2: 10,
				3: 1000,
				4: 10000,
			},
			args: args{key: collections.FloatValue32(1)},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.i.Subtract(tt.args.key)
			got, _ := tt.i.Get(tt.args.key)
			if !cmp.Equal(got.Int(), tt.want) {
				t.Fatalf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatMap32_Update(t *testing.T) {
	type args struct {
		key collections.Data
	}
	tests := []struct {
		name string
		i    FloatMap32
		args args
		want int
	}{
		{
			name: "should increment from the value",
			i: FloatMap32{
				1: 1,
				2: 10,
				3: 1000,
				4: 10000,
			},
			args: args{key: collections.FloatValue32(1)},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.i.Update(tt.args.key)
			got, _ := tt.i.Get(tt.args.key)
			if !cmp.Equal(got.Int(), tt.want) {
				t.Fatalf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
