package counter

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/marcsantiago/collections"
)

func TestStringMap_Delete(t *testing.T) {
	type args struct {
		key collections.Data
	}
	tests := []struct {
		name string
		i    StringMap
		args args
		want StringMap
	}{
		{
			name: "should delete data from the map by key",
			i: StringMap{
				"foo": 1,
				"bar": 10,
			},
			args: args{key: collections.StringValue("foo")},
			want: StringMap{
				"bar": 10,
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

func TestStringMap_Get(t *testing.T) {
	type args struct {
		key collections.Data
	}
	tests := []struct {
		name  string
		i     StringMap
		args  args
		want  collections.Data
		want1 bool
	}{
		{
			name: " should get a value without issue",
			i: StringMap{
				"foo": 1,
				"bar": 10,
			},
			args:  args{key: collections.StringValue("foo")},
			want:  collections.IntValue(1),
			want1: true,
		},
		{
			name: " should not find a value",
			i: StringMap{
				"foo": 1,
				"bar": 10,
			},
			args:  args{key: collections.StringValue("zip")},
			want:  collections.IntValue(0),
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.i.Get(tt.args.key)
			if !cmp.Equal(got, tt.want) {
				t.Fatalf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Fatalf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStringMap_Items(t *testing.T) {
	tests := []struct {
		name string
		i    StringMap
		want []collections.Element
	}{
		{
			name: "should return items",
			i: StringMap{
				"foo": 1,
				"bar": 10,
				"the": 1000,
				"i":   10000,
			},
			want: []collections.Element{
				collections.Element{Key: collections.StringValue("foo"), Value: collections.IntValue(1)},
				collections.Element{Key: collections.StringValue("bar"), Value: collections.IntValue(10)},
				collections.Element{Key: collections.StringValue("the"), Value: collections.IntValue(1000)},
				collections.Element{Key: collections.StringValue("i"), Value: collections.IntValue(10000)},
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

func TestStringMap_Iterate(t *testing.T) {
	tests := []struct {
		name string
		i    StringMap
		want []collections.Element
	}{
		{
			name: " should get a value without issue",
			i: StringMap{
				"foo": 1,
				"bar": 10,
				"the": 1000,
				"i":   10000,
			},
			want: []collections.Element{
				collections.Element{Key: collections.StringValue("foo"), Value: collections.IntValue(1)},
				collections.Element{Key: collections.StringValue("bar"), Value: collections.IntValue(10)},
				collections.Element{Key: collections.StringValue("the"), Value: collections.IntValue(1000)},
				collections.Element{Key: collections.StringValue("i"), Value: collections.IntValue(10000)},
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

func TestStringMap_MostCommon(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		i    StringMap
		args args
		want []collections.Element
	}{
		{
			name: "should return the most common 3 words",
			i: StringMap{
				"foo": 1,
				"bar": 10,
				"the": 1000,
				"i":   10000,
			},
			want: []collections.Element{
				collections.Element{Key: collections.StringValue("i"), Value: collections.IntValue(10000)},
				collections.Element{Key: collections.StringValue("the"), Value: collections.IntValue(1000)},
				collections.Element{Key: collections.StringValue("bar"), Value: collections.IntValue(10)},
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

func TestStringMap_Set(t *testing.T) {
	type args struct {
		key   collections.Data
		value collections.Data
	}
	tests := []struct {
		name string
		i    StringMap
		args args
		want int
	}{
		{
			name: "should override the counter value with the set value",
			i: StringMap{
				"foo": 1,
			},
			args: args{
				key:   collections.StringValue("foo"),
				value: collections.IntValue(100),
			},
			want: 100,
		},
		{
			name: "should not set the counter value",
			i:    make(StringMap),
			args: args{
				key:   collections.StringValue("foo"),
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

func TestStringMap_Subtract(t *testing.T) {
	type args struct {
		key collections.Data
	}
	tests := []struct {
		name string
		i    StringMap
		args args
		want int
	}{
		{
			name: "should decrement from the value",
			i: StringMap{
				"foo": 1,
				"bar": 10,
				"the": 1000,
				"i":   10000,
			},
			args: args{key: collections.StringValue("foo")},
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

func TestStringMap_Update(t *testing.T) {
	type args struct {
		key collections.Data
	}
	tests := []struct {
		name string
		i    StringMap
		args args
		want int
	}{
		{
			name: "should increment from the value",
			i: StringMap{
				"foo": 1,
				"bar": 10,
				"the": 1000,
				"i":   10000,
			},
			args: args{key: collections.StringValue("foo")},
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
