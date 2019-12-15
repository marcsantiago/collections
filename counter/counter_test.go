package counter

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIntSliceCounter_Next(t *testing.T) {
	type fields struct {
		arr []int
	}
	tests := []struct {
		name       string
		fields     fields
		wantValues []int
	}{
		{
			name: "basic test",
			fields: fields{
				arr: []int{2, 2, 3, 4, 4, 4, 10, 11, 11},
			},
			wantValues: []int{2, 1, 3, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Ints(tt.wantValues)
			var values []int
			o := NewIntSliceCounter(tt.fields.arr)
			for element := range o.Next() {
				values = append(values, element.Value.Int())
			}
			sort.Ints(values)
			if !cmp.Equal(tt.wantValues, values) {
				t.Fatalf("IntSliceCounter_Next() got values are not expected diff\n%+v", cmp.Diff(tt.wantValues, values))
			}
		})
	}
}
