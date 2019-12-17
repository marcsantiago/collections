package namedtuple

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/marcsantiago/collections"
)

func TestNamedTuple(t *testing.T) {
	type args struct {
		labels []string
		data   [][]collections.Data
	}
	tests := []struct {
		name       string
		args       args
		want       NamedTuple
		wantAge    int
		wantCity   string
		wantGender string
	}{
		{
			name: "should create a NamedTuple and add data without issue",
			args: args{
				labels: []string{"age", "gender", "city"},
				data: [][]collections.Data{
					[]collections.Data{collections.IntValue(10), collections.StringValue("male"), collections.StringValue("NY")},
					[]collections.Data{collections.IntValue(20), collections.StringValue("male"), collections.StringValue("TX")},
					[]collections.Data{collections.IntValue(30), collections.StringValue("female"), collections.StringValue("NJ")},
				},
			},
			want: NamedTuple{
				labels: map[string]int{
					"age":    0,
					"gender": 1,
					"city":   2,
				},
				elements: [][]collections.Data{
					[]collections.Data{collections.IntValue(10), collections.StringValue("male"), collections.StringValue("NY")},
					[]collections.Data{collections.IntValue(20), collections.StringValue("male"), collections.StringValue("TX")},
					[]collections.Data{collections.IntValue(30), collections.StringValue("female"), collections.StringValue("NJ")},
				},
			},
			wantAge:    10,
			wantCity:   "TX",
			wantGender: "female",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.labels...)
			for _, data := range tt.args.data {
				got = got.Add(data)
			}

			if !cmp.Equal(tt.want, got, deepAllowUnexported(tt.want, got)) {
				t.Fatalf("NamedTuple() created named tuple did not match expected\n%+v", cmp.Diff(tt.want, got, deepAllowUnexported(tt.want, got)))
			}

			age := got.Get(0, "age")
			city := got.Get(1, "city")
			gender := got.Get(2, "gender")

			if !cmp.Equal(tt.wantAge, age.Int()) {
				t.Fatalf("NamedTuple() want age %v, got %v", tt.wantAge, age.Int())
			}

			if !cmp.Equal(tt.wantCity, city.String()) {
				t.Fatalf("NamedTuple() want city %v, got %v", tt.wantCity, city.String())
			}

			if !cmp.Equal(tt.wantGender, gender.String()) {
				t.Fatalf("NamedTuple() want gender %v, got %v", tt.wantGender, gender.String())
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
