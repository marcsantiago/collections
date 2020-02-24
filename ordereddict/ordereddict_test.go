package ordereddict

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/marcsantiago/collections"
)

func TestOrderedDict(t *testing.T) {
	tests := []struct {
		name   string
		rangeN int
		delN   []int
		want   OrderedDict
	}{
		{
			name:   "should set, get, delete with no issue",
			rangeN: 10,
			delN:   []int{0, 5, 8},
			want: OrderedDict{
				keys: []collections.Data{
					collections.IntValue(1),
					collections.IntValue(2),
					collections.IntValue(3),
					collections.IntValue(4),
					collections.IntValue(6),
					collections.IntValue(7),
					collections.IntValue(9),
				},
				hash: map[collections.Data]collections.Data{
					collections.IntValue(1): collections.IntValue(2),
					collections.IntValue(2): collections.IntValue(4),
					collections.IntValue(3): collections.IntValue(6),
					collections.IntValue(4): collections.IntValue(8),
					collections.IntValue(6): collections.IntValue(12),
					collections.IntValue(7): collections.IntValue(14),
					collections.IntValue(9): collections.IntValue(18),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			od := New()
			for i := 0; i < tt.rangeN; i++ {
				key, value := i, i*2
				od.Set(collections.IntValue(key), collections.IntValue(value))
			}

			for _, d := range tt.delN {
				od.Delete(collections.IntValue(d))
			}

			if !cmp.Equal(od, tt.want, deepAllowUnexported(od, tt.want)) {
				t.Fatalf("OrderedDict() %+v\n", cmp.Diff(od, tt.want, deepAllowUnexported(od, tt.want)))
			}
		})
	}
}

func TestOrderedDict_Items(t *testing.T) {
	tests := []struct {
		name   string
		rangeN int
		delN   []int
		want   []collections.Element
	}{
		{
			name:   "should return items in the same order they were inserted",
			rangeN: 10,
			delN:   []int{0, 5, 8},
			want: []collections.Element{
				{Key: collections.IntValue(1), Value: collections.IntValue(2)},
				{Key: collections.IntValue(2), Value: collections.IntValue(4)},
				{Key: collections.IntValue(3), Value: collections.IntValue(6)},
				{Key: collections.IntValue(4), Value: collections.IntValue(8)},
				{Key: collections.IntValue(6), Value: collections.IntValue(12)},
				{Key: collections.IntValue(7), Value: collections.IntValue(14)},
				{Key: collections.IntValue(9), Value: collections.IntValue(18)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			od := New()
			for i := 0; i < tt.rangeN; i++ {
				key, value := i, i*2
				od.Set(collections.IntValue(key), collections.IntValue(value))
			}

			for _, d := range tt.delN {
				od.Delete(collections.IntValue(d))
			}

			items := od.Items()
			if !cmp.Equal(items, tt.want) {
				t.Fatalf("items() %+v\n", cmp.Diff(od, tt.want))
			}
		})
	}
}

// get around unexported mutex
func deepAllowUnexported(vs ...interface{}) cmp.Option {
	m := make(map[reflect.Type]struct{})
	for _, v := range vs {
		structTypes(reflect.ValueOf(v), m)
	}
	var typs []interface{}
	for t := range m {
		typs = append(typs, reflect.New(t).Elem().Interface())
	}
	return cmp.AllowUnexported(typs...)
}

func structTypes(v reflect.Value, m map[reflect.Type]struct{}) {
	if !v.IsValid() {
		return
	}
	switch v.Kind() {
	case reflect.Ptr:
		if !v.IsNil() {
			structTypes(v.Elem(), m)
		}
	case reflect.Interface:
		if !v.IsNil() {
			structTypes(v.Elem(), m)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			structTypes(v.Index(i), m)
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			structTypes(v.MapIndex(k), m)
		}
	case reflect.Struct:
		m[v.Type()] = struct{}{}
		for i := 0; i < v.NumField(); i++ {
			structTypes(v.Field(i), m)
		}
	}
}
