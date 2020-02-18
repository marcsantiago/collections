package chain

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/marcsantiago/collections"
)

func TestMaps_Delete(t *testing.T) {
	type args struct {
		key collections.Data
	}
	tests := []struct {
		name string
		m    Maps
		args args
		want Maps
	}{
		{
			name: "should delete from map 1",
			m: Maps{
				collections.GenericMap{collections.StringValue("foo"): collections.StringValue("bar")},
				collections.GenericMap{collections.StringValue("zip"): collections.StringValue("zap")},
				collections.GenericMap{collections.StringValue("one"): collections.StringValue("two")},
			},
			args: args{key: collections.StringValue("foo")},
			want: Maps{
				collections.GenericMap{},
				collections.GenericMap{collections.StringValue("zip"): collections.StringValue("zap")},
				collections.GenericMap{collections.StringValue("one"): collections.StringValue("two")},
			},
		},
		{
			name: "should delete nothing",
			m: Maps{
				collections.GenericMap{collections.StringValue("zip"): collections.StringValue("zap")},
				collections.GenericMap{collections.StringValue("foo"): collections.StringValue("bar")},
				collections.GenericMap{collections.StringValue("one"): collections.StringValue("two")},
			},
			args: args{key: collections.StringValue("foo")},
			want: Maps{
				collections.GenericMap{collections.StringValue("zip"): collections.StringValue("zap")},
				collections.GenericMap{collections.StringValue("foo"): collections.StringValue("bar")},
				collections.GenericMap{collections.StringValue("one"): collections.StringValue("two")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Delete(tt.args.key)
			if !cmp.Equal(tt.m, tt.want) {
				t.Fatalf("Delete() got %+v want %+v", tt.m, tt.want)
			}
		})
	}
}

func TestMaps_Get(t *testing.T) {
	type args struct {
		key collections.Data
	}
	tests := []struct {
		name  string
		m     Maps
		args  args
		want  collections.Data
		want1 bool
	}{
		{
			name: "should find key in map 1",
			m: Maps{
				collections.GenericMap{collections.StringValue("foo"): collections.StringValue("bar")},
				collections.GenericMap{collections.StringValue("zip"): collections.StringValue("zap")},
				collections.GenericMap{collections.StringValue("one"): collections.StringValue("two")},
			},
			args:  args{key: collections.StringValue("foo")},
			want:  collections.StringValue("bar"),
			want1: true,
		},
		{
			name: "should find key in map 3",
			m: Maps{
				collections.GenericMap{collections.StringValue("foo"): collections.StringValue("bar")},
				collections.GenericMap{collections.StringValue("zip"): collections.StringValue("zap")},
				collections.GenericMap{collections.StringValue("one"): collections.StringValue("two")},
			},
			args:  args{key: collections.StringValue("one")},
			want:  collections.StringValue("two"),
			want1: true,
		},
		{
			name: "should find nothing",
			m: Maps{
				collections.GenericMap{collections.StringValue("foo"): collections.StringValue("bar")},
				collections.GenericMap{collections.StringValue("zip"): collections.StringValue("zap")},
				collections.GenericMap{collections.StringValue("one"): collections.StringValue("two")},
			},
			args:  args{key: collections.StringValue("foobar")},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.m.Get(tt.args.key)
			if !cmp.Equal(got, tt.want) {
				t.Fatalf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Fatalf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMaps_Items(t *testing.T) {
	tests := []struct {
		name string
		m    Maps
		want []collections.Element
	}{
		{
			name: "should get all the items from the chain",
			m: Maps{
				collections.GenericMap{collections.StringValue("foo"): collections.StringValue("bar")},
				collections.GenericMap{collections.StringValue("zip"): collections.StringValue("zap")},
				collections.GenericMap{collections.StringValue("one"): collections.StringValue("two")},
			},
			want: []collections.Element{
				collections.Element{Key: collections.StringValue("foo"), Value: collections.StringValue("bar")},
				collections.Element{Key: collections.StringValue("zip"), Value: collections.StringValue("zap")},
				collections.Element{Key: collections.StringValue("one"), Value: collections.StringValue("two")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Items(); !cmp.Equal(got, tt.want) {
				t.Fatalf("Items() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaps_Iterate(t *testing.T) {
	tests := []struct {
		name string
		m    Maps
		want []collections.Element
	}{
		{
			name: "should iterate over the chain",
			m: Maps{
				collections.GenericMap{collections.StringValue("foo"): collections.StringValue("bar")},
				collections.GenericMap{collections.StringValue("zip"): collections.StringValue("zap")},
				collections.GenericMap{collections.StringValue("one"): collections.StringValue("two")},
			},
			want: []collections.Element{
				collections.Element{Key: collections.StringValue("foo"), Value: collections.StringValue("bar")},
				collections.Element{Key: collections.StringValue("zip"), Value: collections.StringValue("zap")},
				collections.Element{Key: collections.StringValue("one"), Value: collections.StringValue("two")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []collections.Element
			for ele := range tt.m.Iterate() {
				got = append(got, ele)
			}
			if !cmp.Equal(got, tt.want) {
				t.Fatalf("Iterate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaps_NewChild(t *testing.T) {
	type args struct {
		mm collections.Map
	}
	tests := []struct {
		name string
		m    Maps
		args args
		want Maps
	}{
		{
			name: "should create a new child with empty map 1",
			m: Maps{
				collections.GenericMap{collections.StringValue("foo"): collections.StringValue("bar")},
				collections.GenericMap{collections.StringValue("zip"): collections.StringValue("zap")},
				collections.GenericMap{collections.StringValue("one"): collections.StringValue("two")},
			},
			args: args{mm: nil},
			want: Maps{
				collections.GenericMap{},
				collections.GenericMap{collections.StringValue("foo"): collections.StringValue("bar")},
				collections.GenericMap{collections.StringValue("zip"): collections.StringValue("zap")},
				collections.GenericMap{collections.StringValue("one"): collections.StringValue("two")},
			},
		},
		{
			name: "should create a new child with new map",
			m: Maps{
				collections.GenericMap{collections.StringValue("foo"): collections.StringValue("bar")},
				collections.GenericMap{collections.StringValue("zip"): collections.StringValue("zap")},
				collections.GenericMap{collections.StringValue("one"): collections.StringValue("two")},
			},
			args: args{mm: collections.GenericMap{collections.StringValue("candy"): collections.StringValue("land")}},
			want: Maps{
				collections.GenericMap{collections.StringValue("candy"): collections.StringValue("land")},
				collections.GenericMap{collections.StringValue("foo"): collections.StringValue("bar")},
				collections.GenericMap{collections.StringValue("zip"): collections.StringValue("zap")},
				collections.GenericMap{collections.StringValue("one"): collections.StringValue("two")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.NewChild(tt.args.mm); !cmp.Equal(got, tt.want) {
				t.Fatalf("NewChild() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestMaps_Parents(t *testing.T) {
	tests := []struct {
		name string
		m    Maps
		want Maps
	}{
		{
			name: "should get remove the parent map node",
			m: Maps{
				collections.GenericMap{collections.StringValue("foo"): collections.StringValue("bar")},
				collections.GenericMap{collections.StringValue("zip"): collections.StringValue("zap")},
				collections.GenericMap{collections.StringValue("one"): collections.StringValue("two")},
			},
			want: Maps{
				collections.GenericMap{collections.StringValue("zip"): collections.StringValue("zap")},
				collections.GenericMap{collections.StringValue("one"): collections.StringValue("two")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Parents(); !cmp.Equal(got, tt.want) {
				t.Fatalf("Parents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaps_Set(t *testing.T) {
	type args struct {
		key   collections.Data
		value collections.Data
	}
	tests := []struct {
		name string
		m    Maps
		args args
		want Maps
	}{
		{
			name: "should update map 1 only",
			m: Maps{
				collections.GenericMap{collections.StringValue("foo"): collections.StringValue("bar")},
				collections.GenericMap{collections.StringValue("zip"): collections.StringValue("zap")},
				collections.GenericMap{collections.StringValue("one"): collections.StringValue("two")},
			},
			args: args{
				key:   collections.StringValue("candy"),
				value: collections.StringValue("land"),
			},
			want: Maps{
				collections.GenericMap{
					collections.StringValue("foo"):   collections.StringValue("bar"),
					collections.StringValue("candy"): collections.StringValue("land"),
				},
				collections.GenericMap{collections.StringValue("zip"): collections.StringValue("zap")},
				collections.GenericMap{collections.StringValue("one"): collections.StringValue("two")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Set(tt.args.key, tt.args.value)
			if !cmp.Equal(tt.m, tt.want) {
				t.Fatalf("Set() = %v, want %v", tt.m, tt.want)
			}
		})
	}
}
