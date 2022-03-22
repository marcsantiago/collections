package set

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestSet_Add(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want *Set[int]
	}{
		{
			name: "should add values without issue",
			args: args{data: []int{1, 2, 3, 4, 5}},
			want: New(1, 2, 3, 4, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New[int]()
			for _, d := range tt.args.data {
				s.Add(d)
			}

			if !cmp.Equal(s, tt.want, deepAllowUnexported(s, tt.want)) {
				t.Fatalf("Add() got %v want %v", s, tt.want)
			}
		})
	}
}

func TestSet_Clear(t *testing.T) {
	tests := []struct {
		name string
		s    *Set[int]
		want int
	}{
		{
			name: "full set should be empty",
			s:    New(1, 2, 3, 4, 5),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Clear()
			if tt.s.Size() != tt.want {
				t.Fatalf("Clear() got %v want %v", tt.s.Size(), tt.want)
			}
		})
	}
}

func TestSet_Contains(t *testing.T) {
	type args struct {
		data int
	}
	tests := []struct {
		name string
		s    *Set[int]
		args args
		want bool
	}{
		{
			name: "should contain the data",
			s:    New(1, 2, 3, 4, 5),
			args: args{data: 2},
			want: true,
		},
		{
			name: "should not contain the data",
			s:    New(1, 2, 3, 4, 5),
			args: args{data: 10},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.data); got != tt.want {
				t.Fatalf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Copy(t *testing.T) {
	tests := []struct {
		name string
		s    *Set[int]
		want *Set[int]
	}{
		{
			name: "should copy the set identically",
			s:    New(1, 2, 3, 4, 5),
			want: New(1, 2, 3, 4, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Copy(); !cmp.Equal(got, tt.want, deepAllowUnexported(got, tt.want)) {
				t.Fatalf("Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Difference(t *testing.T) {
	type args struct {
		sets []*Set[string]
	}
	tests := []struct {
		name string
		s    *Set[string]
		args args
		want *Set[string]
	}{
		{
			name: "should return the difference",
			args: args{sets: []*Set[string]{
				New("google", "microsoft", "apple"),
				New("banana"),
			}},
			s:    New("apple", "banana", "cherry"),
			want: New("cherry"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Difference(tt.args.sets...); !cmp.Equal(got, tt.want, deepAllowUnexported(got, tt.want)) {
				t.Fatalf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_DifferenceUpdate(t *testing.T) {
	type args struct {
		sets []*Set[string]
	}
	tests := []struct {
		name string
		s    *Set[string]
		args args
		want *Set[string]
	}{
		{
			name: "should return the difference",
			args: args{sets: []*Set[string]{
				New("google", "microsoft", "apple"),
				New("banana"),
			}},
			s:    New("apple", "banana", "cherry"),
			want: New("cherry"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.DifferenceUpdate(tt.args.sets...)
			if !cmp.Equal(tt.s, tt.want, deepAllowUnexported(tt.s, tt.want)) {
				t.Fatalf("DifferenceUpdate() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_Discard(t *testing.T) {
	type args struct {
		data int
	}
	tests := []struct {
		name string
		s    *Set[int]
		args args
		want *Set[int]
	}{
		{
			name: "should remove the data from the set",
			s:    New(1, 2, 3, 4, 5),
			args: args{data: 1},
			want: New(2, 3, 4, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Discard(tt.args.data)
			if !cmp.Equal(tt.s, tt.want, deepAllowUnexported(tt.s, tt.want)) {
				t.Fatalf("Discard() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_Intersection(t *testing.T) {
	type args struct {
		sets []*Set[string]
	}
	tests := []struct {
		name string
		s    *Set[string]
		args args
		want *Set[string]
	}{
		{
			name: "should return the intersection",
			args: args{sets: []*Set[string]{
				New("google", "microsoft", "apple", "banana"),
				New("apple", "banana"),
			}},
			s:    New("apple", "banana", "cherry"),
			want: New("apple", "banana"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Intersection(tt.args.sets...); !cmp.Equal(got, tt.want, deepAllowUnexported(tt.s, tt.want)) {
				t.Fatalf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IntersectionUpdate(t *testing.T) {
	type args struct {
		sets []*Set[string]
	}
	tests := []struct {
		name string
		s    *Set[string]
		args args
		want *Set[string]
	}{
		{
			name: "should return the intersection",
			args: args{sets: []*Set[string]{
				New("google", "microsoft", "apple", "banana"),
				New("apple", "banana"),
			}},
			s:    New("apple", "banana", "cherry"),
			want: New("apple", "banana"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.IntersectionUpdate(tt.args.sets...)
			if got := tt.s.Intersection(tt.args.sets...); !cmp.Equal(got, tt.want, deepAllowUnexported(tt.s, tt.want)) {
				t.Fatalf("IntersectionUpdate() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_IsDisjoint(t *testing.T) {
	type args struct {
		ss *Set[string]
	}
	tests := []struct {
		name string
		s    *Set[string]
		args args
		want bool
	}{
		{
			name: "sets should be disjointed",
			s:    New("apple", "banana", "cherry"),
			args: args{
				New("google", "microsoft"),
			},
			want: true,
		},
		{
			name: "sets should not be disjointed",
			s:    New("apple", "banana", "cherry"),
			args: args{
				New("apple", "google", "microsoft"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsDisjoint(tt.args.ss); got != tt.want {
				t.Fatalf("IsDisjoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IsSubset(t *testing.T) {
	type args struct {
		ss *Set[string]
	}
	tests := []struct {
		name string
		s    *Set[string]
		args args
		want bool
	}{
		{
			name: "sets should be a subset",
			s:    New("apple", "banana", "cherry"),
			args: args{
				New("apple", "banana", "cherry", "google", "microsoft"),
			},
			want: true,
		},
		{
			name: "sets should not be a subset",
			s:    New("apple", "banana", "cherry"),
			args: args{
				New("apple", "google", "microsoft"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsSubset(tt.args.ss); got != tt.want {
				t.Fatalf("IsSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IsSuperset(t *testing.T) {
	type args struct {
		ss *Set[string]
	}
	tests := []struct {
		name string
		s    *Set[string]
		args args
		want bool
	}{
		{
			name: "sets should be a super set",
			s:    New("apple", "banana", "cherry", "google", "microsoft"),
			args: args{
				New("apple", "banana", "cherry"),
			},
			want: true,
		},
		{
			name: "sets should not be a super set",
			s:    New("apple", "banana", "cherry"),
			args: args{
				New("apple", "google", "microsoft"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsSuperset(tt.args.ss); got != tt.want {
				t.Fatalf("IsSuperset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Pop(t *testing.T) {
	tests := []struct {
		name     string
		s        *Set[string]
		popTwice bool
		want     any
		ok       bool
	}{
		{
			name: "should randomly remove an item from the set",
			s:    New("apple"),
			want: "apple",
			ok:   true,
		},
		{
			name:     "should return nothing",
			s:        New("apple"),
			popTwice: true,
			want:     "",
			ok:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.popTwice {
				tt.s.Pop()
			}
			if got, ok := tt.s.Pop(); !cmp.Equal(got, tt.want) || !cmp.Equal(ok, tt.ok) {
				t.Fatalf("Pop() = %v, %v, want %v, %v", got, ok, tt.want, tt.ok)
			}
		})
	}
}

func TestSet_Remove(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		s    *Set[string]
		want *Set[string]
	}{
		{
			name: "should  remove an item from the set",
			args: args{key: "cherry"},
			s:    New("apple", "banana", "cherry"),
			want: New("apple", "banana"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Remove(tt.args.key)
			if !cmp.Equal(tt.s, tt.want, deepAllowUnexported(tt.s, tt.want)) {
				t.Fatalf("Remove() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_SymmetricDifference(t *testing.T) {
	type args struct {
		ss *Set[string]
	}
	tests := []struct {
		name string
		s    *Set[string]
		args args
		want *Set[string]
	}{
		{
			name: "should return symmetric difference",
			s:    New("apple", "banana", "cherry"),
			args: args{
				ss: New("google", "microsoft", "apple"),
			},
			want: New("banana", "cherry", "microsoft", "google"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SymmetricDifference(tt.args.ss); !cmp.Equal(got, tt.want, deepAllowUnexported(tt.s, tt.want)) {
				t.Fatalf("SymmetricDifference() = got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_SymmetricDifferenceUpdate(t *testing.T) {
	type args struct {
		ss *Set[string]
	}
	tests := []struct {
		name string
		s    *Set[string]
		args args
		want *Set[string]
	}{
		{
			name: "should return symmetric difference",
			s:    New("apple", "banana", "cherry"),
			args: args{
				ss: New("google", "microsoft", "apple"),
			},
			want: New("banana", "cherry", "microsoft", "google"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.SymmetricDifferenceUpdate(tt.args.ss)
			if !cmp.Equal(tt.s, tt.want, deepAllowUnexported(tt.s, tt.want)) {
				t.Fatalf("SymmetricDifferenceUpdate() = got %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_Union(t *testing.T) {
	type args struct {
		sets []*Set[string]
	}
	tests := []struct {
		name string
		s    *Set[string]
		args args
		want *Set[string]
	}{
		{
			name: "should add all the sets together",
			s:    New("apple", "banana", "cherry"),
			args: args{sets: []*Set[string]{
				New("google", "microsoft", "apple", "banana"),
				New("apple", "banana"),
			}},
			want: New("apple", "banana", "cherry", "google", "microsoft"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Union(tt.args.sets...); !cmp.Equal(got, tt.want, deepAllowUnexported(tt.s, tt.want)) {
				t.Fatalf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Update(t *testing.T) {
	type args struct {
		sets []*Set[string]
	}
	tests := []struct {
		name string
		s    *Set[string]
		args args
		want *Set[string]
	}{
		{
			name: "should add all the sets together",
			s:    New("apple", "banana", "cherry"),
			args: args{sets: []*Set[string]{
				New("google", "microsoft", "apple", "banana"),
				New("apple", "banana"),
			}},
			want: New("apple", "banana", "cherry", "google", "microsoft"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Update(tt.args.sets...)
			if !cmp.Equal(tt.s, tt.want, deepAllowUnexported(tt.s, tt.want)) {
				t.Fatalf("Update() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func BenchmarkSet_Add(b *testing.B) {
	set := New[int]()
	for i := 0; i < 10000; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Contains(200)
	}
}

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
